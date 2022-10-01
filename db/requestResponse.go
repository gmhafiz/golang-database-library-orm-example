package db

import (
	"database/sql"
	"encoding/json"
)

type CreateUserRequest struct {
	ID              uint   `json:"id"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	FavouriteColour string `json:"favourite_colour"`
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
	ID              uint   `json:"id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	FavouriteColour string `json:"favourite_color,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}

type UserResponseEnt struct {
	ID              uint    `json:"id,omitempty"`
	FirstName       string  `json:"first_name"`
	MiddleName      *string `json:"middle_name,omitempty"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	FavouriteColour string  `json:"favourite_colour"`
	UpdatedAt       string  `json:"updated_at"`
}

type UserResponseWithAddress struct {
	ID              uint              `json:"id,omitempty"`
	FirstName       string            `json:"first_name,omitempty"`
	MiddleName      string            `json:"middle_name,omitempty"`
	LastName        string            `json:"last_name,omitempty"`
	Email           string            `json:"email,omitempty"`
	FavouriteColour string            `json:"favourite_colour,omitempty"`
	Address         AddressForCountry `json:"address,omitempty"`
}

type UserResponseWithAddressesSqlx struct {
	ID              uint   `json:"id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	FavouriteColour string `json:"favourite_colour,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`

	Address []*AddressForCountry `json:"address"`
}

type UserResponseWithAddresses struct {
	ID              uint    `json:"id,omitempty" `
	FirstName       string  `json:"first_name,omitempty"`
	MiddleName      *string `json:"middle_name,omitempty"`
	LastName        string  `json:"last_name,omitempty"`
	Email           string  `json:"email,omitempty"`
	FavouriteColour string  `json:"favourite_colour,omitempty"`

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
	ID        uint           `json:"id,omitempty"`
	Line1     string         `json:"line_1,omitempty"`
	Line2     sql.NullString `json:"line_2,omitempty"`
	Postcode  sql.NullInt32  `json:"postcode,omitempty"`
	City      sql.NullString `json:"city,omitempty"`
	State     sql.NullString `json:"state,omitempty"`
	CountryID sql.NullInt64  `json:"country_id,omitempty"`
}

type CountryResponseWithAddress struct {
	ID   int    `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`

	Addresses []*AddressForCountry `json:"address"`
}

// Scan When scanning the result, we are actually getting an array of uint8.
// These json payload is then unmarshalled into our custom struct.
func (m *CountryResponseWithAddress) Scan(src interface{}) error {
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}
