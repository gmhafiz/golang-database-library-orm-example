package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"godb/db"
	"godb/respond/message"
)

const (
	insert      = "INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, middle_name, last_name, email, favourite_colour, tags, updated_at"
	list        = "SELECT * FROM users ORDER BY id LIMIT 30 OFFSET 0;"
	get         = "SELECT * FROM users WHERE id = $1;"
	update      = "UPDATE users set first_name=$1, middle_name=$2, last_name=$3, email=$4, favourite_colour=$5 WHERE id=$6;"
	updateNamed = "UPDATE users set first_name=:first_name, middle_name=:middle_name, last_name=:last_name, email=:email, favourite_colour=:favourite_colour WHERE id=:id;"
	delete      = "DELETE FROM users where id=$1"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, request *db.CreateUserRequest, hash string) (*db.UserDB, error) {
	var u db.UserDB
	err := r.db.QueryRowContext(ctx, insert,
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
		&u.Tags,
		&u.UpdatedAt,
	)
	if err != nil {
		log.Printf("sqlx.Create: %v\n", err)
		pqErr := err.(*pq.Error)
		return nil, &db.Err{Msg: fmt.Errorf("%s", pqErr.Detail).Error()}
	}

	return &u, nil
}

const preparedStatement = false

func (r *repository) List(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
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

	if preparedStatement {
		return withPrepared(ctx, r)
	}

	rows, err := r.db.QueryxContext(ctx, list)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user records")
	}

	for rows.Next() {
		var u db.UserDB
		err = rows.StructScan(&u)
		if err != nil {
			return nil, errors.New("db scanning error")
		}
		users = append(users, &db.UserResponse{
			ID:              u.ID,
			FirstName:       u.FirstName,
			MiddleName:      u.MiddleName.String,
			LastName:        u.LastName,
			Email:           u.Email,
			FavouriteColour: u.FavouriteColour,
			Tags:            u.Tags,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
		})
	}
	return users, nil
}

func (r *repository) Get(ctx context.Context, userID int64) (*db.UserResponse, error) {
	var u db.UserDB
	err := r.db.GetContext(ctx, &u, get, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return &db.UserResponse{}, &db.Err{Msg: message.ErrInternalError.Error(), Status: http.StatusInternalServerError}
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

func (r *repository) Update(ctx context.Context, f *db.Filter, userID int64, req *db.UserUpdateRequest) (*db.UserResponse, error) {
	if f.Transaction {
		return r.Transaction(ctx, userID, req)
	}
	currUser, err := r.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	currUser.FirstName = req.FirstName
	currUser.MiddleName = req.MiddleName
	currUser.LastName = req.LastName
	currUser.Email = req.Email
	currUser.FavouriteColour = req.FavouriteColour

	_, err = r.db.ExecContext(ctx, update,
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
	return r.db.ExecContext(ctx, delete, userID)
}

func (r *repository) ListFilterWhereIn(ctx context.Context, f *db.Filter) (users []*db.UserResponse, err error) {
	query, args, err := sqlx.In("SELECT * FROM users WHERE last_name IN (?)", f.LastNames)
	if err != nil {
		return nil, fmt.Errorf("error creating query: %w", err)
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query) // no need this for mysql as it defaults to question mark (?)

	var dbScan []*db.UserDB

	err = r.db.SelectContext(ctx, &dbScan, query, args...)
	if err != nil {
		return nil, err
	}

	for _, val := range dbScan {
		users = append(users, &db.UserResponse{
			ID:              val.ID,
			FirstName:       val.FirstName,
			MiddleName:      val.MiddleName.String,
			LastName:        val.LastName,
			Email:           val.Email,
			FavouriteColour: val.FavouriteColour,
			Tags:            val.Tags,
			UpdatedAt:       val.UpdatedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (r *repository) Array(ctx context.Context, userID int64) ([]string, error) {
	selectQuery := "SELECT tags FROM users WHERE users.id = $1"

	var values []string

	err := r.db.QueryRowContext(ctx, selectQuery, userID).Scan(pq.Array(&values))
	if err != nil {
		return nil, err
	}

	return values, nil
}

func withPrepared(ctx context.Context, r *repository) (users []*db.UserResponse, err error) {
	stmt, err := r.db.PrepareContext(ctx, list)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
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
			Tags:            u.Tags,
			UpdatedAt:       u.UpdatedAt.Format(time.RFC3339),
		})
	}
	return users, nil
}
