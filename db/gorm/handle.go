package gorm

import (
	"encoding/json"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"godb/db/sqlx"
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
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/gorm/country", func(router chi.Router) {
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

	h.db.Create(r.Context(), &request, hash)

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	userResponse, err := h.db.List(r.Context())
	if err != nil {
		http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
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

	userResponse := h.db.Get(r.Context(), userID)

	respond.Json(w, http.StatusOK, userResponse)
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

	h.db.Update(userID, &req)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	h.db.Delete(r.Context(), userID)
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {

}
