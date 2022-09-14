package gorm

import (
	"context"

	"gorm.io/gorm"

	"godb/db"
)

func (r *repo) Transaction(ctx context.Context, id int64, req *db.UserUpdateRequest) (*User, error) {
	var u User

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.WithContext(ctx).First(&u).Error; err != nil {
			return err
		}

		u.FirstName = req.FirstName
		u.MiddleName = req.MiddleName
		u.LastName = req.LastName
		u.Email = req.Email
		u.FavouriteColour = req.FavouriteColour

		if err := tx.WithContext(ctx).Save(&u).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &u, nil
}
