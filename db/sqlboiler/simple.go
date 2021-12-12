package sqlboiler

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"godb/db/sqlboiler/models"
	sqlx2 "godb/db/sqlx"
)

type database struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: db,
	}
}

func (r *database) Create(ctx context.Context, request sqlx2.UserRequest, hash string) (*models.User, error) {
	user := &models.User{
		FirstName: request.FirstName,
		MiddleName: null.String{
			String: request.MiddleName,
			Valid:  true,
		},
		LastName: request.LastName,
		Email:    request.Email,
		Password: hash,
	}

	return user, user.Insert(ctx, r.db, boil.Infer())
}

func (r *database) List(ctx context.Context) ([]*sqlx2.UserResponse, error) {
	users, err := models.Users().All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("error getting users")
	}

	var userResponse []*sqlx2.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, &sqlx2.UserResponse{
			ID:         uint(user.ID),
			FirstName:  user.FirstName,
			MiddleName: user.MiddleName.String,
			LastName:   user.LastName,
			Email:      user.Email,
		})
	}
	return userResponse, nil
}

func (r *database) Get(ctx context.Context, userID int64) (*models.User, error) {
	user, err := models.FindUser(ctx, r.db, userID)
	if err != nil {
		return nil, fmt.Errorf("db scanning error")
	}

	return user, nil
}

func (r *database) Update(ctx context.Context, req sqlx2.UserUpdateRequest) (*models.User, error) {
	user := &models.User{
		ID:        int64(req.ID),
		FirstName: req.FirstName,
		MiddleName: null.String{
			String: req.MiddleName,
			Valid:  true,
		},
		LastName: req.LastName,
		Email:    req.Email,
	}

	_, err := user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("error updating user")
	}

	return user, nil
}

func (r *database) Delete(ctx context.Context, userID int64) error {
	u, err := r.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user")
	}

	_, err = u.Delete(ctx, r.db)
	if err != nil {
		return fmt.Errorf("error deleting user")
	}

	return nil
}
