package ent

import (
	"context"
	"fmt"

	"godb/db/ent/ent/gen"
	"godb/db/ent/ent/gen/user"
	"godb/db/sqlx"
)

func (r *database) Create(ctx context.Context, request sqlx.UserRequest, hash string) (*gen.User, error) {
	saved, err := r.db.User.Create().
		SetFirstName(request.FirstName).
		SetNillableMiddleName(&request.MiddleName).
		SetLastName(request.LastName).
		SetEmail(request.Email).
		SetPassword(hash).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error saving user")
	}

	return saved, nil
}

func (r *database) List(ctx context.Context) ([]*gen.User, error) {
	all, err := r.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (r *database) Get(ctx context.Context, userID uint64) (*gen.User, error) {
	u, err := r.db.User.Query().Where(user.ID(uint(userID))).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("db scanning error")
	}
	return u, nil
}

func (r *database) Update(ctx context.Context, userID int64, req *sqlx.UserUpdateRequest) (*gen.User, error) {
	_, err := r.db.User.UpdateOneID(uint(userID)).
		SetFirstName(req.FirstName).
		SetNillableMiddleName(&req.MiddleName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, uint64(userID))
}

func (r *database) Delete(ctx context.Context, userID int64) error {
	return r.db.User.DeleteOneID(uint(userID)).Exec(ctx)
}
