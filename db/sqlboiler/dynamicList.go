package sqlboiler

import (
	"context"
	"strings"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db/sqlboiler/models"
	sqlx2 "godb/db/sqlx"
)

func (r *database) ListFilterByColumn(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	var mods []qm.QueryMod

	if f.Email != "" {
		mods = append(mods, models.UserWhere.Email.EQ(strings.ToLower(f.Email)))
	}

	if f.FirstName != "" {
		// doesn't work with mismatched case
		//mods = append(mods, models.UserWhere.FirstName.EQ(strings.ToLower(f.FirstName)))

		// use `qm.Where()` to use ILIKE
		mods = append(mods, qm.Where("first_name ILIKE ?", strings.ToLower(f.FirstName)))
	}
	if f.FavouriteColour != "" {
		mods = append(mods, models.UserWhere.FavouriteColour.EQ(null.String{String: strings.ToLower(f.FavouriteColour), Valid: true}))
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	all, err := models.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &sqlx2.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour.String,
		})
	}

	return users, nil
}

func (r *database) ListFilterSort(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	var mods []qm.QueryMod

	for key := range f.Base.Sort {
		mods = append(mods, qm.OrderBy(key)) // todo: ORDER col ASC
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	all, err := models.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &sqlx2.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour.String,
		})
	}

	return users, nil
}

func (r *database) ListFilterPagination(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	var mods []qm.QueryMod

	if f.Base.Limit != 0 && !f.Base.DisablePaging {
		mods = append(mods, qm.Limit(int(f.Base.Limit)))
	}
	if f.Base.Offset != 0 && !f.Base.DisablePaging {
		mods = append(mods, qm.Offset(int(f.Base.Offset)))
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	all, err := models.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &sqlx2.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour.String,
		})
	}

	return users, nil
}
