package sqlx

import (
	"context"
	"fmt"
	"strings"
)

func (r *database) ListFilterByColumn(ctx context.Context, filters *Filter) (users []*UserResponse, err error) {
	selectClause := "SELECT * FROM users "  // notice the space at the end?
	paginateClause := " LIMIT 30 OFFSET 0;" // and at the beginning?

	whereClauses := make([]string, 0)
	arguments := make([]interface{}, 0)
	if filters.Email != "" {
		// Note that if postgres database created was non-deterministic (for case-insensitive)
		// then we need to append `COLLATE case_insensitive` at the end of
		// the full query.
		//
		// Here we simply lower our input text and lower the database value.
		whereClauses = append(whereClauses, "LOWER(email) = ?")
		arguments = append(arguments, strings.ToLower(filters.Email))
	}
	if filters.FirstName != "" {
		whereClauses = append(whereClauses, "LOWER(first_name) = ?")
		arguments = append(arguments, strings.ToLower(filters.FirstName))
	}
	if filters.FavouriteColour != "" {
		whereClauses = append(whereClauses, "favourite_colour = ?")
		arguments = append(arguments, strings.ToLower(filters.FavouriteColour))
	}

	fullQuery := selectClause

	if len(whereClauses) > 0 {
		fullQuery += "WHERE "
		for _, clause := range whereClauses {
			fullQuery += clause
		}
	}

	fullQuery += paginateClause

	fullQuery = r.db.Rebind(fullQuery)

	rows, err := r.db.QueryxContext(ctx, fullQuery, arguments...)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for rows.Next() {
		var u userDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
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

func (r *database) ListFilterSort(ctx context.Context, filters *Filter) (users []*UserResponse, err error) {
	selectClause := "SELECT * FROM users "
	paginateClause := " LIMIT 30 OFFSET 0;"
	sortClauses := ""

	fullQuery := selectClause

	if len(filters.Base.Sort) > 0 {
		sortClauses += " ORDER BY "
		for col, order := range filters.Base.Sort {
			sortClauses += fmt.Sprintf(" %s ", col)
			sortClauses += fmt.Sprintf(" %s ", order)
		}
	}

	fullQuery += sortClauses
	fullQuery += paginateClause

	fullQuery = r.db.Rebind(fullQuery)

	rows, err := r.db.QueryxContext(ctx, fullQuery)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for rows.Next() {
		var u userDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	return users, nil
}

func (r *database) ListFilterPagination(ctx context.Context, filters *Filter) (users []*UserResponse, err error) {
	selectClause := "SELECT * FROM users "

	orderClause := " ORDER BY id"
	paginateClause := " LIMIT ? OFFSET ?"

	fullQuery := selectClause
	fullQuery += orderClause
	fullQuery += paginateClause

	fullQuery = r.db.Rebind(fullQuery)

	rows, err := r.db.QueryxContext(ctx, fullQuery, filters.Base.Limit, filters.Base.Offset)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for rows.Next() {
		var u userDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
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
