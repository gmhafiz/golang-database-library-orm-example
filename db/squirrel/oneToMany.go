package squirrel

import (
	"context"
	"database/sql"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"

	"godb/db"
)

func (r repository) CountriesRawJSON(ctx context.Context) (resp []json.RawMessage, err error) {
	s := r.db.Select("* from country_address")
	rows, err := r.db.
		Select("row_to_json(row)").
		FromSelect(s, "row").
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var items []json.RawMessage
	for rows.Next() {
		var rowToJson json.RawMessage
		if err := rows.Scan(&rowToJson); err != nil {
			return nil, err
		}
		items = append(items, rowToJson)
	}

	// or

	var scanned []*Custom1ToManyStruct
	for rows.Next() {
		var rowToJson Custom1ToManyStruct
		if err := rows.Scan(&rowToJson); err != nil {
			return nil, err
		}
		scanned = append(scanned, &rowToJson)
	}

	return items, nil
}

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
			ID:        c.ID,
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

type Custom1ToManyStruct struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address []struct {
		Id        int    `json:"id"`
		Line1     string `json:"line_1"`
		Line2     string `json:"line_2"`
		Postcode  int    `json:"postcode"`
		City      string `json:"city"`
		State     string `json:"state"`
		CountryId int    `json:"country_id"`
	} `json:"address"`
}

func (m *Custom1ToManyStruct) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}
