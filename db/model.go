package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

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
	Password        SecretString   `db:"password"`
	FavouriteColour string         `db:"favourite_colour"`
	UpdatedAt       time.Time      `db:"updated_at"`
}

// SecretString is a custom struct that wraps string type. It has methods that
// implements various interface to help sensitive values like a password from
// being exposed to end user.
type SecretString struct {
	value string
}

func (s SecretString) String() string {
	return "***"
}

func (s SecretString) Scan(src interface{}) error {
	return nil
}

func (s SecretString) Value() (driver.Value, error) {
	return s.value, nil
}

func (s SecretString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type UserResponseWithAddressesSqlxSingleQuery struct {
	ID              uint            `json:"id,omitempty" db:"id"`
	FirstName       string          `json:"first_name,omitempty" db:"first_name"`
	MiddleName      string          `json:"middle_name,omitempty" db:"middle_name"`
	LastName        string          `json:"last_name,omitempty" db:"last_name"`
	Email           string          `json:"email,omitempty" db:"email"`
	FavouriteColour string          `json:"favourite_colour,omitempty" db:"favourite_colour"`
	UpdatedAt       string          `json:"updated_at,omitempty" db:"updated_at"`
	Address         json.RawMessage `json:"addresses" db:"addresses"`
}

// Scan When scanning the result, we are actually getting an array of uint8.
// These json payload is then unmarshalled into our custom struct.
func (m *UserResponseWithAddressesSqlxSingleQuery) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}
