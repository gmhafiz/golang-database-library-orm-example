package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"godb/db"
	"godb/respond/message"
)

// Transaction pattern adapted from https://go.dev/doc/database/execute-transactions
func (r *repository) Transaction(ctx context.Context, userID int64, req *db.UserUpdateRequest) (*db.UserResponse, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}

	// if transaction is done, it will return right away and not do a round-trip
	// to the database thanks to a check in `atomic.CompareAndSwapInt32(&tx.done, 0, 1)`
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;")
	if err != nil {
		return nil, fmt.Errorf("error setting isolation level: %w", err)
	}

	var u db.UserDB

	err = tx.GetContext(ctx, &u, get, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return &db.UserResponse{}, &db.Err{Msg: message.ErrInternalError.Error()}
	}

	u.FirstName = req.FirstName
	u.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  req.MiddleName != "",
	}
	u.LastName = req.LastName
	u.Email = req.Email
	u.FavouriteColour = req.FavouriteColour

	_, err = tx.ExecContext(ctx, update,
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Email,
		u.FavouriteColour,
		u.Tags,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &db.UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
		Tags:            u.Tags,
		UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (r *repository) TransactionIsolation(ctx context.Context, userID int64, req *db.UserUpdateRequest) (*db.UserResponse, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}

	// if transaction is done, it will return right away and not do a round-trip
	// to the database thanks to a check in `atomic.CompareAndSwapInt32(&tx.done, 0, 1)`
	defer tx.Rollback()

	var u db.UserDB

	// cannot use sqlx's GetContext() that made scanning easier.
	rows := tx.QueryRowContext(ctx, get, userID)
	err = rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password, &u.FavouriteColour, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return &db.UserResponse{}, &db.Err{Msg: message.ErrInternalError.Error()}
	}

	u.FirstName = req.FirstName
	u.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  req.MiddleName != "",
	}
	u.LastName = req.LastName
	u.Email = req.Email
	u.FavouriteColour = req.FavouriteColour

	_, err = tx.ExecContext(ctx, update,
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Email,
		u.FavouriteColour,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &db.UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
		Tags:            u.Tags,
		UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
	}, nil
}
