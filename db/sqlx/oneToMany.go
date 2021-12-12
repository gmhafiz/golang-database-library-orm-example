package sqlx

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

const (
	GetWithAddresses  = "SELECT c.id AS country_id, c.name, c.code, a.id AS address_id, a.line_1, a.line_2, a.postcode, a.city, a.state FROM countries c LEFT JOIN addresses a on c.id = a.country_id ORDER BY c.id"
	GetWithAddresses2 = "select row_to_json(row) from (select * from country_address) row"
)

type countryWithAddress struct {
	CountryID   string         `db:"country_id"`
	CountryName string         `db:"name"`
	CountryCode string         `db:"code"`
	AddressID   uint           `db:"address_id"`
	Line1       string         `db:"line_1"`
	Line2       sql.NullString `db:"line_2"`
	Postcode    sql.NullInt32  `db:"postcode"`
	City        sql.NullString `db:"city"`
	State       sql.NullString `db:"state"`
}

// Countries Using array_agg. Here we return an aggregated response that
// aggregates each address to its related country. It is done by creating a
// view beforehand. Then we simply query the view.
/*
	CREATE VIEW country_address as
	select c.id, c.code, c.name,
	       (
	           select array_to_json(array_agg(row_to_json(addresslist.*))) as array_to_json
	           from (
	                    select a.*
	                    from addresses a
	                    where c.id = a.country_id
	                ) addresslist) as addresses
	from countries AS c;
*/
func (r *database) Countries(ctx context.Context) ([]*CountryResponseWithAddress, error) {
	var resp []*CountryResponseWithAddress

	rows2, err := r.db.QueryContext(ctx, GetWithAddresses2)
	if err != nil {
		return nil, fmt.Errorf(`{"message": "db error"}`)
	}
	defer rows2.Close()

	for rows2.Next() {
		var i CountryResponseWithAddress
		err = rows2.Scan(&i)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &i)
	}

	return resp, nil
}

// Scan When scanning the result, we are actually getting an array of uint8.
// These json payload is then unmarshalled into our custom struct.
func (m *CountryResponseWithAddress) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}
