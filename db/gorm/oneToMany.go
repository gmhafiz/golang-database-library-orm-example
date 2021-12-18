package gorm

import (
	"context"
	"fmt"
)

func (r *repo) Countries(ctx context.Context) ([]*Country, error) {
	var countries []*Country

	err := r.db.WithContext(ctx).
		Preload("Address").
		Limit(30).
		Find(&countries).
		Select("*").
		Error
	if err != nil {
		return nil, fmt.Errorf("error loading countries: %w", err)
	}

	return countries, nil
}
