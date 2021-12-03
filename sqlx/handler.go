package sqlx

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"

	"godb/respond"
)

const (
	Insert = "INSERT INTO users (first_name, middle_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)"
	List   = "SELECT * FROM users;"
)

type handler struct {
	db *sqlx.DB
}

func Handle(r *chi.Mux, db *sqlx.DB) {
	h := &handler{
		db: db,
	}

	r.Route("/api/sqlx/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/{userID}", h.Get)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var request UserRequest
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

	_, err = h.db.ExecContext(r.Context(), Insert, request.FirstName, request.MiddleName, request.LastName, request.Email, hash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				http.Error(w, `{"message": "unique key violation"}`, http.StatusBadRequest)
				return
			}
		}
		http.Error(w, `{"message": "db error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type user struct {
	ID         uint           `db:"id"`
	FirstName  string         `db:"first_name"`
	MiddleName sql.NullString `db:"middle_name"`
	LastName   string         `db:"last_name"`
	Email      string         `db:"email"`
	Password   string         `db:"password"`
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	var users []UserResponse
	rows, err := h.db.QueryContext(r.Context(), List)
	if err != nil {
		http.Error(w, `{"message": "db error"}`, http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var u user
		err := rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password)
		if err != nil {
			http.Error(w, `{"message": "db scanning error"}`, http.StatusInternalServerError)
			return
		}
		users = append(users, UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	respond.Json(w, http.StatusOK, users)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("")
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("")
}
