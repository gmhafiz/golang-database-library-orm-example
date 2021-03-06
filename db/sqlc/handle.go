package sqlc

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	sqlx2 "godb/db/sqlx"
	"godb/param"
	"godb/respond"
	"godb/respond/message"
)

type handler struct {
	db *database
}

func Register(r *chi.Mux, db *sqlx.DB) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/sqlc/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/m2m", h.ListM2M)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/sqlc/country", func(router chi.Router) {
		router.Get("/", h.Countries)
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	request := sqlx2.NewUserRequest()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
		return
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}

	user, err := h.db.Create(r.Context(), request, hash)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusCreated, &sqlx2.UserResponse{
		ID:              uint(user.ID),
		FirstName:       user.FirstName,
		MiddleName:      user.MiddleName.String,
		LastName:        user.LastName,
		Email:           user.Email,
		FavouriteColour: string(user.FavouriteColour),
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	f := filters(r.URL.Query())

	users, err := h.db.List(r.Context(), f)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	var userResponse []*sqlx2.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, &sqlx2.UserResponse{
			ID:              uint(user.ID),
			FirstName:       user.FirstName,
			MiddleName:      user.MiddleName.String,
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: string(user.FavouriteColour),
		})
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	user, err := h.db.Get(r.Context(), userID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrDBScan)
		return
	}

	respond.Json(w, http.StatusOK, &sqlx2.UserResponse{
		ID:              uint(user.ID),
		FirstName:       user.FirstName,
		MiddleName:      user.MiddleName.String,
		LastName:        user.LastName,
		Email:           user.Email,
		FavouriteColour: string(user.FavouriteColour),
	})
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	var req sqlx2.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
		return
	}

	if req.FirstName == "" || req.MiddleName == "" || req.LastName == "" ||
		req.Email == "" || req.FavouriteColour == "" {
		respond.Error(w, http.StatusBadRequest, errors.New("required field(s) is/are empty"))
		return
	}

	updated, err := h.db.Update(r.Context(), userID, &req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, &sqlx2.UserResponse{
		ID:              uint(userID),
		FirstName:       updated.FirstName,
		MiddleName:      updated.MiddleName.String,
		LastName:        updated.LastName,
		Email:           updated.Email,
		FavouriteColour: string(updated.FavouriteColour),
	})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	err = h.db.Delete(r.Context(), userID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	res, err := h.db.Countries(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, res)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
