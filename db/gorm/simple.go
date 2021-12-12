package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"godb/db/sqlx"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, u *sqlx.UserRequest, hash string) {
	user := &User{
		FirstName: u.FirstName,
		MiddleName: sql.NullString{
			String: u.MiddleName,
			Valid:  true,
		},
		LastName: u.LastName,
		Email:    u.Email,
		Password: hash,
	}

	r.db.WithContext(ctx).Create(user)
}

func (r *repo) List(ctx context.Context) ([]*sqlx.UserResponse, error) {
	var users []User
	err := r.db.WithContext(ctx).Model(&User{}).Select("*").Scan(&users).Error
	if err != nil {
		return nil, fmt.Errorf(`{"message": "db scanning error"}`)
	}

	var userResponse []*sqlx.UserResponse
	for _, u := range users {
		userResponse = append(userResponse, &sqlx.UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}
	return userResponse, nil
}

func (r *repo) Get(ctx context.Context, userID int64) *sqlx.UserResponse {
	var user User

	r.db.WithContext(ctx).First(&user, userID)

	return &sqlx.UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName.String,
		LastName:   user.LastName,
		Email:      user.Email,
	}
}

func (r *repo) Update(userID int64, req *sqlx.UserUpdateRequest) {
	u := &User{}
	u.ID = uint(userID)
	r.db.First(&u)

	u.FirstName = req.FirstName
	u.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  true,
	}
	u.LastName = req.LastName
	u.Email = req.Email
	r.db.Save(&u)
}

func (r *repo) Delete(ctx context.Context, userID int64) {
	r.db.WithContext(ctx).Delete(&User{}, userID)
}
