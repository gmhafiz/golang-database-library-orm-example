package sqlc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	sqlx2 "godb/db/sqlx"
)

type database struct {
	db *Queries
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: New(db),
	}
}

func (r *database) Create(ctx context.Context, request sqlx2.UserRequest, hash string) (*User, error) {
	u, err := r.db.CreateUser(ctx, CreateUserParams{
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
		return nil, err
	}

	return &u, nil
}

func (r *database) List(ctx context.Context) ([]User, error) {
	return r.db.ListUsers(ctx)
}

func (r *database) Get(ctx context.Context, userID int64) (User, error) {
	return r.db.GetUser(ctx, userID)
}

func (r *database) Update(ctx context.Context, userID int64, req *sqlx2.UserUpdateRequest) error {
	err := r.db.UpdateUser(ctx, UpdateUserParams{
		FirstName: req.FirstName,
		MiddleName: sql.NullString{
			String: req.MiddleName,
			Valid:  true,
		},
		LastName: req.LastName,
		Email:    req.Email,
		ID:       userID,
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (r *database) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteUser(ctx, id)
}
