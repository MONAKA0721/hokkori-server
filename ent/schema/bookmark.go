package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

func (Bookmark) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "post_id"),
	}
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.Time("bookmarked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("post_id"),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("post_id"),
	}
}
