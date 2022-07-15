package squirrel

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"

	"github.com/samber/lo"

	sqlx2 "godb/db/sqlx"
)

func (r repository) Countries(ctx context.Context) (resp []*sqlx2.CountryResponseWithAddress, err error) {
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

	var countries []*Country
	for rows.Next() {
		var country Country
		err := rows.Scan(&country.ID, &country.Name, &country.Code)
		if err != nil {
			return nil, err
		}
		countries = append(countries, &country)
	}

	countryIDs := lo.Map(countries, func(t *Country, _ int) int {
		return t.ID
	})

	rows, err = r.db.Select("*").
		From("addresses").
		Where(sq.Eq{"id": countryIDs}).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var addresses []*Address
	for rows.Next() {
		var address Address
		err := rows.Scan(&address.ID, &address.Line1, &address.Line2, &address.Postcode, &address.City, &address.State, &address.CountryID)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)

	}

	for _, c := range countries {
		country := &sqlx2.CountryResponseWithAddress{
			Id:        c.ID,
			Code:      c.Code,
			Name:      c.Name,
			Addresses: make([]*sqlx2.AddressForCountry, 0),
		}
		resp = append(resp, country)
		for _, a := range addresses {
			if a.CountryID == c.ID {
				country.Addresses = append(country.Addresses, &sqlx2.AddressForCountry{
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
