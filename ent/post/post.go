// Code generated by ent, DO NOT EDIT.

package post

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSpoiled holds the string denoting the spoiled field in the database.
	FieldSpoiled = "spoiled"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeHashtags holds the string denoting the hashtags edge name in mutations.
	EdgeHashtags = "hashtags"
	// EdgeWork holds the string denoting the work edge name in mutations.
	EdgeWork = "work"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeLikedUsers holds the string denoting the liked_users edge name in mutations.
	EdgeLikedUsers = "liked_users"
	// EdgeBookmarkedUsers holds the string denoting the bookmarked_users edge name in mutations.
	EdgeBookmarkedUsers = "bookmarked_users"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeBookmarks holds the string denoting the bookmarks edge name in mutations.
	EdgeBookmarks = "bookmarks"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "posts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_posts"
	// HashtagsTable is the table that holds the hashtags relation/edge. The primary key declared below.
	HashtagsTable = "post_hashtags"
	// HashtagsInverseTable is the table name for the Hashtag entity.
	// It exists in this package in order to avoid circular dependency with the "hashtag" package.
	HashtagsInverseTable = "hashtags"
	// WorkTable is the table that holds the work relation/edge.
	WorkTable = "posts"
	// WorkInverseTable is the table name for the Work entity.
	// It exists in this package in order to avoid circular dependency with the "work" package.
	WorkInverseTable = "works"
	// WorkColumn is the table column denoting the work relation/edge.
	WorkColumn = "work_posts"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "posts"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "post_category"
	// LikedUsersTable is the table that holds the liked_users relation/edge. The primary key declared below.
	LikedUsersTable = "likes"
	// LikedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikedUsersInverseTable = "users"
	// BookmarkedUsersTable is the table that holds the bookmarked_users relation/edge. The primary key declared below.
	BookmarkedUsersTable = "bookmarks"
	// BookmarkedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BookmarkedUsersInverseTable = "users"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "post_id"
	// BookmarksTable is the table that holds the bookmarks relation/edge.
	BookmarksTable = "bookmarks"
	// BookmarksInverseTable is the table name for the Bookmark entity.
	// It exists in this package in order to avoid circular dependency with the "bookmark" package.
	BookmarksInverseTable = "bookmarks"
	// BookmarksColumn is the table column denoting the bookmarks relation/edge.
	BookmarksColumn = "post_id"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldContent,
	FieldType,
	FieldSpoiled,
	FieldThumbnail,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"post_category",
	"user_posts",
	"work_posts",
}

var (
	// HashtagsPrimaryKey and HashtagsColumn2 are the table columns denoting the
	// primary key for the hashtags relation (M2M).
	HashtagsPrimaryKey = []string{"post_id", "hashtag_id"}
	// LikedUsersPrimaryKey and LikedUsersColumn2 are the table columns denoting the
	// primary key for the liked_users relation (M2M).
	LikedUsersPrimaryKey = []string{"user_id", "post_id"}
	// BookmarkedUsersPrimaryKey and BookmarkedUsersColumn2 are the table columns denoting the
	// primary key for the bookmarked_users relation (M2M).
	BookmarkedUsersPrimaryKey = []string{"user_id", "post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	ThumbnailValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLetter Type = "letter"
	TypePraise Type = "praise"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLetter, TypePraise:
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
