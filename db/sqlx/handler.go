package sqlx

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"godb/param"
	"godb/respond"
)

type handler struct {
	db *database
}

func Register(r *chi.Mux, db *sqlx.DB) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/sqlx/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/m2m", h.ListM2M)
		router.Get("/{userID}", h.Get)
		router.Get("/{userID}/address", h.Countries)
		router.Put("/{userID}", h.Update)
		router.Delete("/{userID}", h.Delete)
	})

	r.Route("/api/sqlx/country", func(router chi.Router) {
		router.Get("/", h.Countries)
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

	u, err := h.db.Create(r.Context(), &request, hash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				http.Error(w, `{"message": "`+ErrUniqueKeyViolation.Error()+`"}`, http.StatusBadRequest)
				return
			default:
				http.Error(w, `{"message": "`+ErrDefault.Error()+`"}`, http.StatusBadRequest)
				return
			}
		}

		http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusOK, &UserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName.String,
		LastName:   u.LastName,
		Email:      u.Email,
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.List(r.Context())
	if err != nil {
		http.Error(w, `{"message": `+err.Error()+`}`, http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusOK, users)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+err.Error()+`}`, http.StatusBadRequest)
		return
	}

	u, err := h.db.Get(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, u)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+err.Error()+`}`, http.StatusBadRequest)
		return
	}

	var req UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"message": "bad request"}`, http.StatusBadRequest)
		return
	}

	u, err := h.db.Update(r.Context(), userID, &req)
	if err != nil {
		http.Error(w, `{"message": "error updating"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, &UserResponse{
		ID:         uint(userID),
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName,
		LastName:   u.LastName,
		Email:      u.Email,
	})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		http.Error(w, `{"message": `+param.ErrParam.Error()+`}`, http.StatusBadRequest)
		return
	}

	_, err = h.db.Delete(r.Context(), userID)
	if err != nil {
		http.Error(w, `{"message": "error deleting"}`, http.StatusBadRequest)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.db.Countries(r.Context())
	if err != nil {
		http.Error(w, `{"message": "error retrieving"}`, http.StatusInternalServerError)
		return
	}

	respond.Json(w, http.StatusOK, addresses)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		http.Error(w, `{"message": `+err.Error()+`}`, http.StatusBadRequest)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
