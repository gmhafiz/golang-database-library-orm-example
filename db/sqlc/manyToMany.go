package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"godb/db"
	"godb/db/sqlc/pg"
	"godb/respond/message"
)

func (r *database) ListM2M(ctx context.Context) ([]*db.UserResponseWithAddressesSqlx, error) {
	users, err := r.db.SelectUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error selecting users")
	}

	var all []*db.UserResponseWithAddressesSqlx
	for _, u := range users {
		all = append(all, &db.UserResponseWithAddressesSqlx{
			ID:         uint(u.ID),
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
			Tags:       u.Tags,
			Address:    []*db.AddressForCountry{},
			UpdatedAt:  u.UpdatedAt.Format(time.RFC3339),
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
						user.Address = append(user.Address, &db.AddressForCountry{
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

func (r *database) ListM2MOneQuery(ctx context.Context) ([]*db.UserResponseWithAddressesSqlxSingleQuery, error) {
	dbResponse, err := r.db.ListM2MOneQuery(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return nil, &db.Err{Msg: message.ErrInternalError.Error(), Status: http.StatusInternalServerError}
	}

	resp := make([]*db.UserResponseWithAddressesSqlxSingleQuery, 0)

	for _, dbRow := range dbResponse {
		row := &db.UserResponseWithAddressesSqlxSingleQuery{
			ID:              uint(dbRow.ID),
			FirstName:       dbRow.FirstName,
			MiddleName:      dbRow.MiddleName.String,
			LastName:        dbRow.LastName,
			Email:           dbRow.Email,
			FavouriteColour: string(dbRow.FavouriteColour),
			Tags:            dbRow.Tags,
			Address:         dbRow.Addresses,
		}
		resp = append(resp, row)
	}

	return resp, nil
}

func getUserIDs(users []pg.SelectUsersRow) (ids []int32) {
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

func getAddressIDs(uas []pg.SelectUserAddressesRow) (ids []int32) {
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
