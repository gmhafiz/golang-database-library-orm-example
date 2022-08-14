package squirrel

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"

	"godb/db"
)

func (r repository) Countries(ctx context.Context) (resp []*db.CountryResponseWithAddress, err error) {
	rows, err := r.db.
		Select("*").
		From("countries").
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var countries []*db.Country
	for rows.Next() {
		var country db.Country
		err := rows.Scan(&country.ID, &country.Name, &country.Code)
		if err != nil {
			return nil, err
		}
		countries = append(countries, &country)
	}

	countryIDs := lo.Map(countries, func(t *db.Country, _ int) int {
		return t.ID
	})

	rows, err = r.db.Select("*").
		From("addresses").
		Where(sq.Eq{"id": countryIDs}).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var addresses []*db.AddressDB
	for rows.Next() {
		var address db.AddressDB
		err := rows.Scan(&address.ID, &address.Line1, &address.Line2, &address.Postcode, &address.City, &address.State, &address.CountryID)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)

	}

	for _, c := range countries {
		country := &db.CountryResponseWithAddress{
			Id:        c.ID,
			Code:      c.Code,
			Name:      c.Name,
			Addresses: make([]*db.AddressForCountry, 0),
		}
		resp = append(resp, country)
		for _, a := range addresses {
			if a.CountryID == c.ID {
				country.Addresses = append(country.Addresses, &db.AddressForCountry{
					ID:       uint(a.ID),
					Line1:    a.Line1,
					Line2:    a.Line2,
					Postcode: a.Postcode,
					City:     a.City,
					State:    a.State,
				})
			}
		}
	}

	return resp, nil
}
