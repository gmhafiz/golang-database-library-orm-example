package ent

import (
	"context"

	"godb/db/ent/ent/gen"
)

func (r *database) Countries(ctx context.Context) ([]*gen.Country, error) {
	return r.db.Country.
		Query().
		WithAddresses().
		All(ctx)
}
