package gorm

import (
	"encoding/json"
	"errors"
	"godb/db"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"godb/param"
	"godb/respond"
)

type handler struct {
	db *repo
}

func Register(r *chi.Mux, db *gorm.DB) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/gorm/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/m2m", h.ListM2M)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Patch("/{userID}", h.Patch)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/gorm/country", func(router chi.Router) {
		router.Get("/", h.Countries)
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	request := db.NewUserRequest()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, errors.New("bad request"))
		return
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, errors.New("internal error"))
		return
	}

	user, err := h.db.Create(r.Context(), request, hash)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusCreated, user)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	f := db.Filters(r.URL.Query())

	userResponse, err := h.db.List(r.Context(), f)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	userResponse, err := h.db.Get(r.Context(), userID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	var req db.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
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

	respond.Json(w, http.StatusOK, updated)
}

func (h *handler) Patch(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}
	var req db.UserPatchRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	updated, err := h.db.Patch(r.Context(), userID, &req)
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
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	resp, err := h.db.Countries(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, resp)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
