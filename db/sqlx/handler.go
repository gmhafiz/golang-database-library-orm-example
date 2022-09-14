package sqlx

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"godb/db"
	"godb/param"
	"godb/respond"
	"godb/respond/message"
)

type handler struct {
	db *repository
}

//type Err struct {
//	Msg    string
//	Status int
//}
//
//func (e *Err) Error() string {
//	return e.Msg
//}

func Register(r *chi.Mux, db *sqlx.DB, dbType string) {
	h := &handler{
		db: NewRepo(db),
	}

	r.Route("/api/sqlx/user", func(router chi.Router) {
		router.Post("/", h.Create)
		router.Get("/", h.List)
		router.Get("/m2m", h.ListM2M)
		router.Get("/m2mOneQuery", h.ListM2MOneQuery)
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
	// 1. real-world application must perform request validation

	// 2. Transform into request struct
	var request db.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
		return
	}

	// Some business logic, preferably in its own layer.
	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}

	// 3. Call data access layer. Preferably called from business logic layer.
	u, err := h.db.Create(r.Context(), &request, hash)
	if err != nil {
		var errStruct *db.Err
		if errors.As(err, &errStruct) {
			respond.Error(w, http.StatusBadRequest, errStruct)
			return
		}
		respond.Error(w, http.StatusInternalServerError, message.ErrDefault)
		return
	}

	// 4. Respond with custom struct
	respond.Json(w, http.StatusCreated, &db.UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
		UpdatedAt:       u.UpdatedAt.String(),
	})
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	f := db.Filters(r.URL.Query())

	users, err := h.db.List(r.Context(), f)
	if err != nil {
		var errStruct *db.Err
		if errors.As(err, &errStruct) {
			respond.Error(w, http.StatusInternalServerError, errStruct)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, users)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	u, err := h.db.Get(r.Context(), userID)
	if err != nil {
		var errStruct *db.Err
		if errors.As(err, &errStruct) {
			respond.Error(w, errStruct.Status, errStruct)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, u)
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

	if req.FirstName == "" || req.LastName == "" ||
		req.Email == "" || req.FavouriteColour == "" {
		respond.Error(w, http.StatusBadRequest, errors.New("required field(s) is/are empty"))
		return
	}

	u, err := h.db.Update(r.Context(), f, userID, &req)
	if err != nil {
		var errStruct *db.Err
		if errors.As(err, &errStruct) {
			respond.Error(w, http.StatusInternalServerError, errStruct)
			return
		}

		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.Json(w, http.StatusOK, &db.UserResponse{
		ID:              uint(userID),
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
		UpdatedAt:       u.UpdatedAt,
	})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := param.Int64(r, "userID")
	if err != nil {
		respond.Error(w, http.StatusBadRequest, param.ErrParam)
		return
	}

	_, err = h.db.Delete(r.Context(), userID)
	if err != nil {
		var errStruct *db.Err
		if errors.As(err, &errStruct) {
			respond.Error(w, http.StatusInternalServerError, errStruct)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *handler) Countries(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.db.Countries(r.Context())
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrRetrieving)
		return
	}

	respond.Json(w, http.StatusOK, addresses)
}

func (h *handler) ListM2M(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2M(r.Context())
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Json(w, http.StatusOK, users)
}

func (h *handler) ListM2MOneQuery(w http.ResponseWriter, r *http.Request) {
	users, err := h.db.ListM2MOneQuery(r.Context())
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.Json(w, http.StatusOK, users)
}
