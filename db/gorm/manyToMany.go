package gorm

import (
	"context"
	"fmt"
)

func (r *repo) ListM2M(ctx context.Context) ([]*User, error) {
	var users []*User

	err := r.db.WithContext(ctx).
		Preload("Addresses").
		Find(&users).
		Select("*").
		Limit(30).
		Error
	if err != nil {
		return nil, fmt.Errorf("error loading user with addresses: %w", err)
	}

	return users, nil
}
