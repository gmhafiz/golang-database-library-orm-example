package squirrel

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

type Address struct {
	ID       int    `db:"id"`
	Line1    string `db:"line_1"`
	Line2    string `db:"line_2"`
	Postcode int32  `db:"postcode"`
	City     string `db:"city"`
	State    string `db:"state"`

	CountryID int `db:"country_id"`
}

type Country struct {
	ID   int    `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}
