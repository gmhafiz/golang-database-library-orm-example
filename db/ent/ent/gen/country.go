// Code generated by entc, DO NOT EDIT.

package gen

import (
	"fmt"
	"godb/db/ent/ent/gen/country"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Country is the model entity for the Country schema.
type Country struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges CountryEdges `json:"edges"`
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// Addresses holds the value of the addresses edge.
	Addresses []*Address `json:"addresses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AddressesOrErr returns the Addresses value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) AddressesOrErr() ([]*Address, error) {
	if e.loadedTypes[0] {
		return e.Addresses, nil
	}
	return nil, &NotLoadedError{edge: "addresses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			values[i] = new(sql.NullInt64)
		case country.FieldCode, country.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Country", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint(value.Int64)
		case country.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		case country.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		}
	}
	return nil
}

// QueryAddresses queries the "addresses" edge of the Country entity.
func (c *Country) QueryAddresses() *AddressQuery {
	return (&CountryClient{config: c.config}).QueryAddresses(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return (&CountryClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("gen: Country is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", code=")
	builder.WriteString(c.Code)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Countries is a parsable slice of Country.
type Countries []*Country

func (c Countries) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
