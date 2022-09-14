package squirrel

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"godb/db"
	"godb/respond/message"
)

type repository struct {
	db    sq.StatementBuilderType
	forTx *sqlx.DB
}

func (r repository) Create(ctx context.Context, f *db.Filter, request *db.CreateUserRequest, hash string) (*db.UserDB, error) {
	var u db.UserDB

	query := r.db.Insert("users").
		Columns("first_name", "middle_name", "last_name", "email", "password", "favourite_colour").
		Values(request.FirstName, request.MiddleName, request.LastName, request.Email, hash, request.FavouriteColour).
		Suffix(`RETURNING "id", "first_name", "middle_name", "last_name", "email", "favourite_colour"`)

	err := query.
		QueryRowContext(ctx).
		Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.FavouriteColour)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r repository) List(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	if len(f.LastNames) > 0 {
		return r.ListFilterWhereIn(ctx, f)
	}

	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}
	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}

	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	rows, err := r.db.
		Select("*").
		From("users").
		Limit(uint64(f.Base.Limit)).
		Offset(uint64(f.Base.Offset)).
		OrderBy("id").
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var u db.UserDB
		err = rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.FavouriteColour,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &db.UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			UpdatedAt:       u.UpdatedAt.String(),
		})
	}

	return users, nil
}

func (r repository) Get(ctx context.Context, userID int64) (*db.UserResponse, error) {
	rows := r.db.
		Select("*").
		From("users").
		Where(sq.Eq{"id": userID}).
		QueryRowContext(ctx)

	var u db.UserDB
	err := rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password, &u.FavouriteColour, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		return nil, err
	}

	return &db.UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
		UpdatedAt:       u.UpdatedAt.String(),
	}, nil
}

func (r repository) Update(ctx context.Context, id int64, f *db.Filter, req *db.UserUpdateRequest) (*db.UserResponse, error) {
	if f.Transaction {
		return r.Transaction(ctx, id, req)
	}

	currUser, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = req.MiddleName
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = req.FavouriteColour

	_, err = r.db.Update("users").
		Set("first_name", currUser.FirstName).
		Set("middle_name", currUser.MiddleName).
		Set("last_name", currUser.LastName).
		Set("email", currUser.Email).
		Set("favourite_colour", currUser.FavouriteColour).
		Where(sq.Eq{"id": id}).
		ExecContext(ctx)
	if err != nil {
		return nil, nil
	}

	return r.Get(ctx, id)
}

func (r repository) Delete(ctx context.Context, id int64) (sql.Result, error) {
	_, err := r.db.Delete("users").
		Where(sq.Eq{"id": id}).
		ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r repository) ListFilterWhereIn(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	var dbScan []*db.UserDB

	rows, err := r.db.
		Select("*").
		From("users").
		Where(sq.Eq{"last_name": f.LastNames}).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var u db.UserDB
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.FavouriteColour,
		)
		if err != nil {
			return nil, err
		}
		dbScan = append(dbScan, &u)
	}

	for _, val := range dbScan {
		users = append(users, &db.UserResponse{
			ID:              val.ID,
			FirstName:       val.FirstName,
			MiddleName:      val.MiddleName.String,
			LastName:        val.LastName,
			Email:           val.Email,
			FavouriteColour: val.FavouriteColour,
			UpdatedAt:       val.UpdatedAt.String(),
		})
	}

	return users, nil
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{
		db: sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			RunWith(db.DB),

		forTx: db,
	}
}
