package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
)

var customTimeMixin = []ent.Mixin{
	mixin.AnnotateFields(
		mixin.CreateTime{},
		entgql.OrderField("CREATE_TIME"),
	),
	mixin.AnnotateFields(
		mixin.UpdateTime{},
		entgql.OrderField("UPDATE_TIME"),
	),
}
