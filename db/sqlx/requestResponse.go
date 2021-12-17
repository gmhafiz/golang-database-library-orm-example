package sqlx

import (
	"database/sql"

	"godb/db/ent/ent/gen"
)

type UserRequest struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UserUpdateRequest struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}

type UserResponse struct {
	ID         uint   `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}

type UserResponseWithAddress struct {
	ID         uint              `json:"id,omitempty"`
	FirstName  string            `json:"first_name"`
	MiddleName string            `json:"middle_name,omitempty"`
	LastName   string            `json:"last_name"`
	Email      string            `json:"email"`
	Address    AddressForCountry `json:"address"`
}

type UserResponseWithAddressesSqlx struct {
	ID         uint                `json:"id,omitempty"`
	FirstName  string              `json:"first_name"`
	MiddleName string              `json:"middle_name,omitempty"`
	LastName   string              `json:"last_name"`
	Email      string              `json:"email"`
	Address    []AddressForCountry `json:"address"`
}

type UserResponseWithAddresses struct {
	ID         uint           `json:"id,omitempty" `
	FirstName  string         `json:"first_name"`
	MiddleName *string        `json:"middle_name,omitempty"`
	LastName   string         `json:"last_name"`
	Email      string         `json:"email"`
	Address    []*gen.Address `json:"address"`
}

type AddressResponse struct {
	ID        uint           `json:"id,omitempty"`
	Line1     string         `json:"line_1,omitempty"`
	Line2     sql.NullString `json:"line_2"`
	Postcode  sql.NullInt32  `json:"postcode"`
	City      sql.NullString `json:"city"`
	State     sql.NullString `json:"state"`
	CountryID sql.NullInt64  `json:"country_id,omitempty"`
}

type CountryResponseWithAddress struct {
	Id        int                 `json:"id"`
	Code      string              `json:"code"`
	Name      string              `json:"name"`
	Addresses []AddressForCountry `json:"address"`
}

type AddressForCountry struct {
	ID       uint   `json:"id,omitempty"`
	Line1    string `json:"line_1,omitempty"`
	Line2    string `json:"line_2"`
	Postcode int32  `json:"postcode"`
	City     string `json:"city"`
	State    string `json:"state"`
}
