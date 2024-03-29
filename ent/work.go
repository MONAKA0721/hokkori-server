// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// Work is the model entity for the Work schema.
type Work struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Thumbnail holds the value of the "thumbnail" field.
	Thumbnail string `json:"thumbnail,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkQuery when eager-loading is set.
	Edges WorkEdges `json:"edges"`
}

// WorkEdges holds the relations/edges for other nodes in the graph.
type WorkEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Drafts holds the value of the drafts edge.
	Drafts []*Draft `json:"drafts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]*int
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// DraftsOrErr returns the Drafts value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEdges) DraftsOrErr() ([]*Draft, error) {
	if e.loadedTypes[1] {
		return e.Drafts, nil
	}
	return nil, &NotLoadedError{edge: "drafts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Work) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case work.FieldID:
			values[i] = new(sql.NullInt64)
		case work.FieldTitle, work.FieldThumbnail:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Work", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Work fields.
func (w *Work) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case work.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case work.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				w.Title = value.String
			}
		case work.FieldThumbnail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail", values[i])
			} else if value.Valid {
				w.Thumbnail = value.String
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Work entity.
func (w *Work) QueryPosts() *PostQuery {
	return (&WorkClient{config: w.config}).QueryPosts(w)
}

// QueryDrafts queries the "drafts" edge of the Work entity.
func (w *Work) QueryDrafts() *DraftQuery {
	return (&WorkClient{config: w.config}).QueryDrafts(w)
}

// Update returns a builder for updating this Work.
// Note that you need to call Work.Unwrap() before calling this method if this Work
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Work) Update() *WorkUpdateOne {
	return (&WorkClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Work entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Work) Unwrap() *Work {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Work is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Work) String() string {
	var builder strings.Builder
	builder.WriteString("Work(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("title=")
	builder.WriteString(w.Title)
	builder.WriteString(", ")
	builder.WriteString("thumbnail=")
	builder.WriteString(w.Thumbnail)
	builder.WriteByte(')')
	return builder.String()
}

// Works is a parsable slice of Work.
type Works []*Work

func (w Works) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
