package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	sqlx2 "godb/db/sqlx"
	"godb/respond/message"
	"log"
	"net/http"
)

type database struct {
	db *Queries
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: New(db),
	}
}

func (r *database) Create(ctx context.Context, request *sqlx2.UserRequest, hash string) (*User, error) {
	u, _ := r.db.CreateUser(ctx, CreateUserParams{
		FirstName: request.FirstName,
		MiddleName: sql.NullString{
			String: request.MiddleName,
			Valid:  true,
		},
		LastName:        request.LastName,
		Email:           request.Email,
		FavouriteColour: ValidColours(request.FavouriteColour),
		Password:        hash,
	})
	err := &pgconn.PgError{
		Severity:         "ERROR",
		Code:             "23505",
		Message:          "duplicate key value violates unique constraint \"users_email_key\"",
		Detail:           "Key (email)=(john@example.com) already exists.",
		Hint:             "",
		Position:         0,
		InternalPosition: 0,
		InternalQuery:    "",
		Where:            "",
		SchemaName:       "public",
		TableName:        "users",
		ColumnName:       "",
		DataTypeName:     "",
		ConstraintName:   "users_email_key",
		File:             "nbtinsert.c",
		Line:             0,
		Routine:          "_bt_check_unique",
	}
	if err != nil {
		log.Println(err)

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return &User{}, &sqlx2.Err{
					Msg:    message.ErrUniqueKeyViolation.Error(),
					Status: http.StatusBadRequest,
				}
			default:
				return &User{}, &sqlx2.Err{
					Msg:    message.ErrDefault.Error(),
					Status: http.StatusInternalServerError,
				}
			}
		}

		//if errors.Is(err, sql.ErrNoRows) {
		//	return &User{}, &sqlx2.Err{Msg: message.ErrRecordNotFound.Error()}
		//}
		return nil, err
	}

	return &u, nil
}

func (r *database) List(ctx context.Context, f *Filter) (l []ListUsersRow, err error) {
	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}

	if len(f.LastName) > 0 {
		return r.ListFilterWhereIn(ctx, f)
	}

	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}
	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	return r.db.ListUsers(ctx)
}

func (r *database) Get(ctx context.Context, userID int64) (GetUserRow, error) {
	return r.db.GetUser(ctx, userID)
}

func (r *database) Update(ctx context.Context, userID int64, req *sqlx2.UserUpdateRequest) (*GetUserRow, error) {
	currUser, err := r.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no record found")
		}
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  true,
	}
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = ValidColours(req.FavouriteColour)

	err = r.db.UpdateUser(ctx, UpdateUserParams{
		FirstName:       currUser.FirstName,
		MiddleName:      currUser.MiddleName,
		LastName:        currUser.LastName,
		Email:           currUser.Email,
		FavouriteColour: currUser.FavouriteColour,
		ID:              userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error updating the user: %w", err)
	}

	u, err := r.db.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting a user: %w", err)
	}

	return &u, nil
}

func (r *database) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteUser(ctx, id)
}
