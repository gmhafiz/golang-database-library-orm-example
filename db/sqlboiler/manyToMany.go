package sqlboiler

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db"
	"godb/db/sqlboiler/models"
)

func (r *database) ListM2M(ctx context.Context) (interface{}, error) {
	uas, err := models.Users(
		qm.Load(models.UserRels.Addresses),
		qm.Limit(30),
		qm.OrderBy(models.UserTableColumns.ID),
	).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var all []*db.UserResponseWithAddresses
	for _, ua := range uas {

		u := &db.UserResponseWithAddresses{
			ID:              uint(ua.ID),
			FirstName:       ua.FirstName,
			MiddleName:      &ua.MiddleName.String,
			LastName:        ua.LastName,
			Email:           ua.Email,
			FavouriteColour: ua.FavouriteColour,
			Tags:            ua.Tags,
			Address:         transform(ua.R.Addresses), // sqlboiler does not serialise child relationship
		}
		all = append(all, u)
	}

	return all, nil
}

func transform(slice models.AddressSlice) (addresses []*db.Address) {
	for _, address := range slice {
		a := &db.Address{
			ID:       uint(address.ID),
			Line1:    address.Line1,
			Line2:    &address.Line2.String,
			Postcode: address.Postcode.Int,
			State:    address.State.String,
		}

		addresses = append(addresses, a)
	}

	return addresses
}
