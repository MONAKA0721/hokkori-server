// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MONAKA0721/hokkori/ent/bookmark"
	"github.com/MONAKA0721/hokkori/ent/category"
	"github.com/MONAKA0721/hokkori/ent/draft"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
	"github.com/MONAKA0721/hokkori/ent/like"
	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/ent/schema"
	"github.com/MONAKA0721/hokkori/ent/user"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookmarkFields := schema.Bookmark{}.Fields()
	_ = bookmarkFields
	// bookmarkDescBookmarkedAt is the schema descriptor for bookmarked_at field.
	bookmarkDescBookmarkedAt := bookmarkFields[0].Descriptor()
	// bookmark.DefaultBookmarkedAt holds the default value on creation for the bookmarked_at field.
	bookmark.DefaultBookmarkedAt = bookmarkDescBookmarkedAt.Default.(func() time.Time)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	draftMixin := schema.Draft{}.Mixin()
	draftMixinFields0 := draftMixin[0].Fields()
	_ = draftMixinFields0
	draftMixinFields1 := draftMixin[1].Fields()
	_ = draftMixinFields1
	draftFields := schema.Draft{}.Fields()
	_ = draftFields
	// draftDescCreateTime is the schema descriptor for create_time field.
	draftDescCreateTime := draftMixinFields0[0].Descriptor()
	// draft.DefaultCreateTime holds the default value on creation for the create_time field.
	draft.DefaultCreateTime = draftDescCreateTime.Default.(func() time.Time)
	// draftDescUpdateTime is the schema descriptor for update_time field.
	draftDescUpdateTime := draftMixinFields1[0].Descriptor()
	// draft.DefaultUpdateTime holds the default value on creation for the update_time field.
	draft.DefaultUpdateTime = draftDescUpdateTime.Default.(func() time.Time)
	// draft.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	draft.UpdateDefaultUpdateTime = draftDescUpdateTime.UpdateDefault.(func() time.Time)
	hashtagFields := schema.Hashtag{}.Fields()
	_ = hashtagFields
	// hashtagDescTitle is the schema descriptor for title field.
	hashtagDescTitle := hashtagFields[0].Descriptor()
	// hashtag.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	hashtag.TitleValidator = hashtagDescTitle.Validators[0].(func(string) error)
	likeFields := schema.Like{}.Fields()
	_ = likeFields
	// likeDescLikedAt is the schema descriptor for liked_at field.
	likeDescLikedAt := likeFields[0].Descriptor()
	// like.DefaultLikedAt holds the default value on creation for the liked_at field.
	like.DefaultLikedAt = likeDescLikedAt.Default.(func() time.Time)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postMixinFields1 := postMixin[1].Fields()
	_ = postMixinFields1
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreateTime is the schema descriptor for create_time field.
	postDescCreateTime := postMixinFields0[0].Descriptor()
	// post.DefaultCreateTime holds the default value on creation for the create_time field.
	post.DefaultCreateTime = postDescCreateTime.Default.(func() time.Time)
	// postDescUpdateTime is the schema descriptor for update_time field.
	postDescUpdateTime := postMixinFields1[0].Descriptor()
	// post.DefaultUpdateTime holds the default value on creation for the update_time field.
	post.DefaultUpdateTime = postDescUpdateTime.Default.(func() time.Time)
	// post.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	post.UpdateDefaultUpdateTime = postDescUpdateTime.UpdateDefault.(func() time.Time)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[0].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[1].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescThumbnail is the schema descriptor for thumbnail field.
	postDescThumbnail := postFields[4].Descriptor()
	// post.ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	post.ThumbnailValidator = postDescThumbnail.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	workFields := schema.Work{}.Fields()
	_ = workFields
	// workDescTitle is the schema descriptor for title field.
	workDescTitle := workFields[0].Descriptor()
	// work.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	work.TitleValidator = workDescTitle.Validators[0].(func(string) error)
	// workDescThumbnail is the schema descriptor for thumbnail field.
	workDescThumbnail := workFields[1].Descriptor()
	// work.ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	work.ThumbnailValidator = workDescThumbnail.Validators[0].(func(string) error)
}
