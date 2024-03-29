package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	"godb/db"
	"godb/respond/message"
)

const (
	QueryUserWithAddresses = `
		SELECT u.*, a.* 
		FROM users u 
			LEFT JOIN user_addresses ua on u.id = ua.user_id 
			LEFT JOIN addresses a on a.id = ua.address_id 
		WHERE u.id = $1`

	QueryUsers = `
		SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email, u.updated_at
		FROM "users" u
		ORDER BY u.id
		LIMIT 30;
`
	QueryUsersAddress = `
		SELECT DISTINCT ua.user_id AS user_id, ua.address_id AS address_id 
		FROM "addresses" a 
			LEFT JOIN "user_addresses" ua ON a.id = ua.address_id 
		WHERE ua.user_id IN (?)
		ORDER BY ua.user_id;
`
	QueryAddress = `
		SELECT a.* 
		FROM addresses a
		WHERE a.id IN (?)
		ORDER BY a.id;
`
)

type userAddress struct {
	UserID    int `db:"user_id"`
	AddressID int `db:"address_id"`
}

func (r *repository) ListM2M(ctx context.Context) ([]*db.UserResponseWithAddressesSqlx, error) {
	users, err := r.db.QueryContext(ctx, QueryUsers)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}
	defer users.Close()

	var all []*db.UserResponseWithAddressesSqlx
	for users.Next() {
		var u db.UserDB
		if err := users.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.UpdatedAt); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		all = append(all, &db.UserResponseWithAddressesSqlx{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
			Tags:            u.Tags,
			Address:         []*db.AddressForCountry{},
		})
	}

	userIDs := getUserIDs(all)

	query, args, err := sqlx.In(QueryUsersAddress, userIDs)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	userAddresses, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}
	defer userAddresses.Close()

	var uas []*userAddress
	for userAddresses.Next() {
		var ua userAddress
		if err := userAddresses.Scan(&ua.UserID, &ua.AddressID); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		uas = append(uas, &ua)
	}
	defer userAddresses.Close()

	addressIDs := getAddressIDs(uas)
	query, args, err = sqlx.In(QueryAddress, addressIDs)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	addresses, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}
	defer addresses.Close()

	var allAddresses []*address
	for addresses.Next() {
		var a address
		if err := addresses.Scan(&a.ID, &a.Line1, &a.Line2, &a.Postcode, &a.City, &a.State, &a.CountryID); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		allAddresses = append(allAddresses, &a)
	}

	// now we attach address to the user
	for _, u := range uas {
		for _, user := range all {
			if u.UserID == int(user.ID) {
				for _, addr := range allAddresses {
					if addr.ID == uint(u.AddressID) {
						user.Address = append(user.Address, &db.AddressForCountry{
							ID:       addr.ID,
							Line1:    addr.Line1,
							Line2:    addr.Line2.String,
							Postcode: addr.Postcode.Int32,
							City:     addr.City.String,
							State:    addr.State.String,
						})
					}
				}
			}
		}
	}

	return all, nil
}

func (r *repository) ListM2MOneQuery(ctx context.Context) ([]*db.UserResponseWithAddressesSqlxSingleQuery, error) {
	m2mQuery := `SELECT row_to_json(row) FROM (SELECT u.id,
       u.first_name,
       u.middle_name,
       u.last_name,
       u.email,
       u.favourite_colour,
       u.tags,
       u.updated_at,
       array_to_json(array_agg(row_to_json(a.*))) AS addresses
FROM addresses a
         RIGHT JOIN user_addresses ua ON ua.address_id = a.id
         RIGHT JOIN users u on u.id = ua.user_id
GROUP BY u.id) row;
`
	var res []*db.UserResponseWithAddressesSqlxSingleQuery

	rows, err := r.db.QueryContext(ctx, m2mQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.SelectContext(ctx, &res, m2mQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return nil, &db.Err{Msg: message.ErrInternalError.Error(), Status: http.StatusInternalServerError}
	}

	return res, nil
}

func getAddressIDs(uas []*userAddress) (ids []int) {
	seen := make(map[int]bool)
	for _, item := range uas {
		ok := seen[item.AddressID]
		if !ok {
			ids = append(ids, item.AddressID)
			seen[item.AddressID] = true
		}
	}

	return ids
}

func getUserIDs(users []*db.UserResponseWithAddressesSqlx) (ids []interface{}) {
	seen := make(map[int]bool)
	for _, user := range users {
		ok := seen[int(user.ID)]
		if !ok {
			ids = append(ids, int(user.ID))
			seen[int(user.ID)] = true
		}
	}

	return ids
}

type address struct {
	ID        uint           `db:"id"`
	Line1     string         `db:"line_1"`
	Line2     sql.NullString `db:"line_2"`
	Postcode  sql.NullInt32  `db:"postcode"`
	City      sql.NullString `db:"city"`
	State     sql.NullString `db:"state"`
	CountryID uint           `db:"country_id"`
}

//func (r *repository) ListAddress(ctx context.Context, userID int64) error {
//	rows, err := r.db.QueryContext(ctx, UserWithAddresses, userID)
//	if err != nil {
//		return fmt.Errorf("db error")
//	}
//	defer rows.Close()
//
//	var items []address
//	for rows.Next() {
//		var ua address
//		if err := rows.Scan(
//			&ua.ID,
//			&ua.Line1,
//			&ua.Line2,
//			&ua.Postcode,
//			&ua.City,
//			&ua.State,
//		); err != nil {
//			return err
//		}
//		items = append(items, ua)
//
//	}
//
//	if err := rows.Close(); err != nil {
//		return err
//	}
//	if err := rows.Err(); err != nil {
//		return err
//	}
//
//	//var resp []*UserResponseWithAddress
//	//
//	//for _, i := range items {
//	//
//	//}
//
//	return nil
//}
