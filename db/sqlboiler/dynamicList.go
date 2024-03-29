package sqlboiler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db"
	"godb/db/sqlboiler/models"
)

func (r *database) ListFilterByColumn(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	var mods []qm.QueryMod

	if f.Email != "" {
		mods = append(mods, models.UserWhere.Email.EQ(strings.ToLower(f.Email)))
	}

	if f.FirstName != "" {
		// The following doesn't work with mismatched case
		//mods = append(mods, models.UserWhere.FirstName.EQ(strings.ToLower(f.FirstName)))

		// Instead, use `qm.Where()` to use ILIKE
		mods = append(mods, qm.Where("first_name ILIKE ?", strings.ToLower(f.FirstName)))
	}
	if f.FavouriteColour != "" {
		mods = append(mods, models.UserWhere.FavouriteColour.EQ(strings.ToLower(f.FavouriteColour)))
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	boil.DebugMode = true

	all, err := models.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &db.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour,
			Tags:            i.Tags,
			UpdatedAt:       i.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (r *database) ListFilterSort(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	var mods []qm.QueryMod

	for key, order := range f.Base.Sort {
		// vulnerable to sql injection
		//mods = append(mods, qm.OrderBy(fmt.Sprintf("%s %s", key, order)))

		switch key {
		// whitelist columns.
		case "first_name", "last_name", "middle_name", "email", "favourite_colour":
			mods = append(mods, qm.OrderBy(fmt.Sprintf("%s %s", key, order)))
		}
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	all, err := models.Users(mods...).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &db.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour,
			Tags:            i.Tags,
			UpdatedAt:       i.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (r *database) ListFilterPagination(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	var mods []qm.QueryMod

	if f.Base.Limit != 0 && !f.Base.DisablePaging {
		mods = append(mods, qm.Limit(f.Base.Limit))
	}
	if f.Base.Offset != 0 && !f.Base.DisablePaging {
		mods = append(mods, qm.Offset(f.Base.Offset))
	}

	mods = append(mods, qm.OrderBy(models.UserColumns.ID))

	all, err := models.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, i := range all {
		users = append(users, &db.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour,
			Tags:            i.Tags,
			UpdatedAt:       i.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}
