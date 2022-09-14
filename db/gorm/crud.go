package gorm

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"godb/db"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, u *db.CreateUserRequest, hash string) (*User, error) {
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

func (r *repo) List(ctx context.Context, f *db.Filter) ([]*User, error) {
	if f.Email != "" || f.FirstName != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}
	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}
	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}
	if len(f.LastNames) > 0 {
		return r.ListFilterWhereIn(ctx, f)
	}

	var users []*User
	err := r.db.WithContext(ctx).
		Select([]string{"id", "first_name", "middle_name", "last_name", "email", "favourite_colour"}).
		Limit(f.Base.Limit).
		Offset(f.Base.Offset).
		Order("id").
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

	err := r.db.WithContext(ctx).
		// First() also can accept a `var user []*User` which can return more than one record!
		First(&user, userID).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) Update(ctx context.Context, userID int64, f *db.Filter, req *db.UserUpdateRequest) (*User, error) {
	if f.Transaction {
		return r.Transaction(ctx, userID, req)
	}

	var u User
	u.ID = int(userID)
	if err := r.db.WithContext(ctx).First(&u).Error; err != nil {
		return nil, err
	}

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

func (r *repo) Patch(ctx context.Context, userID int64, req *db.UserPatchRequest) (*User, error) {
	var u User
	if err := copier.Copy(&u, req); err != nil {
		return nil, err
	}

	err := r.db.Debug().
		WithContext(ctx).
		Model(&u).
		Where("id", userID). // order is important. Cannot be after Updates()
		Updates(&u).
		Error
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, userID)
}

func (r *repo) Delete(ctx context.Context, userID int64) error {
	return r.db.WithContext(ctx).Delete(&User{}, userID).Error
}

func (r *repo) ListFilterWhereIn(ctx context.Context, f *db.Filter) (users []*User, err error) {
	err = r.db.WithContext(ctx).
		Where("last_name IN ?", f.LastNames).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
