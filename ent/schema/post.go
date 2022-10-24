package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
		field.Bool("spoiled"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("posts").Unique().Required(),
		edge.To("hashtags", Hashtag.Type),
		edge.From("work", Work.Type).Ref("posts").Unique().Required(),
		edge.To("category", Category.Type).Unique().Required(),
		edge.From("liked_users", User.Type).
			Ref("liked_posts").
			Through("likes", Like.Type),
		edge.From("bookmarked_users", User.Type).
			Ref("bookmarked_posts").
			Through("bookmarks", Bookmark.Type),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Post) Mixin() []ent.Mixin {
	return customTimeMixin
}
