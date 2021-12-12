package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

// Annotations of the User.
//func (Country) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		entsql.Annotation{Table: "entCountries"},
//	}
//}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("code"),
		field.String("name"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("entAddresses", Address.Type).Ref("entCountry").Unique(),
		edge.To("addresses", Address.Type),
	}
}
