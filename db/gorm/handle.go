package gorm

import (
	"encoding/json"
	"github.com/alexedwards/argon2id"
	"godb/db/sqlx"
	"godb/respond"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func Handle(r *chi.Mux, db *gorm.DB) {
	h := &handler{
		db: db,
	}

	r.Route("/api/gorm/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
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

	h.db.Create(&UserGorm{
		FirstName:  request.FirstName,
		MiddleName: request.MiddleName,
		LastName:   request.LastName,
		Email:      request.Email,
		Password:   hash,
	})

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	var users []UserGorm
	err := h.db.Model(&UserGorm{}).Select("*").Scan(&users).Error
	if err != nil {
		http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
		return
	}

	var userResponse []sqlx.UserResponse
	for _, u := range users {
		userResponse = append(userResponse, sqlx.UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	var user UserGorm
	h.db.Take(&user)

	userResponse := sqlx.UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Email:      user.Email,
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

}