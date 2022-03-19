package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Letter holds the schema definition for the Letter entity.
type Letter struct {
	ent.Schema
}

// Fields of the Letter.
func (Letter) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").NotEmpty(),
	}
}

// Edges of the Letter.
func (Letter) Edges() []ent.Edge {
	return nil
}
