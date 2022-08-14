package db

import "database/sql"

type AddressForCountry struct {
	ID       uint   `json:"id,omitempty"`
	Line1    string `json:"line_1,omitempty"`
	Line2    string `json:"line_2,omitempty"`
	Postcode int32  `json:"postcode,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
}

type AddressDB struct {
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

type UserDB struct {
	ID              uint           `db:"id"`
	FirstName       string         `db:"first_name"`
	MiddleName      sql.NullString `db:"middle_name"`
	LastName        string         `db:"last_name"`
	Email           string         `db:"email"`
	Password        string         `db:"password"`
	FavouriteColour string         `db:"favourite_colour"`
}

//type ValidColours string
//
//const (
//	ValidColoursRed   ValidColours = "red"
//	ValidColoursGreen ValidColours = "green"
//	ValidColoursBlue  ValidColours = "blue"
//)
