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
