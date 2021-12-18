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
	return r.db.User.Query().
		//Select(user.FieldFirstName).
		//Where(
		//	user.ID(1),
		//).
		Limit(30).
		All(ctx)
}

func (r *database) Get(ctx context.Context, userID uint64) (*gen.User, error) {
	return r.db.User.Query().Where(user.ID(uint(userID))).First(ctx)
}

func (r *database) Update(ctx context.Context, userID int64, req *sqlx.UserUpdateRequest) (*gen.User, error) {
	return r.db.Debug().User.UpdateOneID(uint(userID)).
		SetFirstName(req.FirstName).
		SetNillableMiddleName(&req.MiddleName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		Save(ctx)
}

func (r *database) Delete(ctx context.Context, userID int64) error {
	return r.db.User.DeleteOneID(uint(userID)).Exec(ctx)
}
