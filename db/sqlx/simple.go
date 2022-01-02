package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	Insert      = "INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, middle_name, last_name, email, favourite_colour"
	List        = "SELECT * FROM users LIMIT 30 OFFSET 0;"
	Get         = "SELECT * FROM users WHERE id = $1;"
	Update      = "UPDATE users set first_name=$1, middle_name=$2, last_name=$3, email=$4, favourite_colour=$5 WHERE id=$6;"
	UpdateNamed = "UPDATE users set first_name=:first_name, middle_name=:middle_name, last_name=:last_name, email=:email, favourite_colour=:favourite_colour WHERE id=:id;"
	Delete      = "DELETE FROM users where id=$1"
)

var (
	ErrUniqueKeyViolation = fmt.Errorf("unique key violation")
	ErrDefault            = fmt.Errorf("whoops, something wrong happened")
)

type database struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: db,
	}
}

func (r *database) Create(ctx context.Context, request *UserRequest, hash string) (*userDB, error) {
	var u userDB
	err := r.db.QueryRowContext(ctx, Insert,
		request.FirstName,
		request.MiddleName,
		request.LastName,
		request.Email,
		hash,
		request.FavouriteColour,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.FavouriteColour,
	)
	if err != nil {
		log.Printf("sqlx.Create: %v\n", err)
		return nil, fmt.Errorf("error creating user record: %w", err)
	}

	return &u, nil
}

func (r *database) List(ctx context.Context, filters *Filter) (users []*UserResponse, err error) {
	if filters.FirstName != "" || filters.Email != "" || filters.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, filters)
	}
	if len(filters.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, filters)
	}
	if filters.Base.Page > 1 {
		return r.ListFilterPagination(ctx, filters)
	}

	rows, err := r.db.QueryxContext(ctx, List)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user records")
	}

	for rows.Next() {
		var u userDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
		})
	}
	return users, nil
}

func (r *database) Get(ctx context.Context, userID int64) (*UserResponse, error) {
	var u userDB
	err := r.db.GetContext(ctx, &u, Get, userID)
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	return &UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
	}, nil
}

func (r *database) Update(ctx context.Context, userID int64, req *UserUpdateRequest) (*UserResponse, error) {
	currUser, err := r.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = req.MiddleName
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = req.FavouriteColour

	_, err = r.db.ExecContext(ctx, Update,
		currUser.FirstName,
		currUser.MiddleName,
		currUser.LastName,
		currUser.Email,
		currUser.FavouriteColour,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, userID)

	//type updateNamed struct {
	//	ID              int64  `db:"id"`
	//	FirstName       string `db:"first_name"`
	//	MiddleName      string `db:"middle_name"`
	//	LastName        string `db:"last_name"`
	//	Email           string `db:"email"`
	//	FavouriteColour string `db:"favourite_colour"`
	//}
	//
	//update := updateNamed{
	//	ID:              userID,
	//	FirstName:       req.FirstName,
	//	MiddleName:      req.MiddleName,
	//	LastName:        req.LastName,
	//	Email:           req.Email,
	//	FavouriteColour: req.FavouriteColour,
	//}
	//
	//return r.db.NamedExecContext(ctx, UpdateNamed, update)
}

func (r *database) Delete(ctx context.Context, userID int64) (sql.Result, error) {
	return r.db.ExecContext(ctx, Delete, userID)
}

type userDB struct {
	ID              uint           `db:"id"`
	FirstName       string         `db:"first_name"`
	MiddleName      sql.NullString `db:"middle_name"`
	LastName        string         `db:"last_name"`
	Email           string         `db:"email"`
	Password        string         `db:"password"`
	FavouriteColour string         `db:"favourite_colour"`
}
