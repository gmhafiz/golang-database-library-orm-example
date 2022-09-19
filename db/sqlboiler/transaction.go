package sqlboiler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"godb/db"
	"godb/db/sqlboiler/models"
)

func (r *database) Transaction(ctx context.Context, id int64, req db.UserUpdateRequest) (*models.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}

	defer tx.Rollback()

	// (Optional) Set our executor with this transaction
	r.exec = tx

	// (Optional) Choose to call our refactored Get() method that uses a transaction.
	u, err := r.GetForTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("got user using GetForTransaction(): %v \n", u)

	// Otherwise, we follow the established pattern.
	user, err := models.FindUser(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no record found")
		}
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
	_, err = user.Update(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	_ = tx.Commit()

	return user, nil
}

// GetForTransaction is like Get, but executes using transaction
func (r *database) GetForTransaction(ctx context.Context, userID int64) (*models.User, error) {
	user, err := models.FindUser(ctx, r.exec, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no record found")
		}
		return nil, err
	}

	return user, nil
}

func (r *database) TransactionUsingHelper(ctx context.Context, id int64, req db.UserUpdateRequest) (*models.User, error) {
	u := &models.User{}

	err := Tx(ctx, r.db.DB, func(tx *sql.Tx) error {
		user, err := models.FindUser(ctx, tx, id)

		if err != nil {
			return err
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
		_, err = user.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}

		u = user

		return nil

	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func Tx(ctx context.Context, db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
