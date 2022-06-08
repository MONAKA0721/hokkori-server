package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Text("title").NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Enum("type").Values("letter", "praise"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
