package sqlc

import (
	"context"
	"database/sql"
	"log"
)

func (r *database) ListFilterByColumn(ctx context.Context, f *Filter) (l []ListUsersRow, err error) {
	p := ListDynamicUsersParams{
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
		l = append(l, ListUsersRow{
			ID:        user.ID,
			FirstName: user.FirstName,
			MiddleName: sql.NullString{
				String: user.MiddleName.String,
				Valid:  true,
			},
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: user.FavouriteColour,
		})
	}

	return l, nil
}

func (r *database) ListFilterSort(ctx context.Context, f *Filter) (l []ListUsersRow, err error) {
	p := ListDynamicUsersParams{
		// fixme: optional filter favourite_colour
		SqlOffset: int32(f.Base.Offset),
		SqlLimit:  int32(f.Base.Limit),
	}

	if len(f.Base.Sort) > 0 {
		for col, order := range f.Base.Sort {
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
		l = append(l, ListUsersRow{
			ID:        user.ID,
			FirstName: user.FirstName,
			MiddleName: sql.NullString{
				String: user.MiddleName.String,
				Valid:  true,
			},
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: user.FavouriteColour,
		})
	}
	return l, nil
}

func (r *database) ListFilterPagination(ctx context.Context, f *Filter) (l []ListUsersRow, err error) {
	p := ListDynamicUsersParams{
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
		l = append(l, ListUsersRow{
			ID:        user.ID,
			FirstName: user.FirstName,
			MiddleName: sql.NullString{
				String: user.MiddleName.String,
				Valid:  true,
			},
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: user.FavouriteColour,
		})
	}
	return l, nil
}

func (r *database) ListFilterWhereIn(ctx context.Context, f *Filter) ([]ListUsersRow, error) {
	return []ListUsersRow{}, nil
}
