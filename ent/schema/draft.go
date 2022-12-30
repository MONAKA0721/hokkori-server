package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Draft holds the schema definition for the Draft entity.
type Draft struct {
	ent.Schema
}

// Fields of the Draft.
func (Draft) Fields() []ent.Field {
	return []ent.Field{
		field.Text("praise_title"),
		field.Text("letter_title"),
		field.Text("praise_content"),
		field.Text("letter_content"),
		field.Bool("praise_spoiled"),
		field.Bool("letter_spoiled"),
	}
}

// Edges of the Draft.
func (Draft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("drafts").Unique().Required(),
		edge.To("hashtags", Hashtag.Type),
		edge.From("work", Work.Type).Ref("drafts").Unique(),
		edge.To("category", Category.Type).Unique(),
	}
}

func (Draft) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Draft) Mixin() []ent.Mixin {
	return customTimeMixin
}
