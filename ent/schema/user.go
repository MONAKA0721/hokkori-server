package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("username").Optional(),
		field.Text("profile").Optional(),
		field.String("avatar_url").Optional(),
		field.Int("age").Optional().Comment("1:10代 2:20代 3:30代 4:40代 5:50代 6:60代以上"),
		field.Int("gender").Optional().Comment("1:男 2:女 3:選択しない"),
		field.Ints("interests").Optional().Comment("1:ポケットモンスター 2:どうぶつの森 3:スーパーマリオ 4:スプラトゥーン 5:ゼルダの伝説 6:モンスターハンター 7:ドラゴンクエスト 8:ファイナルファンタジー 9:ニーア 10:桃太郎電鉄 11:パワプロ 12:メタルギア 13:マインクラフト 14:ソニック"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("liked_posts", Post.Type).
			Through("likes", Like.Type),
		edge.To("bookmarked_posts", Post.Type).
			Through("bookmarks", Bookmark.Type),
		edge.To("following", User.Type).
			From("followers"),
		edge.To("drafts", Draft.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
