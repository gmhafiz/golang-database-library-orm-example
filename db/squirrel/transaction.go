package squirrel

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"

	"godb/db"
	"godb/respond/message"
)

func (r repository) Transaction(ctx context.Context, id int64, req *db.UserUpdateRequest) (*db.UserResponse, error) {
	tx, err := r.forTx.Beginx()
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}

	defer tx.Rollback()

	rows := r.db.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		RunWith(tx).
		QueryRowContext(ctx)

	var u db.UserDB
	err = rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password, &u.FavouriteColour, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		return nil, err
	}

	u.FirstName = req.FirstName
	u.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  req.MiddleName != "",
	}
	u.LastName = req.LastName
	u.Email = req.Email
	u.FavouriteColour = req.FavouriteColour

	_, err = r.db.Update("users").
		Set("first_name", u.FirstName).
		Set("middle_name", u.MiddleName).
		Set("last_name", u.LastName).
		Set("email", u.Email).
		Set("favourite_colour", u.FavouriteColour).
		Where(sq.Eq{"id": id}).
		RunWith(tx).
		ExecContext(ctx)
	if err != nil {
		return nil, nil
	}

	rows = r.db.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		RunWith(tx).
		QueryRowContext(ctx)

	err = rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password, &u.FavouriteColour, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		return nil, err
	}

	_ = tx.Commit()

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
