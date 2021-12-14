package sqlx

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Insert = "INSERT INTO users (first_name, middle_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, middle_name, last_name, email, password"
	List   = "SELECT * FROM users;"
	Get    = "SELECT * FROM users WHERE id = $1;"
	Update = "UPDATE users set first_name=$1, middle_name=$2, last_name=$3, email=$4 WHERE id=$5;"
	Delete = "DELETE FROM users where id=$1"
)

var (
	ErrUniqueKeyViolation = fmt.Errorf("unique key violation")
	ErrDefault            = fmt.Errorf("whoops, something wrong happened")
)

type database struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *database {
	return &database{
		db: db,
	}
}

func (r *database) Create(ctx context.Context, request *UserRequest, hash string) (*userDB, error) {
	var u userDB
	err := r.db.QueryRowContext(ctx, Insert,
		request.FirstName,
		request.MiddleName,
		request.LastName,
		request.Email,
		hash,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating user record: %w", err)
	}

	return &u, nil
}

func (r *database) List(ctx context.Context) (users []*UserResponse, err error) {
	rows, err := r.db.QueryContext(ctx, List)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user records")
	}

	for rows.Next() {
		var u userDB
		err := rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.Email, &u.Password)
		if err != nil {
			return nil, fmt.Errorf("db scanning error")
		}
		users = append(users, &UserResponse{
			ID:         u.ID,
			FirstName:  u.FirstName,
			MiddleName: u.MiddleName.String,
			LastName:   u.LastName,
			Email:      u.Email,
		})
	}
	return users, nil
}

func (r *database) Get(ctx context.Context, userID int64) (*UserResponse, error) {
	var u userDB
	err := r.db.GetContext(ctx, &u, Get, userID)
	if err != nil {
		return nil, fmt.Errorf("db error")
	}

	return &UserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		MiddleName: u.MiddleName.String,
		LastName:   u.LastName,
		Email:      u.Email,
	}, nil
}

func (r *database) Update(ctx context.Context, userID int64, req *UserUpdateRequest) (sql.Result, error) {
	return r.db.ExecContext(ctx, Update,
		req.FirstName,
		req.MiddleName,
		req.LastName,
		req.Email,
		userID,
	)
}

func (r *database) Delete(ctx context.Context, userID int64) (sql.Result, error) {
	return r.db.ExecContext(ctx, Delete, userID)
}

type userDB struct {
	ID         uint           `db:"id"`
	FirstName  string         `db:"first_name"`
	MiddleName sql.NullString `db:"middle_name"`
	LastName   string         `db:"last_name"`
	Email      string         `db:"email"`
	Password   string         `db:"password"`
}
