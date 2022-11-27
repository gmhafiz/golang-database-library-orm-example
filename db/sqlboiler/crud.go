package sqlboiler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"godb/db"
	"godb/db/sqlboiler/models"
	"godb/respond/message"
)

type database struct {
	db   *sqlx.DB
	exec boil.ContextExecutor
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: db,
	}
}

func (r *database) Create(ctx context.Context, request *db.CreateUserRequest, hash string) (*models.User, error) {
	user := &models.User{
		FirstName: request.FirstName,
		MiddleName: null.String{
			String: request.MiddleName,
			Valid:  request.MiddleName != "",
		},
		LastName: request.LastName,
		Email:    request.Email,
		Password: hash,
		FavouriteColour: null.String{
			String: request.FavouriteColour,
			Valid:  request.FavouriteColour != "",
		},
	}

	return user, user.Insert(ctx, r.db, boil.Infer())
}

func (r *database) List(ctx context.Context, f *db.Filter) ([]*db.UserResponse, error) {
	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
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

	users, err := models.Users(
		qm.Offset(f.Base.Offset),
		qm.Limit(f.Base.Limit),
		qm.OrderBy(models.UserColumns.ID),
	).
		All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("error getting users")
	}

	var userResponse []*db.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, &db.UserResponse{
			ID:              uint(user.ID),
			FirstName:       user.FirstName,
			MiddleName:      user.MiddleName.String,
			LastName:        user.LastName,
			Email:           user.Email,
			FavouriteColour: user.FavouriteColour.String,
			UpdatedAt:       user.UpdatedAt.String(),
		})
	}
	return userResponse, nil
}

func (r *database) Get(ctx context.Context, userID int64) (*models.User, error) {
	user, err := models.FindUser(ctx, r.db, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no record found")
		}
		log.Println(err)
		return &models.User{}, &db.Err{Msg: message.ErrInternalError.Error(), Status: http.StatusInternalServerError}
	}

	return user, nil
}

func (r *database) Update(ctx context.Context, id int64, f *db.Filter, req db.UserUpdateRequest) (*models.User, error) {
	if f.Transaction {
		return r.Transaction(ctx, id, req)
		//return r.TransactionUsingHelper(ctx, id, req) // (Optional) Another transaction pattern.
	}

	user, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.MiddleName = null.String{
		String: req.MiddleName,
		Valid:  req.MiddleName != "",
	}
	user.LastName = req.LastName
	user.Email = req.Email
	user.FavouriteColour = null.String{
		String: req.FavouriteColour,
		Valid:  req.FavouriteColour != "",
	}

	// Ignore number of affected rows with underscore
	_, err = user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update method that deletes password
// Do not to this. It will delete password from database!
/*
UPDATE "users" SET "first_name"=$1,"middle_name"=$2,"last_name"=$3,"email"=$4,"password"=$5 WHERE "id"=$6
[John { true} Doe john-changed@example.com  13]

*/
//func (r *database) Update(ctx context.Context, id int64, req db.UserUpdateRequest) (*models.User, error) {
//	boil.DebugMode = true
//	defer func() {
//		boil.DebugMode = false
//	}()
//	user := &models.User{
//		ID:        id,
//		FirstName: req.FirstName,
//		MiddleName: null.String{
//			String: req.MiddleName,
//			Valid:  true,
//		},
//		LastNames: req.LastNames,
//		Email:    req.Email,
//      FavouriteColour: req.FavouriteColour,
//	}
//
//	_, err := user.Update(ctx, r.db, boil.Infer())
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}

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

func (r *database) ListFilterWhereIn(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	// Accepts slice of interface, not slice of string. Not generic. So need to
	// convert each element to interface{}, or 'any' in Go v1.18
	args := lo.Map(f.LastNames, func(t string, _ int) any {
		return t
	})

	all, err := models.Users(
		//qm.WhereIn("last_name", args...), // Does not work. Needs IN operator
		//qm.WhereIn("last_name IN ($1, $2)", "Donovan", "Campbell"), // Is what we want
		qm.WhereIn("last_name IN ?", args...), // instead, just give it a `?`.
	).
		All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("error getting users")
	}

	for _, i := range all {
		users = append(users, &db.UserResponse{
			ID:              uint(i.ID),
			FirstName:       i.FirstName,
			MiddleName:      i.MiddleName.String,
			LastName:        i.LastName,
			Email:           i.Email,
			FavouriteColour: i.FavouriteColour.String,
			UpdatedAt:       i.UpdatedAt.String(),
		})
	}

	return users, nil
}
