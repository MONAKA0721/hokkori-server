// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/ent/predicate"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// WorkUpdate is the builder for updating Work entities.
type WorkUpdate struct {
	config
	hooks    []Hook
	mutation *WorkMutation
}

// Where appends a list predicates to the WorkUpdate builder.
func (wu *WorkUpdate) Where(ps ...predicate.Work) *WorkUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WorkUpdate) SetTitle(s string) *WorkUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetThumbnail sets the "thumbnail" field.
func (wu *WorkUpdate) SetThumbnail(s string) *WorkUpdate {
	wu.mutation.SetThumbnail(s)
	return wu
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (wu *WorkUpdate) SetNillableThumbnail(s *string) *WorkUpdate {
	if s != nil {
		wu.SetThumbnail(*s)
	}
	return wu
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (wu *WorkUpdate) ClearThumbnail() *WorkUpdate {
	wu.mutation.ClearThumbnail()
	return wu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (wu *WorkUpdate) AddPostIDs(ids ...int) *WorkUpdate {
	wu.mutation.AddPostIDs(ids...)
	return wu
}

// AddPosts adds the "posts" edges to the Post entity.
func (wu *WorkUpdate) AddPosts(p ...*Post) *WorkUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.AddPostIDs(ids...)
}

// Mutation returns the WorkMutation object of the builder.
func (wu *WorkUpdate) Mutation() *WorkMutation {
	return wu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (wu *WorkUpdate) ClearPosts() *WorkUpdate {
	wu.mutation.ClearPosts()
	return wu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (wu *WorkUpdate) RemovePostIDs(ids ...int) *WorkUpdate {
	wu.mutation.RemovePostIDs(ids...)
	return wu
}

// RemovePosts removes "posts" edges to Post entities.
func (wu *WorkUpdate) RemovePosts(p ...*Post) *WorkUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wu.RemovePostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := work.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Work.title": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Thumbnail(); ok {
		if err := work.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail", err: fmt.Errorf(`ent: validator failed for field "Work.thumbnail": %w`, err)}
		}
	}
	return nil
}

func (wu *WorkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldTitle,
		})
	}
	if value, ok := wu.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldThumbnail,
		})
	}
	if wu.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: work.FieldThumbnail,
		})
	}
	if wu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !wu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WorkUpdateOne is the builder for updating a single Work entity.
type WorkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkMutation
}

// SetTitle sets the "title" field.
func (wuo *WorkUpdateOne) SetTitle(s string) *WorkUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetThumbnail sets the "thumbnail" field.
func (wuo *WorkUpdateOne) SetThumbnail(s string) *WorkUpdateOne {
	wuo.mutation.SetThumbnail(s)
	return wuo
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillableThumbnail(s *string) *WorkUpdateOne {
	if s != nil {
		wuo.SetThumbnail(*s)
	}
	return wuo
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (wuo *WorkUpdateOne) ClearThumbnail() *WorkUpdateOne {
	wuo.mutation.ClearThumbnail()
	return wuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (wuo *WorkUpdateOne) AddPostIDs(ids ...int) *WorkUpdateOne {
	wuo.mutation.AddPostIDs(ids...)
	return wuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (wuo *WorkUpdateOne) AddPosts(p ...*Post) *WorkUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.AddPostIDs(ids...)
}

// Mutation returns the WorkMutation object of the builder.
func (wuo *WorkUpdateOne) Mutation() *WorkMutation {
	return wuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (wuo *WorkUpdateOne) ClearPosts() *WorkUpdateOne {
	wuo.mutation.ClearPosts()
	return wuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (wuo *WorkUpdateOne) RemovePostIDs(ids ...int) *WorkUpdateOne {
	wuo.mutation.RemovePostIDs(ids...)
	return wuo
}

// RemovePosts removes "posts" edges to Post entities.
func (wuo *WorkUpdateOne) RemovePosts(p ...*Post) *WorkUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wuo.RemovePostIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkUpdateOne) Select(field string, fields ...string) *WorkUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Work entity.
func (wuo *WorkUpdateOne) Save(ctx context.Context) (*Work, error) {
	var (
		err  error
		node *Work
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Work)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WorkMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkUpdateOne) SaveX(ctx context.Context) *Work {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := work.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Work.title": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Thumbnail(); ok {
		if err := work.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail", err: fmt.Errorf(`ent: validator failed for field "Work.thumbnail": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorkUpdateOne) sqlSave(ctx context.Context) (_node *Work, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Work.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, work.FieldID)
		for _, f := range fields {
			if !work.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != work.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldTitle,
		})
	}
	if value, ok := wuo.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldThumbnail,
		})
	}
	if wuo.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: work.FieldThumbnail,
		})
	}
	if wuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !wuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   work.PostsTable,
			Columns: []string{work.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Work{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
