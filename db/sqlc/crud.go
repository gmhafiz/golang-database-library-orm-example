package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"godb/db"
	"godb/db/sqlc/mariadb"
	"godb/db/sqlc/pg"
	"godb/respond/message"
)

type database struct {
	db      *pg.Queries
	mariaDB *mariadb.Queries

	tx   pg.DBTX
	sqlx *sqlx.DB

	dbType string
}

func NewRepo(db *sqlx.DB, dbType string) *database {
	return &database{
		db:      pg.New(db),
		mariaDB: mariadb.New(db),
		sqlx:    db,
		dbType:  dbType,
	}
}

func (r *database) Create(ctx context.Context, request *db.CreateUserRequest, hash string) (*pg.User, error) {
	u, err := r.db.CreateUser(ctx, pg.CreateUserParams{
		FirstName: request.FirstName,
		MiddleName: sql.NullString{
			String: request.MiddleName,
			Valid:  request.MiddleName != "",
		},
		LastName:        request.LastName,
		Email:           request.Email,
		FavouriteColour: pg.ValidColours(request.FavouriteColour),
		Password:        hash,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return &pg.User{}, &db.Err{
					Msg:    message.ErrUniqueKeyViolation.Error(),
					Status: http.StatusBadRequest,
				}
			default:
				return &pg.User{}, &db.Err{
					Msg:    message.ErrDefault.Error(),
					Status: http.StatusInternalServerError,
				}
			}
		}

		return nil, err
	}

	return &u, nil
}

func (r *database) List(ctx context.Context, f *db.Filter) (l []db.UserResponse, err error) {
	if f.FirstName != "" || f.Email != "" || f.FavouriteColour != "" {
		return r.ListFilterByColumn(ctx, f)
	}

	if len(f.LastNames) > 0 {
		return r.ListFilterWhereIn(ctx, f)
	}

	if len(f.Base.Sort) > 0 {
		return r.ListFilterSort(ctx, f)
	}
	if f.Base.Page > 1 {
		return r.ListFilterPagination(ctx, f)
	}

	users, err := r.db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, row := range users {
		l = append(l, db.UserResponse{
			ID:              uint(row.ID),
			FirstName:       row.FirstName,
			MiddleName:      row.MiddleName.String,
			LastName:        row.LastName,
			Email:           row.Email,
			FavouriteColour: string(row.FavouriteColour),
			Tags:            row.Tags,
			UpdatedAt:       row.UpdatedAt.Format(time.RFC3339),
		})
	}

	return l, nil
}

func (r *database) Get(ctx context.Context, userID int64) (*db.UserResponse, error) {
	res, err := r.db.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.UserResponse{}, &db.Err{Msg: message.ErrRecordNotFound.Error(), Status: http.StatusNotFound}
		}
		log.Println(err)
		return &db.UserResponse{}, &db.Err{Msg: message.ErrInternalError.Error(), Status: http.StatusInternalServerError}
	}

	return &db.UserResponse{
		ID:              uint(res.ID),
		FirstName:       res.FirstName,
		MiddleName:      res.MiddleName.String,
		LastName:        res.LastName,
		Email:           res.Email,
		FavouriteColour: string(res.FavouriteColour),
		Tags:            res.Tags,
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (r *database) Update(ctx context.Context, userID int64, f *db.Filter, req *db.UserUpdateRequest) (*pg.GetUserRow, error) {
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

	err = r.db.UpdateUser(ctx, pg.UpdateUserParams{
		FirstName: currUser.FirstName,
		MiddleName: sql.NullString{
			String: currUser.MiddleName,
			Valid:  currUser.MiddleName != "",
		},
		LastName:        currUser.LastName,
		Email:           currUser.Email,
		FavouriteColour: pg.ValidColours(currUser.FavouriteColour),
		ID:              userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error updating the user: %w", err)
	}

	u, err := r.db.GetUser(ctx, userID)
	if err != nil {
		return nil, errors.New("error getting a user")
	}

	return &u, nil
}

func (r *database) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteUser(ctx, id)
}

func (r *database) ListFilterWhereIn(ctx context.Context, f *db.Filter) (result []db.UserResponse, err error) {
	switch r.dbType {
	case "postgres", "postgresql", "psql", "pgsql", "pgx":
		users, err := r.db.SelectWhereInLastNames(ctx, f.LastNames)
		if err != nil {
			return nil, errors.New("error getting users")
		}

		for _, val := range users {
			result = append(result, db.UserResponse{
				ID:              uint(val.ID),
				FirstName:       val.FirstName,
				MiddleName:      val.MiddleName.String,
				LastName:        val.LastName,
				Email:           val.Email,
				FavouriteColour: string(val.FavouriteColour),
				Tags:            val.Tags,
				UpdatedAt:       val.UpdatedAt.Format(time.RFC3339),
			})
		}
	case "mysql", "mariadb":
		// No native support for mysql/mariadb :(
		// See https://github.com/kyleconroy/sqlc/issues/695

		// So use sqlx.In(), or try this:
		in, err := r.mariaDB.SelectWhereInLastNamesIn(ctx, f.LastNames)
		if err != nil {
			return nil, err
		}
		fmt.Println(in)

		// Or the long way:
		mysqlQuery := fmt.Sprintf(`SELECT * FROM users WHERE last_name IN (?%v);`, strings.Repeat(",?", len(f.LastNames)-1))

		args := lo.Map(f.LastNames, func(t string, _ int) any {
			return t
		})

		rows, err := r.sqlx.QueryContext(ctx, mysqlQuery, args...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var item db.UserDB
			err := rows.Scan(
				&item.ID,
				&item.FirstName,
				&item.MiddleName,
				&item.LastName,
				&item.Email,
				&item.Password,
				&item.FavouriteColour,
				&item.Tags,
			)
			if err != nil {
				return nil, err
			}
			result = append(result, db.UserResponse{
				ID:              item.ID,
				FirstName:       item.FirstName,
				MiddleName:      item.MiddleName.String,
				LastName:        item.LastName,
				Email:           item.Email,
				FavouriteColour: item.FavouriteColour,
				Tags:            item.Tags,
				UpdatedAt:       item.UpdatedAt.Format(time.RFC3339),
			})
		}
	}

	return result, nil
}
