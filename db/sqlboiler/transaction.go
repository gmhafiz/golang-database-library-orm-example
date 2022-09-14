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
