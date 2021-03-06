package ent

import (
	"context"
	"fmt"
	"log"

	"godb/db/ent/ent/gen"
	"godb/db/ent/ent/gen/user"
	"godb/db/sqlx"
)

func (r *database) Create(ctx context.Context, request *sqlx.UserRequest, hash string) (*gen.User, error) {
	saved, err := r.db.User.Create().
		SetFirstName(request.FirstName).
		SetNillableMiddleName(&request.MiddleName).
		SetLastName(request.LastName).
		SetEmail(request.Email).
		SetPassword(hash).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error saving user: %w", err)
	}

	return saved, nil
}

func (r *database) List(ctx context.Context, f *Filter) ([]*gen.User, error) {
	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}
	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}
	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	return r.db.User.Query().
		Limit(10).
		Offset(0).
		Order(gen.Asc(user.FieldID)).
		All(ctx)
}

func (r *database) Get(ctx context.Context, userID uint64) (*gen.User, error) {
	return r.db.User.Query().Where(user.ID(uint(userID))).First(ctx)
}

func (r *database) Update(ctx context.Context, userID int64, req *sqlx.UserUpdateRequest) (*gen.User, error) {
	return r.db.User.UpdateOneID(uint(userID)).
		SetFirstName(req.FirstName).
		SetNillableMiddleName(&req.MiddleName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		Save(ctx)
}

func (r *database) Delete(ctx context.Context, userID int64) error {
	return r.db.User.DeleteOneID(uint(userID)).Exec(ctx)
}
