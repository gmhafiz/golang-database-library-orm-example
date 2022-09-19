package ent

import (
	"context"
	"fmt"

	"godb/db"
	"godb/db/ent/ent/gen"
)

func (r *database) Transaction(ctx context.Context, id int64, req *db.UserUpdateRequest) (*gen.User, error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}
	defer tx.Rollback()

	user, err := tx.User.UpdateOneID(uint(id)).
		SetFirstName(req.FirstName).
		SetNillableMiddleName(&req.MiddleName).
		SetLastName(req.LastName).
		SetEmail(req.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func WithTx(ctx context.Context, client *gen.Client, fn func(tx *gen.Tx) error) error {
	tx, err := client.Tx(ctx)

	// Or the following by specifying the isolation level.
	//tx, err := client.BeginTx(ctx, &sql.TxOptions{
	//	Isolation: sql.LevelRepeatableRead,
	//	ReadOnly:  false,
	//})

	if err != nil {
		return err
	}

	_ = tx.Rollback()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func AnotherTransactionPattern(ctx context.Context, client *gen.Client, userID int64, req *db.UserUpdateRequest) (*gen.User, error) {
	u := &gen.User{}

	if err := WithTx(ctx, client, func(tx *gen.Tx) error {

		user, err := tx.User.UpdateOneID(uint(userID)).
			SetFirstName(req.FirstName).
			SetNillableMiddleName(&req.MiddleName).
			SetLastName(req.LastName).
			SetEmail(req.Email).
			Save(ctx)
		if err != nil {
			return err
		}

		u = user

		return nil

	}); err != nil {
		return nil, err
	}

	return u, nil
}
