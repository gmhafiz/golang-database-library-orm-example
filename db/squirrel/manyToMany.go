package squirrel

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"
)

type userAddress struct {
	UserID    int `db:"user_id"`
	AddressID int `db:"address_id"`
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

type UserWithAddresses struct {
	ID              uint           `db:"id" json:"id"`
	FirstName       string         `db:"first_name" json:"first_name"`
	MiddleName      sql.NullString `db:"middle_name" json:"middle_name"`
	LastName        string         `db:"last_name" json:"last_name"`
	Email           string         `db:"email" json:"email"`
	FavouriteColour string         `db:"favourite_colour" json:"favourite_colour"`
	UpdatedAt       string         `db:"updated_at" json:"updated_at"`

	Address []*AddressForCountry `json:"address" json:"address"`
}

type AddressForCountry struct {
	ID       uint           `db:"id" json:"id"`
	Line1    string         `db:"line_1" json:"line_1"`
	Line2    sql.NullString `db:"line_2" json:"line_2"`
	Postcode int32          `db:"postcode" json:"postcode"`
	City     sql.NullString `db:"city" json:"city"`
	State    string         `db:"state" json:"state"`
}

func (r repository) ListM2M(ctx context.Context) (interface{}, error) {
	rows, err := r.db.
		Select("u.id, u.first_name, u.middle_name, u.last_name, u.email, u.favourite_colour, u.updated_at").
		From("users u").
		QueryContext(ctx)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var all []*UserWithAddresses
	var userIDs []uint
	for rows.Next() {
		u := UserWithAddresses{Address: []*AddressForCountry{}}
		err := rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.FavouriteColour, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		all = append(all, &u)
		userIDs = append(userIDs, u.ID)
	}

	rows, err = r.db.
		Select("*").
		From("user_addresses").
		QueryContext(ctx)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	//var addressIDs []int
	var uas []*userAddress
	for rows.Next() {
		var ua userAddress
		if err := rows.Scan(&ua.UserID, &ua.AddressID); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		uas = append(uas, &ua)
		//addressIDs = append(addressIDs, ua.AddressID)
	}

	addressIDs := lo.Map(uas, func(t *userAddress, _ int) int {
		return t.AddressID
	})
	addressIDs = lo.Uniq(addressIDs)

	rows, err = r.db.
		Select("*").
		From("addresses").
		Where(sq.Eq{"addresses.id": addressIDs}).
		QueryContext(ctx)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	var allAddresses []*address
	for rows.Next() {
		var a address
		if err := rows.Scan(&a.ID, &a.Line1, &a.Line2, &a.Postcode, &a.City, &a.State, &a.CountryID); err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		allAddresses = append(allAddresses, &a)
	}

	for _, u := range uas {
		for _, user := range all {
			if uint(u.UserID) == user.ID {
				for _, addr := range allAddresses {
					if addr.ID == uint(u.AddressID) {
						user.Address = append(user.Address, &AddressForCountry{
							ID:    addr.ID,
							Line1: addr.Line1,
							Line2: sql.NullString{
								String: addr.Line2.String,
								Valid:  addr.Line2.Valid,
							},
							Postcode: addr.Postcode.Int32,
							City: sql.NullString{
								String: addr.City.String,
								Valid:  addr.City.Valid,
							},
							State: addr.State.String,
						})
					}
				}
			}
		}
	}

	return all, nil
}
