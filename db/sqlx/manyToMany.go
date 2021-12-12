package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	UserWithAddresses = `
		SELECT u.*, a.* 
		FROM users u 
			LEFT JOIN user_addresses ua on u.id = ua.user_id 
			LEFT JOIN addresses a on a.id = ua.address_id 
		WHERE u.id = $1`

	Users = `
		SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email 
		FROM "users" u 
		LIMIT 30;`
	UsersAddress = `
		SELECT DISTINCT ua.address_id 
		FROM "addresses" a 
			LEFT JOIN "user_addresses" ua ON a.id = ua.address_id 
		WHERE ua.user_id IN (?);`
	Address = `
		SELECT a.* 
		FROM addresses a
		WHERE a.id in (?);
`
)

type userAddress struct {
	ID int `db:"id"`
}

func (r *database) ListM2M(ctx context.Context) (*UserResponseWithAddress, error) {
	users, err := r.db.QueryContext(ctx, Users)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}
	defer users.Close()

	var all []*UserResponseWithAddress
	for users.Next() {
		var u user
		if err := users.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		all = append(all, &UserResponseWithAddress{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	userIDs := getUserIDs(all)

	query, args, err := sqlx.In(UsersAddress, userIDs)
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
		if err := userAddresses.Scan(&ua.ID); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		uas = append(uas, &ua)
	}
	defer userAddresses.Close()

	addressIDs := getAddressIDs(uas)
	query, args, err = sqlx.In(Address, addressIDs)
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

	// now we attach address to  the user
	for _, u := range all {
		print(u)
	}

	return nil, nil
}

func getAddressIDs(uas []*userAddress) (ids []int) {
	seen := make(map[int]bool)
	for _, item := range uas {
		ok := seen[item.ID]
		if !ok {
			ids = append(ids, item.ID)
			seen[item.ID] = true
		}
	}

	return ids
}

func getUserIDs(users []*UserResponseWithAddress) (ids []int) {
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

//func (r *database) ListAddress(ctx context.Context, userID int64) error {
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
