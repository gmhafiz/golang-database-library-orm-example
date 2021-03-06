// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "line_1", Type: field.TypeString},
		{Name: "line_2", Type: field.TypeString, Nullable: true},
		{Name: "postcode", Type: field.TypeUint},
		{Name: "state", Type: field.TypeString},
		{Name: "country_addresses", Type: field.TypeUint, Nullable: true},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "addresses_countries_addresses",
				Columns:    []*schema.Column{AddressesColumns[5]},
				RefColumns: []*schema.Column{CountriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "favourite_colour", Type: field.TypeEnum, Enums: []string{"red", "green", "blue"}, Default: "green"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// AddressUsersColumns holds the columns for the "address_users" table.
	AddressUsersColumns = []*schema.Column{
		{Name: "address_id", Type: field.TypeUint},
		{Name: "user_id", Type: field.TypeUint},
	}
	// AddressUsersTable holds the schema information for the "address_users" table.
	AddressUsersTable = &schema.Table{
		Name:       "address_users",
		Columns:    AddressUsersColumns,
		PrimaryKey: []*schema.Column{AddressUsersColumns[0], AddressUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "address_users_address_id",
				Columns:    []*schema.Column{AddressUsersColumns[0]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "address_users_user_id",
				Columns:    []*schema.Column{AddressUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		CountriesTable,
		UsersTable,
		AddressUsersTable,
	}
)

func init() {
	AddressesTable.ForeignKeys[0].RefTable = CountriesTable
	AddressUsersTable.ForeignKeys[0].RefTable = AddressesTable
	AddressUsersTable.ForeignKeys[1].RefTable = UsersTable
}
