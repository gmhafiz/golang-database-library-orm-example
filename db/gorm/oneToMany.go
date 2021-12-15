package gorm

import (
	"context"
	"fmt"
)

func (r *repo) Countries(ctx context.Context) ([]*Country, error) {
	var country []*Country

	err := r.db.WithContext(ctx).Preload("Address").Find(&country).Error
	if err != nil {
		return nil, fmt.Errorf("error loading countries: %w", err)
	}

	return country, nil
}
