package sqlc

import (
	"context"
	"godb/db"
	"godb/db/sqlc/pg"
	"log"
	"time"
)

func (r *database) ListFilterByColumn(ctx context.Context, f *db.Filter) (l []db.UserResponse, err error) {
	p := pg.ListDynamicUsersParams{
		FirstName: f.FirstName,
		Email:     f.Email,
		//FavouriteColourPresent: f.FavouriteColour,
		//FavouriteColour:        ValidColours(f.FavouriteColour),
		SqlOffset: int32(f.Base.Offset),
		SqlLimit:  int32(f.Base.Limit),
	}
	dynamicList, err := r.db.ListDynamicUsers(ctx, p)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, user := range dynamicList {
		l = append(l, db.UserResponse{
			ID:              uint(user.ID),
			FirstName:       user.FirstName,
			MiddleName:      user.MiddleName.String,
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: string(user.FavouriteColour),
			Tags:            user.Tags,
			UpdatedAt:       user.UpdatedAt.Format(time.RFC3339),
		})
	}

	return l, nil
}

func (r *database) ListFilterSort(ctx context.Context, f *db.Filter) (l []db.UserResponse, err error) {
	p := pg.ListDynamicUsersParams{
		// fixme: optional filter favourite_colour
		SqlOffset: int32(f.Base.Offset),
		SqlLimit:  int32(f.Base.Limit),
	}

	if len(f.Base.Sort) > 0 {
		for col, order := range f.Base.Sort {
			// repeat with each column you want to sort.
			if col == "first_name" {
				if order == "ASC" {
					p.FirstNameAsc = col
				} else {
					p.FirstNameDesc = col
				}
			}
		}
	}

	dynamicList, err := r.db.ListDynamicUsers(ctx, p)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, user := range dynamicList {
		l = append(l, db.UserResponse{
			ID:              uint(user.ID),
			FirstName:       user.FirstName,
			MiddleName:      user.MiddleName.String,
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: string(user.FavouriteColour),
			Tags:            user.Tags,
			UpdatedAt:       user.UpdatedAt.Format(time.RFC3339),
		})
	}

	return l, nil
}

func (r *database) ListFilterPagination(ctx context.Context, f *db.Filter) (l []db.UserResponse, err error) {
	p := pg.ListDynamicUsersParams{
		// fixme: optional filter favourite_colour
		SqlOffset: int32(f.Base.Offset),
		SqlLimit:  int32(f.Base.Limit),
	}

	dynamicList, err := r.db.ListDynamicUsers(ctx, p)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, user := range dynamicList {
		l = append(l, db.UserResponse{
			ID:              uint(user.ID),
			FirstName:       user.FirstName,
			MiddleName:      user.MiddleName.String,
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: string(user.FavouriteColour),
			Tags:            user.Tags,
			UpdatedAt:       user.UpdatedAt.Format(time.RFC3339),
		})
	}
	return l, nil
}
