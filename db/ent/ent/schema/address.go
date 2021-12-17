package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("line_1"),
		field.String("line_2").Nillable().Optional(),
		field.Uint("postcode"),
		field.String("state"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("country", Country.Type).Ref("addresses").Unique(),
		edge.To("users", User.Type),
	}
}
