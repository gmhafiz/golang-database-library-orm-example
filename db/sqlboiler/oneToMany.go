package sqlboiler

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db/sqlboiler/models"
	"godb/db/sqlx"
)

func (r *database) Countries(ctx context.Context) ([]*sqlx.CountryResponseWithAddress, error) {
	countries, err := models.Countries(qm.Load(models.CountryRels.Addresses)).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var all []*sqlx.CountryResponseWithAddress
	for _, country := range countries {
		resp := &sqlx.CountryResponseWithAddress{
			Id:        int(country.ID),
			Code:      country.Code,
			Name:      country.Name,
			Addresses: getAddress(country.R.Addresses),
		}
		all = append(all, resp)
	}

	return all, err
}

func getAddress(addresses models.AddressSlice) []*sqlx.AddressForCountry {
	var all []*sqlx.AddressForCountry
	for _, address := range addresses {
		all = append(all, &sqlx.AddressForCountry{
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
