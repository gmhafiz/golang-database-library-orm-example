package sqlx

import (
	"context"
	"fmt"
	"strings"
	"time"

	"godb/db"
)

func (r *repository) ListFilterByColumn(ctx context.Context, filters *db.Filter) (users []*db.UserResponse, err error) {
	selectClause := "SELECT * FROM users " // notice the space at the end?
	paginateClause := " LIMIT 30 OFFSET 0" // and at the beginning?

	whereClauses := make([]string, 0)
	arguments := make([]interface{}, 0)

	if filters.FirstName != "" {
		// Note that if postgres collation was non-deterministic (for case-insensitive)
		// then we need to append `COLLATE case_insensitive` at the end of
		// the full query.
		//
		// Modify existing column for case_insensitive collation with
		// 		CREATE COLLATION case_insensitive (provider = icu, locale = 'und-u-ks-level2', deterministic = false);
		//      ALTER TABLE users ALTER COLUMN  first_name SET DATA TYPE varchar COLLATE "case_insensitive";
		//
		// Here we simply lower our input text and lower the repository value.
		whereClauses = append(whereClauses, "LOWER(first_name) = ?")
		arguments = append(arguments, strings.ToLower(filters.FirstName))
	}
	if filters.FavouriteColour != "" {
		whereClauses = append(whereClauses, "favourite_colour = ?")
		arguments = append(arguments, strings.ToLower(filters.FavouriteColour))
	}
	if filters.Email != "" {
		whereClauses = append(whereClauses, "LOWER(email) = ?")
		arguments = append(arguments, strings.ToLower(filters.Email))
	}

	fullQuery := selectClause

	if len(whereClauses) > 0 {
		fullQuery += "WHERE "
		fullQuery += strings.Join(whereClauses, " AND ")
	}

	fullQuery += " ORDER by id" // space at beginning

	fullQuery += paginateClause

	fullQuery = r.db.Rebind(fullQuery)

	stmt, err := r.db.PrepareContext(ctx, fullQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(stmt)

	rows, err := r.db.QueryxContext(ctx, fullQuery, arguments...)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for rows.Next() {
		var u db.UserDB
		err = rows.StructScan(&u)
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
			Tags:            u.Tags,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (r *repository) ListFilterSort(ctx context.Context, filters *db.Filter) (users []*db.UserResponse, err error) {
	selectClause := "SELECT * FROM users "
	paginateClause := " LIMIT 10 OFFSET 0;"
	sortClauses := ""

	fullQuery := selectClause

	sortClauses += " ORDER BY "
	sortJoined := ""
	var sort []string
	for col, order := range filters.Base.Sort {
		sort = append(sort, fmt.Sprintf(" %s ", col)+" "+fmt.Sprintf(" %s ", order))
		sortJoined = strings.Join(sort, ",")
	}
	sortClauses += sortJoined

	fullQuery += sortClauses
	fullQuery += paginateClause

	fullQuery = r.db.Rebind(fullQuery)

	// try to protect against sql injection
	//_, err = r.db.Prepare(fullQuery)
	//if err != nil {
	//	return nil, err
	//}

	rows, err := r.db.QueryxContext(ctx, fullQuery)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	for rows.Next() {
		var u db.UserDB
		err = rows.StructScan(&u)
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
			Tags:            u.Tags,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (r *repository) ListFilterPagination(ctx context.Context, filters *db.Filter) (users []*db.UserResponse, err error) {
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
		var u db.UserDB
		err = rows.StructScan(&u)
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
			Tags:            u.Tags,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}
