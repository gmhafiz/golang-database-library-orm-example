package sqlx

import "database/sql"

type userDB struct {
	ID              uint           `db:"id"`
	FirstName       string         `db:"first_name"`
	MiddleName      sql.NullString `db:"middle_name"`
	LastName        string         `db:"last_name"`
	Email           string         `db:"email"`
	Password        string         `db:"password"`
	FavouriteColour string         `db:"favourite_colour"`
}
