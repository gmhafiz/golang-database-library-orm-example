package sqlc

import (
	"context"
	"fmt"

	"godb/db/sqlx"
)

func (r *database) ListM2M(ctx context.Context) ([]*sqlx.UserResponseWithAddressesSqlx, error) {
	users, err := r.db.SelectUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error selecting users")
	}

	var all []*sqlx.UserResponseWithAddressesSqlx
	for _, u := range users {
		all = append(all, &sqlx.UserResponseWithAddressesSqlx{
			ID:         uint(u.ID),
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}

	userIDs := getUserIDs(users)

	uas, err := r.db.SelectUserAddresses(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("error selecting user address pivot")
	}

	addressIDs := getAddressIDs(uas)

	address, err := r.db.SelectAddress(ctx, addressIDs)
	if err != nil {
		return nil, fmt.Errorf("error selecting address")
	}

	// now we attach address to the user
	for _, u := range uas {
		for _, user := range all {
			if u.UserID.Int64 == int64(user.ID) {
				for _, addr := range address {
					if addr.ID == u.AddressID.Int64 {
						user.Address = append(user.Address, sqlx.AddressForCountry{
							ID:       uint(addr.ID),
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

func getUserIDs(users []SelectUsersRow) (ids []int32) {
	seen := make(map[int]bool)
	for _, user := range users {
		ok := seen[int(user.ID)]
		if !ok {
			ids = append(ids, int32(user.ID))
			seen[int(user.ID)] = true
		}
	}

	return ids
}

func getAddressIDs(uas []SelectUserAddressesRow) (ids []int32) {
	seen := make(map[int64]bool)
	for _, item := range uas {
		ok := seen[item.UserID.Int64]
		if !ok {
			ids = append(ids, int32(item.UserID.Int64))
			seen[item.UserID.Int64] = true
		}
	}

	return ids
}
