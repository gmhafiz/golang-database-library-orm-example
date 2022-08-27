// Code generated by sqlc. DO NOT EDIT.
// source: mariadb-query.sql

package mariadb

import (
	"context"
	"database/sql"
)

const countriesWithAddress = `-- name: CountriesWithAddress :many
SELECT c.id AS country_id,
       c.name,
       c.code,
       a.id AS address_id,
       a.line_1,
       a.line_2,
       a.postcode,
       a.city,
       a.state
FROM countries c
         LEFT JOIN addresses a on c.id = a.country_id
ORDER BY c.id
`

type CountriesWithAddressRow struct {
	CountryID int64
	Name      string
	Code      string
	AddressID sql.NullInt64
	Line1     sql.NullString
	Line2     sql.NullString
	Postcode  sql.NullInt32
	City      sql.NullString
	State     sql.NullString
}

func (q *Queries) CountriesWithAddress(ctx context.Context) ([]CountriesWithAddressRow, error) {
	rows, err := q.db.QueryContext(ctx, countriesWithAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CountriesWithAddressRow
	for rows.Next() {
		var i CountriesWithAddressRow
		if err := rows.Scan(
			&i.CountryID,
			&i.Name,
			&i.Code,
			&i.AddressID,
			&i.Line1,
			&i.Line2,
			&i.Postcode,
			&i.City,
			&i.State,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const countriesWithAddressAggregate = `-- name: CountriesWithAddressAggregate :many
select id, code, name, json_arrayagg from country_address
`

func (q *Queries) CountriesWithAddressAggregate(ctx context.Context) ([]CountryAddress, error) {
	rows, err := q.db.QueryContext(ctx, countriesWithAddressAggregate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CountryAddress
	for rows.Next() {
		var i CountryAddress
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.JsonArrayagg,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createUser = `-- name: CreateUser :exec


INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	Password        string
	FavouriteColour UsersFavouriteColour
}

// OFFSET ? LIMIT ? ;
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.FavouriteColour,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, middle_name, last_name, email, favourite_colour
FROM users
WHERE id = ?
`

type GetUserRow struct {
	ID              int64
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	FavouriteColour UsersFavouriteColour
}

func (q *Queries) GetUser(ctx context.Context, id int64) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.Email,
		&i.FavouriteColour,
	)
	return i, err
}

const listDynamicUsers = `-- name: ListDynamicUsers :many
SELECT id, first_name, middle_name, last_name, email, password, favourite_colour
FROM users
WHERE (@first_name = '' OR first_name LIKE '%' || @first_name || '%')
  AND (@email = '' OR email = LOWER(@email) )
ORDER BY (CASE
              WHEN @first_name_desc = 'first_name' THEN first_name
              WHEN @email_desc = 'email' THEN email
    END) DESC,
         (CASE
              WHEN @first_name_asc = 'first_name' THEN first_name
              WHEN @email_asc = 'email' THEN email
             END)
`

type ListDynamicUsersRow struct {
	ID              int64
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	Password        string
	FavouriteColour UsersFavouriteColour
}

func (q *Queries) ListDynamicUsers(ctx context.Context) ([]ListDynamicUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listDynamicUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListDynamicUsersRow
	for rows.Next() {
		var i ListDynamicUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.FavouriteColour,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, middle_name, last_name, email, favourite_colour
FROM users
ORDER BY id
LIMIT 30
OFFSET 0
`

type ListUsersRow struct {
	ID              int64
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	FavouriteColour UsersFavouriteColour
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.FavouriteColour,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAddress = `-- name: SelectAddress :many
SELECT a.id, a.line_1, a.line_2, a.postcode, a.city, a.state, a.country_id
FROM addresses a
WHERE a.id = ?
`

func (q *Queries) SelectAddress(ctx context.Context, id int64) ([]Address, error) {
	rows, err := q.db.QueryContext(ctx, selectAddress, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Address
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.Line1,
			&i.Line2,
			&i.Postcode,
			&i.City,
			&i.State,
			&i.CountryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserAddresses = `-- name: SelectUserAddresses :many
SELECT DISTINCT ua.user_id, ua.address_id
FROM addresses a
         LEFT JOIN user_addresses ua ON a.id = ua.address_id
WHERE ua.user_id = ?
`

func (q *Queries) SelectUserAddresses(ctx context.Context, userID sql.NullInt64) ([]UserAddress, error) {
	rows, err := q.db.QueryContext(ctx, selectUserAddresses, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAddress
	for rows.Next() {
		var i UserAddress
		if err := rows.Scan(&i.UserID, &i.AddressID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUsers = `-- name: SelectUsers :many
SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email, u.favourite_colour
FROM users u
LIMIT 30
`

type SelectUsersRow struct {
	ID              int64
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	FavouriteColour UsersFavouriteColour
}

func (q *Queries) SelectUsers(ctx context.Context) ([]SelectUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, selectUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUsersRow
	for rows.Next() {
		var i SelectUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.FavouriteColour,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectWhereInLastNames = `-- name: SelectWhereInLastNames :many
SELECT id, first_name, middle_name, last_name, email, password, favourite_colour, updated_at FROM users WHERE last_name in (?)
`

func (q *Queries) SelectWhereInLastNames(ctx context.Context, lastName string) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, selectWhereInLastNames, lastName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.FavouriteColour,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = ?,
    middle_name = ?,
    last_name = ?,
    email = ?,
    favourite_colour = ?
WHERE id = ?
`

type UpdateUserParams struct {
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	Email           string
	FavouriteColour UsersFavouriteColour
	ID              int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.Email,
		arg.FavouriteColour,
		arg.ID,
	)
	return err
}