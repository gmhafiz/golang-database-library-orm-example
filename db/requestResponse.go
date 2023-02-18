package db

import (
	"database/sql"
	"encoding/json"
)

type CreateUserRequest struct {
	ID              uint     `json:"id"`
	FirstName       string   `json:"first_name"`
	MiddleName      string   `json:"middle_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	Tags            []string `json:"tags"`
	FavouriteColour string   `json:"favourite_colour"`
}

type UserUpdateRequest struct {
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	FavouriteColour string `json:"favourite_colour"`
}

type UserPatchRequest struct {
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	FavouriteColour string `json:"favourite_colour"`
}

type UserResponse struct {
	ID              uint     `json:"id"`
	FirstName       string   `json:"first_name"`
	MiddleName      string   `json:"middle_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	FavouriteColour string   `json:"favourite_color"`
	Tags            []string `json:"tags"`
	UpdatedAt       string   `json:"updated_at"`
}

type UserResponseEnt struct {
	ID              uint    `json:"id"`
	FirstName       string  `json:"first_name"`
	MiddleName      *string `json:"middle_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	FavouriteColour string  `json:"favourite_colour"`
	UpdatedAt       string  `json:"updated_at"`
}

type UserResponseWithAddress struct {
	ID              uint              `json:"id"`
	FirstName       string            `json:"first_name"`
	MiddleName      string            `json:"middle_name"`
	LastName        string            `json:"last_name"`
	Email           string            `json:"email"`
	FavouriteColour string            `json:"favourite_colour"`
	Address         AddressForCountry `json:"address"`
}

type UserResponseWithAddressesSqlx struct {
	ID              uint     `json:"id"`
	FirstName       string   `json:"first_name"`
	MiddleName      string   `json:"middle_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	FavouriteColour string   `json:"favourite_colour"`
	Tags            []string `json:"tags"`
	UpdatedAt       string   `json:"updated_at"`

	Address []*AddressForCountry `json:"address"`
}

type UserResponseWithAddresses struct {
	ID              uint     `json:"id"`
	FirstName       string   `json:"first_name"`
	MiddleName      *string  `json:"middle_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	FavouriteColour string   `json:"favourite_colour"`
	Tags            []string `json:"tags"`

	Address []*Address `json:"address"`
}

// Address is the model entity for the Address schema.
type Address struct {
	ID       uint    `json:"id,omitempty"`
	Line1    string  `json:"line_1,omitempty"`
	Line2    *string `json:"line_2,omitempty"`
	Postcode int     `json:"postcode,omitempty"`
	State    string  `json:"state,omitempty"`
}

type AddressResponse struct {
	ID        uint           `json:"id"`
	Line1     string         `json:"line_1"`
	Line2     sql.NullString `json:"line_2"`
	Postcode  sql.NullInt32  `json:"postcode"`
	City      sql.NullString `json:"city"`
	State     sql.NullString `json:"state"`
	CountryID sql.NullInt64  `json:"country_id"`
}

type CountryResponseWithAddress struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	Addresses []*AddressForCountry `json:"address"`
}

// Scan When scanning the result, we are actually getting an array of uint8.
// These json payload is then unmarshalled into our custom struct.
func (m *CountryResponseWithAddress) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}
