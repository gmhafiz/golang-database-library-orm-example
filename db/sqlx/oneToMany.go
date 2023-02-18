package sqlx

import (
	"context"
	"fmt"

	"godb/db"
)

const (
	GetWithAddresses  = "SELECT c.id AS country_id, c.name, c.code, a.id AS address_id, a.line_1, a.line_2, a.postcode, a.city, a.state FROM countries c LEFT JOIN addresses a on c.id = a.country_id ORDER BY c.id"
	GetWithAddresses2 = "select row_to_json(row) from (select * from country_address) row"
)

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
	                ) addresslist) as address
	from countries AS c;
*/
func (r *repository) Countries(ctx context.Context) ([]*db.CountryResponseWithAddress, error) {
	var resp []*db.CountryResponseWithAddress

	rows, err := r.db.QueryContext(ctx, GetWithAddresses2)
	if err != nil {
		return nil, fmt.Errorf(`{"message": "db error"}`)
	}
	defer rows.Close()

	for rows.Next() {
		var i db.CountryResponseWithAddress
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		if i.Addresses == nil {
			i.Addresses = make([]*db.AddressForCountry, 0)
		}
		resp = append(resp, &i)
	}

	return resp, nil
}
