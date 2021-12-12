package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
//func (User) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		entsql.Annotation{Table: "entUsers"},
//	}
//}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("first_name"),
		field.String("middle_name").Nillable().Optional(),
		field.String("last_name"),
		field.String("email").Unique(),
		field.String("password").Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("addresses", Address.Type).Ref("users"),
	}
}
