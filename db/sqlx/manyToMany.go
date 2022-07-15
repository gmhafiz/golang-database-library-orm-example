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
		LIMIT 30;
`
	UsersAddress = `
		SELECT DISTINCT ua.user_id AS user_id, ua.address_id AS address_id 
		FROM "addresses" a 
			LEFT JOIN "user_addresses" ua ON a.id = ua.address_id 
		WHERE ua.user_id IN (?);
`
	Address = `
		SELECT a.* 
		FROM addresses a
		WHERE a.id IN (?);
`
)

type userAddress struct {
	UserID    int `db:"user_id"`
	AddressID int `db:"address_id"`
}

func (r *repository) ListM2M(ctx context.Context) ([]*UserResponseWithAddressesSqlx, error) {
	users, err := r.db.QueryContext(ctx, Users)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}
	defer users.Close()

	var all []*UserResponseWithAddressesSqlx
	for users.Next() {
		var u userDB
		if err := users.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		all = append(all, &UserResponseWithAddressesSqlx{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
			Address:    []*AddressForCountry{},
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
		if err := userAddresses.Scan(&ua.UserID, &ua.AddressID); err != nil {
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

	// now we attach address to the user
	for _, u := range uas {
		for _, user := range all {
			if u.UserID == int(user.ID) {
				for _, addr := range allAddresses {
					if addr.ID == uint(u.AddressID) {
						user.Address = append(user.Address, &AddressForCountry{
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

func getUserIDs(users []*UserResponseWithAddressesSqlx) (ids []interface{}) {
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
