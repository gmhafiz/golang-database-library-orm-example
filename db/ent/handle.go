package ent

import (
	"encoding/json"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"

	"godb/db/ent/ent/gen"
	"godb/db/sqlx"
	"godb/param"
	"godb/respond"
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
	var request sqlx.UserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `{"message": "bad request"}`, http.StatusBadRequest)
		return
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		http.Error(w, `{"message": "internal error"}`, http.StatusInternalServerError)
		return
	}

	saved, err := h.db.Create(r.Context(), request, hash)
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusCreated, &sqlx.UserResponse{
		ID:         saved.ID,
		FirstName:  saved.FirstName,
		MiddleName: *saved.MiddleName,
		LastName:   saved.LastName,
		Email:      saved.Email,
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	all, err := h.db.List(r.Context())
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	var users []*sqlx.UserResponse
	for _, u := range all {
		users = append(users, &sqlx.UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: *u.MiddleName,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	respond.Json(w, http.StatusOK, users)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.UInt64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	u, err := h.db.Get(r.Context(), userID)
	if err != nil {
		http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, &sqlx.UserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		MiddleName: *u.MiddleName,
		LastName:   u.LastName,
		Email:      u.Email,
	})
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	var req sqlx.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"message": "bad request"}`, http.StatusBadRequest)
		return
	}

	updated, err := h.db.Update(r.Context(), userID, &req)
	if err != nil {
		http.Error(w, `{"message": "error updating"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, updated)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	err = h.db.Delete(r.Context(), userID)
	if err != nil {
		http.Error(w, `{"message": "error deleting"}`, http.StatusBadRequest)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	all, err := h.db.Countries(r.Context())

	if err != nil {
		http.Error(w, `{"message": "error retrieving"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, all)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		http.Error(w, `{"message": `+err.Error()+`}`, http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
