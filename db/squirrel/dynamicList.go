package squirrel

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"godb/db"
)

func (r repository) ListFilterByColumn(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	builder := r.db.
		Select("*").
		From("users").
		OrderBy("id")

	if f.Email != "" {
		//builder = builder.Where(sq.Eq{"email": f.Email})
		builder = builder.Where(sq.Like{"email": "%" + f.Email + "%"})
	}

	if f.FirstName != "" {
		//builder = builder.Where(sq.Eq{"first_name": f.FirstName})
		builder = builder.Where(sq.Like{"first_name": "%" + f.FirstName + "%"})
	}

	if f.FavouriteColour != "" {
		//builder = builder.Where(sq.Eq{"favourite_colour": f.FavouriteColour})
		builder = builder.Where(sq.Like{"favourite_colour": "%" + f.FavouriteColour + "%"})
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
		var u db.UserDB
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
		users = append(users, &db.UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			UpdatedAt:       u.UpdatedAt.String(),
		})
	}

	return users, nil
}

func (r repository) ListFilterSort(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	builder := r.db.Select("*").
		From("users")

	for col, order := range f.Base.Sort {
		// the following commented line is vulnerable to sql injection if
		// user input is not sanitised
		//builder = builder.OrderBy(col + " " + order)

		switch col {
		case "first_name":
			builder = builder.OrderBy("first_name" + " " + order)
		case "last_name":
			builder = builder.OrderBy("last_name" + " " + order)
		case "middle_name":
			builder = builder.OrderBy("middle_name" + " " + order)
		case "email":
			builder = builder.OrderBy("email" + " " + order)
		case "favourite_colour":
			builder = builder.OrderBy("favourite_colour" + " " + order)
		}
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
		var u db.UserDB
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
		users = append(users, &db.UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			UpdatedAt:       u.UpdatedAt.String(),
		})
	}

	return users, nil
}

func (r repository) ListFilterPagination(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
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
		var u db.UserDB
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
		users = append(users, &db.UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			UpdatedAt:       u.UpdatedAt.String(),
		})
	}

	return users, nil
}
