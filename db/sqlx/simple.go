package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"godb/respond/message"
)

const (
	Insert      = "INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, middle_name, last_name, email, favourite_colour"
	List        = "SELECT * FROM users ORDER BY id LIMIT 30 OFFSET 0;"
	Get         = "SELECT * FROM users WHERE id = $1;"
	Update      = "UPDATE users set first_name=$1, middle_name=$2, last_name=$3, email=$4, favourite_colour=$5 WHERE id=$6;"
	UpdateNamed = "UPDATE users set first_name=:first_name, middle_name=:middle_name, last_name=:last_name, email=:email, favourite_colour=:favourite_colour WHERE id=:id;"
	Delete      = "DELETE FROM users where id=$1"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, request *UserRequest, hash string) (*userDB, error) {
	var u userDB
	err := r.db.QueryRowContext(ctx, Insert,
		request.FirstName,
		request.MiddleName,
		request.LastName,
		request.Email,
		hash,
		request.FavouriteColour,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.FavouriteColour,
	)
	if err != nil {
		log.Printf("sqlx.Create: %v\n", err)
		return nil, &Err{Msg: fmt.Errorf("error creating user record: %w", err).Error()}
	}

	return &u, nil
}

func (r *repository) List(ctx context.Context, f *Filter) (users []*UserResponse, err error) {
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

	rows, err := r.db.QueryxContext(ctx, List)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user records")
	}

	for rows.Next() {
		var u userDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
		})
	}
	return users, nil
}

func (r *repository) Get(ctx context.Context, userID int64) (*UserResponse, error) {
	var u userDB
	err := r.db.GetContext(ctx, &u, Get, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &UserResponse{}, &Err{Msg: message.ErrRecordNotFound.Error()}
		}
		log.Println(err)
		return &UserResponse{}, &Err{Msg: message.ErrInternalError.Error()}
	}

	return &UserResponse{
		ID:              u.ID,
		FirstName:       u.FirstName,
		MiddleName:      u.MiddleName.String,
		LastName:        u.LastName,
		Email:           u.Email,
		FavouriteColour: u.FavouriteColour,
	}, nil
}

func (r *repository) Update(ctx context.Context, userID int64, req *UserUpdateRequest) (*UserResponse, error) {
	currUser, err := r.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = req.MiddleName
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = req.FavouriteColour

	_, err = r.db.ExecContext(ctx, Update,
		currUser.FirstName,
		currUser.MiddleName,
		currUser.LastName,
		currUser.Email,
		currUser.FavouriteColour,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, userID)
}

func (r *repository) Delete(ctx context.Context, userID int64) (sql.Result, error) {
	return r.db.ExecContext(ctx, Delete, userID)
}

func (r *repository) ListFilterWhereIn(ctx context.Context, f *Filter) (users []*UserResponse, err error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE last_name IN (?)", f.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating query: %w", err)
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query) // no need this for mysql as it defaults to ?

	var dbScan []*userDB

	err = r.db.SelectContext(ctx, &dbScan, query, args...)
	if err != nil {
		return nil, err
	}

	for _, val := range dbScan {
		users = append(users, &UserResponse{
			ID:              val.ID,
			FirstName:       val.FirstName,
			MiddleName:      val.MiddleName.String,
			LastName:        val.LastName,
			Email:           val.Email,
			FavouriteColour: val.FavouriteColour,
		})
	}

	return users, nil
}
