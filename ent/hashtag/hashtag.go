// Code generated by ent, DO NOT EDIT.

package hashtag

const (
	// Label holds the string label denoting the hashtag type in the database.
	Label = "hashtag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeDrafts holds the string denoting the drafts edge name in mutations.
	EdgeDrafts = "drafts"
	// Table holds the table name of the hashtag in the database.
	Table = "hashtags"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "post_hashtags"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// DraftsTable is the table that holds the drafts relation/edge. The primary key declared below.
	DraftsTable = "draft_hashtags"
	// DraftsInverseTable is the table name for the Draft entity.
	// It exists in this package in order to avoid circular dependency with the "draft" package.
	DraftsInverseTable = "drafts"
)

// Columns holds all SQL columns for hashtag fields.
var Columns = []string{
	FieldID,
	FieldTitle,
}

var (
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"post_id", "hashtag_id"}
	// DraftsPrimaryKey and DraftsColumn2 are the table columns denoting the
	// primary key for the drafts relation (M2M).
	DraftsPrimaryKey = []string{"draft_id", "hashtag_id"}
)

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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)
