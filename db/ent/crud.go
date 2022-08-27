package ent

import (
	"context"
	"errors"
	"fmt"
	"godb/db"
	"log"

	"godb/db/ent/ent/gen"
	"godb/db/ent/ent/gen/user"
)

func (r *database) Create(ctx context.Context, request *db.UserRequest, hash string) (*gen.User, error) {
	saved, err := r.db.Debug().User.Create().
		SetFirstName(request.FirstName).
		SetNillableMiddleName(nil). // Does not insert anything to this column
		//SetMiddleName(request.MiddleName). // Inserts empty string
		SetLastName(request.LastName).
		SetEmail(request.Email).
		SetFavouriteColour(user.FavouriteColour(request.FavouriteColour)).
		SetPassword(hash).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error saving user: %w", err)
	}

	return saved, nil
}

func (r *database) List(ctx context.Context, f *filter) ([]*gen.User, error) {
	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}

	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}

	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	if f.PaginateLastID != 0 {
		return r.ListFilterPaginationByID(ctx, f)
	}

	if len(f.LastNames) > 0 {
		return r.ListFilterWhereIn(ctx, f)
	}

	return r.db.User.Query().
		Order(gen.Asc(user.FieldID)).
		Limit(30).
		Offset(0).
		All(ctx)
}

func (r *database) Get(ctx context.Context, userID uint64) (*gen.User, error) {
	u, err := r.db.User.Query().Where(user.ID(uint(userID))).First(ctx)
	if err != nil {
		if gen.IsNotFound(err) {
			return nil, errors.New("no record found")
		}
		return nil, err
	}

	return u, nil
}

func (r *database) Update(ctx context.Context, userID int64, req *db.UserUpdateRequest) (*gen.User, error) {
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

func (r *database) ListFilterWhereIn(ctx context.Context, f *filter) ([]*gen.User, error) {
	return r.db.User.Query().
		Where(user.LastNameIn(f.LastNames...)).
		All(ctx)
}
