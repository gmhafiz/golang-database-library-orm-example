package sqlboiler

import (
	"context"
	"godb/db"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db/sqlboiler/models"
)

func (r *database) Countries(ctx context.Context) ([]*db.CountryResponseWithAddress, error) {
	countries, err := models.Countries(
		qm.Load(models.CountryRels.Addresses),
		qm.Limit(30),
	).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var all []*db.CountryResponseWithAddress
	for _, country := range countries {
		resp := &db.CountryResponseWithAddress{
			Id:        int(country.ID),
			Code:      country.Code,
			Name:      country.Name,
			Addresses: getAddress(country.R.Addresses),
		}
		all = append(all, resp)
	}

	return all, err
}

func getAddress(addresses models.AddressSlice) []*db.AddressForCountry {
	var all []*db.AddressForCountry
	for _, address := range addresses {
		all = append(all, &db.AddressForCountry{
			ID:       uint(address.ID),
			Line1:    address.Line1,
			Line2:    address.Line2.String,
			Postcode: int32(address.Postcode.Int),
			City:     address.City.String,
			State:    address.State.String,
		})
	}

	return all
}
