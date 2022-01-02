package gorm

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"godb/db/sqlx"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, u *sqlx.UserRequest, hash string) (*User, error) {
	user := &User{
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName,
		LastName:        u.LastName,
		Email:           u.Email,
		Password:        hash,
		FavouriteColour: u.FavouriteColour,
	}

	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repo) List(ctx context.Context, f *Filter) ([]*User, error) {
	if f.Email != "" || f.FirstName != "" {
		return r.ListFilterByColumn(ctx, f)
	}
	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}
	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	var users []*User
	err := r.db.WithContext(ctx).
		Select([]string{"id", "first_name", "middle_name", "last_name", "email", "favourite_colour"}).
		Limit(int(f.Base.Limit)).
		Offset(f.Base.Offset).
		Find(&users).
		Error
	//err := r.db.WithContext(ctx).Select([]string{"id", "first_name", "middle_name", "last_name", "email"}).Find(&users, User{FirstName: "John"}).Limit(30).Error
	//err := r.db.WithContext(ctx).Find(&users, User{FirstName: "John"}).Limit(30).Error
	if err != nil {
		return nil, fmt.Errorf(`{"message": "db scanning error"}`)
	}

	return users, nil
}

func (r *repo) Get(ctx context.Context, userID int64) (*User, error) {
	var user User

	err := r.db.WithContext(ctx).First(&user, userID).Error // First() also can accept a `var user []*User`
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) Update(ctx context.Context, userID int64, req *sqlx.UserUpdateRequest) (*User, error) {
	u := &User{}
	u.ID = int(userID)
	r.db.First(&u)

	u.FirstName = req.FirstName
	u.MiddleName = req.MiddleName
	u.LastName = req.LastName
	u.Email = req.Email
	u.FavouriteColour = req.FavouriteColour
	err := r.db.WithContext(ctx).Save(&u).Error
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, userID)
}

func (r *repo) Delete(ctx context.Context, userID int64) error {
	return r.db.WithContext(ctx).Delete(&User{}, userID).Error
}
