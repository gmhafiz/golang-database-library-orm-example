package sqlboiler

import (
	"encoding/json"
	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	sqlx2 "godb/db/sqlx"
	"godb/param"
	"godb/respond"
	"net/http"
)

type handler struct {
	db *database
}

func Register(r *chi.Mux, db *sqlx.DB) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/sqlboiler/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/sqlboiler/country", func(router chi.Router) {
		router.Get("/", h.Countries)
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var request sqlx2.UserRequest
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

	u, err := h.db.Create(r.Context(), request, hash)
	if err != nil {
		http.Error(w, `{"message": "internal error"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusCreated, &sqlx2.UserResponse{
		ID:         uint(u.ID),
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName.String,
		LastName:   u.LastName,
		Email:      u.Email,
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	userResponse, err := h.db.List(r.Context())
	if err != nil {
		http.Error(w, `{"message": "internal error"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	u, err := h.db.Get(r.Context(), userID)
	if err != nil {
		http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, &sqlx2.UserResponse{
		ID:         uint(u.ID),
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName.String,
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

	var req sqlx2.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"message": "bad request"}`, http.StatusBadRequest)
		return
	}

	updated, err := h.db.Update(r.Context(), userID, req)
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
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
	addresses, err := h.db.Countries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusOK, addresses)
}
