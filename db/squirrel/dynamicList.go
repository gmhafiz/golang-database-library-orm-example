package squirrel

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"

	sqlx2 "godb/db/sqlx"
)

func (r repository) ListFilterByColumn(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	builder := r.db.
		Select("*").
		From("users").
		OrderBy("id")

	if f.Email != "" {
		builder = builder.Where(sq.Eq{"email": f.Email})
	}

	if f.FirstName != "" {
		builder = builder.Where(sq.Eq{"first_name": f.FirstName})
	}

	if f.FavouriteColour != "" {
		builder = builder.Where(sq.Eq{"favourite_colour": f.FavouriteColour})
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var u userDB
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.FavouriteColour,
		)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &sqlx2.UserResponse{
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

func (r repository) ListFilterSort(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	builder := r.db.Select("*").
		From("users")

	for col, order := range f.Base.Sort {
		builder = builder.OrderBy(col + " " + order)
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var u userDB
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.FavouriteColour,
		)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &sqlx2.UserResponse{
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

func (r repository) ListFilterPagination(ctx context.Context, f *Filter) (users []*sqlx2.UserResponse, err error) {
	rows, err := r.db.
		Select("*").
		From("users").
		Limit(uint64(f.Base.Limit)).
		Offset(uint64(f.Base.Offset)).
		OrderBy("id").
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var u userDB
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.FavouriteColour,
		)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &sqlx2.UserResponse{
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
