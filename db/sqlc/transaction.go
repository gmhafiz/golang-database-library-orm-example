package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"godb/db"
	"godb/db/sqlc/pg"
)

func (r *database) Transaction(ctx context.Context, id int64, req *db.UserUpdateRequest) (*pg.GetUserRow, error) {
	tx, err := r.sqlx.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("fails to start transaction error: %w", err)
	}

	defer tx.Rollback()

	qtx := pg.New(tx).WithTx(tx)

	currUser, err := qtx.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no record found")
		}
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = sql.NullString{
		String: req.MiddleName,
		Valid:  req.MiddleName != "",
	}
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = pg.ValidColours(req.FavouriteColour)

	err = qtx.UpdateUser(ctx, pg.UpdateUserParams{
		FirstName:       currUser.FirstName,
		MiddleName:      currUser.MiddleName,
		LastName:        currUser.LastName,
		Email:           currUser.Email,
		FavouriteColour: currUser.FavouriteColour,
	})
	if err != nil {
		return nil, fmt.Errorf("error updating the user: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &currUser, nil
}
