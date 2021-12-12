package sqlc

import (
	"context"
	"encoding/json"
)

func (r *database) Countries(ctx context.Context) ([]json.RawMessage, error) {
	return r.db.CountriesWithAddressAggregate(ctx)
}
