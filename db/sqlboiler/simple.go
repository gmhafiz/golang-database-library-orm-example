package sqlboiler

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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
	users, err := models.Users(qm.Limit(30)).All(ctx, r.db)
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
	return models.FindUser(ctx, r.db, userID)
}

func (r *database) Update(ctx context.Context, id int64, req sqlx2.UserUpdateRequest) (*models.User, error) {
	user := &models.User{
		ID:        id,
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
		return nil, err
	}

	return user, nil
}

func (r *database) Delete(ctx context.Context, userID int64) error {
	//boil.DebugMode = true

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
