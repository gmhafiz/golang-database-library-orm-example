package sqlc

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"godb/respond"
	sqlx2 "godb/sqlx"
)

type handler struct {
	db *Queries
}

func Handle(r *chi.Mux, db *sqlx.DB) {
	h := &handler{
		db: New(db),
	}

	r.Route("/api/sqlc/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
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

	user, err := h.db.CreateUser(r.Context(), CreateUserParams{
		FirstName: request.FirstName,
		MiddleName: sql.NullString{
			String: request.MiddleName,
			Valid:  true,
		},
		LastName: request.LastName,
		Email:    request.Email,
		Password: hash,
	})
	if err != nil {
		return
	}

	respond.Json(w, http.StatusCreated, sqlx2.UserResponse{
		ID:         uint(user.ID),
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName.String,
		LastName:   user.LastName,
		Email:      user.Email,
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListUsers(r.Context())
	if err != nil {
		http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
		return
	}

	var userResponse []sqlx2.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, sqlx2.UserResponse{
			ID:         uint(user.ID),
			FirstName:  user.FirstName,
			MiddleName: user.MiddleName.String,
			LastName:   user.LastName,
			Email:      user.Email,
		})
	}

	respond.Json(w, http.StatusOK, userResponse)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

}
