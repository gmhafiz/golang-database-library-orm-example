package ent

import (
	"encoding/json"
	"godb/db"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"

	"godb/db/ent/ent/gen"
	"godb/param"
	"godb/respond"
	"godb/respond/message"
)

//go:generate ent generate --feature privacy --target ./ent/gen ./ent/schema

type handler struct {
	db *database
}

func Register(r *chi.Mux, db *gen.Client) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/ent/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/m2m", h.ListM2M)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/ent/country", func(router chi.Router) {
		router.Get("/", h.Countries)
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var request db.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}

	saved, err := h.db.Create(r.Context(), &request, hash)
	if err != nil {
		if gen.IsConstraintError(err) {
			respond.Error(w, http.StatusBadRequest, message.ErrUniqueKeyViolation)
			return
		}
		respond.Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}

	respond.Json(w, http.StatusCreated, saved)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	f := filters(r.URL.Query())

	all, err := h.db.List(r.Context(), f)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, all)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.UInt64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	u, err := h.db.Get(r.Context(), userID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, &db.UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      *u.MiddleName,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour.String(),
		UpdatedAt:       u.UpdatedAt.String(),
	})
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	f := db.Filters(r.URL.Query())

	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	var req db.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	updated, err := h.db.Update(r.Context(), userID, f, &req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, updated)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	err = h.db.Delete(r.Context(), userID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrDeleting)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	all, err := h.db.Countries(r.Context())

	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, all)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrDeleting)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
