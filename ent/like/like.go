// Code generated by ent, DO NOT EDIT.

package like

import (
	"time"
)

const (
	// Label holds the string label denoting the like type in the database.
	Label = "like"
	// FieldLikedAt holds the string denoting the liked_at field in the database.
	FieldLikedAt = "liked_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// PostFieldID holds the string denoting the ID field of the Post.
	PostFieldID = "id"
	// Table holds the table name of the like in the database.
	Table = "likes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "likes"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_id"
)

// Columns holds all SQL columns for like fields.
var Columns = []string{
	FieldLikedAt,
	FieldUserID,
	FieldPostID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultLikedAt holds the default value on creation for the "liked_at" field.
	DefaultLikedAt func() time.Time
)