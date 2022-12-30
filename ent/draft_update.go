// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/category"
	"github.com/MONAKA0721/hokkori/ent/draft"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
	"github.com/MONAKA0721/hokkori/ent/predicate"
	"github.com/MONAKA0721/hokkori/ent/user"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// DraftUpdate is the builder for updating Draft entities.
type DraftUpdate struct {
	config
	hooks    []Hook
	mutation *DraftMutation
}

// Where appends a list predicates to the DraftUpdate builder.
func (du *DraftUpdate) Where(ps ...predicate.Draft) *DraftUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdateTime sets the "update_time" field.
func (du *DraftUpdate) SetUpdateTime(t time.Time) *DraftUpdate {
	du.mutation.SetUpdateTime(t)
	return du
}

// SetPraiseTitle sets the "praise_title" field.
func (du *DraftUpdate) SetPraiseTitle(s string) *DraftUpdate {
	du.mutation.SetPraiseTitle(s)
	return du
}

// SetLetterTitle sets the "letter_title" field.
func (du *DraftUpdate) SetLetterTitle(s string) *DraftUpdate {
	du.mutation.SetLetterTitle(s)
	return du
}

// SetPraiseContent sets the "praise_content" field.
func (du *DraftUpdate) SetPraiseContent(s string) *DraftUpdate {
	du.mutation.SetPraiseContent(s)
	return du
}

// SetLetterContent sets the "letter_content" field.
func (du *DraftUpdate) SetLetterContent(s string) *DraftUpdate {
	du.mutation.SetLetterContent(s)
	return du
}

// SetPraiseSpoiled sets the "praise_spoiled" field.
func (du *DraftUpdate) SetPraiseSpoiled(b bool) *DraftUpdate {
	du.mutation.SetPraiseSpoiled(b)
	return du
}

// SetLetterSpoiled sets the "letter_spoiled" field.
func (du *DraftUpdate) SetLetterSpoiled(b bool) *DraftUpdate {
	du.mutation.SetLetterSpoiled(b)
	return du
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (du *DraftUpdate) SetOwnerID(id int) *DraftUpdate {
	du.mutation.SetOwnerID(id)
	return du
}

// SetOwner sets the "owner" edge to the User entity.
func (du *DraftUpdate) SetOwner(u *User) *DraftUpdate {
	return du.SetOwnerID(u.ID)
}

// AddHashtagIDs adds the "hashtags" edge to the Hashtag entity by IDs.
func (du *DraftUpdate) AddHashtagIDs(ids ...int) *DraftUpdate {
	du.mutation.AddHashtagIDs(ids...)
	return du
}

// AddHashtags adds the "hashtags" edges to the Hashtag entity.
func (du *DraftUpdate) AddHashtags(h ...*Hashtag) *DraftUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.AddHashtagIDs(ids...)
}

// SetWorkID sets the "work" edge to the Work entity by ID.
func (du *DraftUpdate) SetWorkID(id int) *DraftUpdate {
	du.mutation.SetWorkID(id)
	return du
}

// SetNillableWorkID sets the "work" edge to the Work entity by ID if the given value is not nil.
func (du *DraftUpdate) SetNillableWorkID(id *int) *DraftUpdate {
	if id != nil {
		du = du.SetWorkID(*id)
	}
	return du
}

// SetWork sets the "work" edge to the Work entity.
func (du *DraftUpdate) SetWork(w *Work) *DraftUpdate {
	return du.SetWorkID(w.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (du *DraftUpdate) SetCategoryID(id int) *DraftUpdate {
	du.mutation.SetCategoryID(id)
	return du
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (du *DraftUpdate) SetNillableCategoryID(id *int) *DraftUpdate {
	if id != nil {
		du = du.SetCategoryID(*id)
	}
	return du
}

// SetCategory sets the "category" edge to the Category entity.
func (du *DraftUpdate) SetCategory(c *Category) *DraftUpdate {
	return du.SetCategoryID(c.ID)
}

// Mutation returns the DraftMutation object of the builder.
func (du *DraftUpdate) Mutation() *DraftMutation {
	return du.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (du *DraftUpdate) ClearOwner() *DraftUpdate {
	du.mutation.ClearOwner()
	return du
}

// ClearHashtags clears all "hashtags" edges to the Hashtag entity.
func (du *DraftUpdate) ClearHashtags() *DraftUpdate {
	du.mutation.ClearHashtags()
	return du
}

// RemoveHashtagIDs removes the "hashtags" edge to Hashtag entities by IDs.
func (du *DraftUpdate) RemoveHashtagIDs(ids ...int) *DraftUpdate {
	du.mutation.RemoveHashtagIDs(ids...)
	return du
}

// RemoveHashtags removes "hashtags" edges to Hashtag entities.
func (du *DraftUpdate) RemoveHashtags(h ...*Hashtag) *DraftUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return du.RemoveHashtagIDs(ids...)
}

// ClearWork clears the "work" edge to the Work entity.
func (du *DraftUpdate) ClearWork() *DraftUpdate {
	du.mutation.ClearWork()
	return du
}

// ClearCategory clears the "category" edge to the Category entity.
func (du *DraftUpdate) ClearCategory() *DraftUpdate {
	du.mutation.ClearCategory()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DraftUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DraftMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DraftUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DraftUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DraftUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DraftUpdate) defaults() {
	if _, ok := du.mutation.UpdateTime(); !ok {
		v := draft.UpdateDefaultUpdateTime()
		du.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DraftUpdate) check() error {
	if _, ok := du.mutation.OwnerID(); du.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Draft.owner"`)
	}
	return nil
}

func (du *DraftUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   draft.Table,
			Columns: draft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: draft.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: draft.FieldUpdateTime,
		})
	}
	if value, ok := du.mutation.PraiseTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldPraiseTitle,
		})
	}
	if value, ok := du.mutation.LetterTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldLetterTitle,
		})
	}
	if value, ok := du.mutation.PraiseContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldPraiseContent,
		})
	}
	if value, ok := du.mutation.LetterContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldLetterContent,
		})
	}
	if value, ok := du.mutation.PraiseSpoiled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: draft.FieldPraiseSpoiled,
		})
	}
	if value, ok := du.mutation.LetterSpoiled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: draft.FieldLetterSpoiled,
		})
	}
	if du.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.OwnerTable,
			Columns: []string{draft.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.OwnerTable,
			Columns: []string{draft.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.HashtagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedHashtagsIDs(); len(nodes) > 0 && !du.mutation.HashtagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.HashtagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.WorkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.WorkTable,
			Columns: []string{draft.WorkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.WorkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.WorkTable,
			Columns: []string{draft.WorkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   draft.CategoryTable,
			Columns: []string{draft.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   draft.CategoryTable,
			Columns: []string{draft.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{draft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DraftUpdateOne is the builder for updating a single Draft entity.
type DraftUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DraftMutation
}

// SetUpdateTime sets the "update_time" field.
func (duo *DraftUpdateOne) SetUpdateTime(t time.Time) *DraftUpdateOne {
	duo.mutation.SetUpdateTime(t)
	return duo
}

// SetPraiseTitle sets the "praise_title" field.
func (duo *DraftUpdateOne) SetPraiseTitle(s string) *DraftUpdateOne {
	duo.mutation.SetPraiseTitle(s)
	return duo
}

// SetLetterTitle sets the "letter_title" field.
func (duo *DraftUpdateOne) SetLetterTitle(s string) *DraftUpdateOne {
	duo.mutation.SetLetterTitle(s)
	return duo
}

// SetPraiseContent sets the "praise_content" field.
func (duo *DraftUpdateOne) SetPraiseContent(s string) *DraftUpdateOne {
	duo.mutation.SetPraiseContent(s)
	return duo
}

// SetLetterContent sets the "letter_content" field.
func (duo *DraftUpdateOne) SetLetterContent(s string) *DraftUpdateOne {
	duo.mutation.SetLetterContent(s)
	return duo
}

// SetPraiseSpoiled sets the "praise_spoiled" field.
func (duo *DraftUpdateOne) SetPraiseSpoiled(b bool) *DraftUpdateOne {
	duo.mutation.SetPraiseSpoiled(b)
	return duo
}

// SetLetterSpoiled sets the "letter_spoiled" field.
func (duo *DraftUpdateOne) SetLetterSpoiled(b bool) *DraftUpdateOne {
	duo.mutation.SetLetterSpoiled(b)
	return duo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (duo *DraftUpdateOne) SetOwnerID(id int) *DraftUpdateOne {
	duo.mutation.SetOwnerID(id)
	return duo
}

// SetOwner sets the "owner" edge to the User entity.
func (duo *DraftUpdateOne) SetOwner(u *User) *DraftUpdateOne {
	return duo.SetOwnerID(u.ID)
}

// AddHashtagIDs adds the "hashtags" edge to the Hashtag entity by IDs.
func (duo *DraftUpdateOne) AddHashtagIDs(ids ...int) *DraftUpdateOne {
	duo.mutation.AddHashtagIDs(ids...)
	return duo
}

// AddHashtags adds the "hashtags" edges to the Hashtag entity.
func (duo *DraftUpdateOne) AddHashtags(h ...*Hashtag) *DraftUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.AddHashtagIDs(ids...)
}

// SetWorkID sets the "work" edge to the Work entity by ID.
func (duo *DraftUpdateOne) SetWorkID(id int) *DraftUpdateOne {
	duo.mutation.SetWorkID(id)
	return duo
}

// SetNillableWorkID sets the "work" edge to the Work entity by ID if the given value is not nil.
func (duo *DraftUpdateOne) SetNillableWorkID(id *int) *DraftUpdateOne {
	if id != nil {
		duo = duo.SetWorkID(*id)
	}
	return duo
}

// SetWork sets the "work" edge to the Work entity.
func (duo *DraftUpdateOne) SetWork(w *Work) *DraftUpdateOne {
	return duo.SetWorkID(w.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (duo *DraftUpdateOne) SetCategoryID(id int) *DraftUpdateOne {
	duo.mutation.SetCategoryID(id)
	return duo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (duo *DraftUpdateOne) SetNillableCategoryID(id *int) *DraftUpdateOne {
	if id != nil {
		duo = duo.SetCategoryID(*id)
	}
	return duo
}

// SetCategory sets the "category" edge to the Category entity.
func (duo *DraftUpdateOne) SetCategory(c *Category) *DraftUpdateOne {
	return duo.SetCategoryID(c.ID)
}

// Mutation returns the DraftMutation object of the builder.
func (duo *DraftUpdateOne) Mutation() *DraftMutation {
	return duo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (duo *DraftUpdateOne) ClearOwner() *DraftUpdateOne {
	duo.mutation.ClearOwner()
	return duo
}

// ClearHashtags clears all "hashtags" edges to the Hashtag entity.
func (duo *DraftUpdateOne) ClearHashtags() *DraftUpdateOne {
	duo.mutation.ClearHashtags()
	return duo
}

// RemoveHashtagIDs removes the "hashtags" edge to Hashtag entities by IDs.
func (duo *DraftUpdateOne) RemoveHashtagIDs(ids ...int) *DraftUpdateOne {
	duo.mutation.RemoveHashtagIDs(ids...)
	return duo
}

// RemoveHashtags removes "hashtags" edges to Hashtag entities.
func (duo *DraftUpdateOne) RemoveHashtags(h ...*Hashtag) *DraftUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return duo.RemoveHashtagIDs(ids...)
}

// ClearWork clears the "work" edge to the Work entity.
func (duo *DraftUpdateOne) ClearWork() *DraftUpdateOne {
	duo.mutation.ClearWork()
	return duo
}

// ClearCategory clears the "category" edge to the Category entity.
func (duo *DraftUpdateOne) ClearCategory() *DraftUpdateOne {
	duo.mutation.ClearCategory()
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DraftUpdateOne) Select(field string, fields ...string) *DraftUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Draft entity.
func (duo *DraftUpdateOne) Save(ctx context.Context) (*Draft, error) {
	var (
		err  error
		node *Draft
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DraftMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Draft)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DraftMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DraftUpdateOne) SaveX(ctx context.Context) *Draft {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DraftUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DraftUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DraftUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdateTime(); !ok {
		v := draft.UpdateDefaultUpdateTime()
		duo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DraftUpdateOne) check() error {
	if _, ok := duo.mutation.OwnerID(); duo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Draft.owner"`)
	}
	return nil
}

func (duo *DraftUpdateOne) sqlSave(ctx context.Context) (_node *Draft, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   draft.Table,
			Columns: draft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: draft.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Draft.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, draft.FieldID)
		for _, f := range fields {
			if !draft.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != draft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: draft.FieldUpdateTime,
		})
	}
	if value, ok := duo.mutation.PraiseTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldPraiseTitle,
		})
	}
	if value, ok := duo.mutation.LetterTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldLetterTitle,
		})
	}
	if value, ok := duo.mutation.PraiseContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldPraiseContent,
		})
	}
	if value, ok := duo.mutation.LetterContent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: draft.FieldLetterContent,
		})
	}
	if value, ok := duo.mutation.PraiseSpoiled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: draft.FieldPraiseSpoiled,
		})
	}
	if value, ok := duo.mutation.LetterSpoiled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: draft.FieldLetterSpoiled,
		})
	}
	if duo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.OwnerTable,
			Columns: []string{draft.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.OwnerTable,
			Columns: []string{draft.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.HashtagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedHashtagsIDs(); len(nodes) > 0 && !duo.mutation.HashtagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.HashtagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   draft.HashtagsTable,
			Columns: draft.HashtagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hashtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.WorkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.WorkTable,
			Columns: []string{draft.WorkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.WorkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   draft.WorkTable,
			Columns: []string{draft.WorkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: work.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   draft.CategoryTable,
			Columns: []string{draft.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   draft.CategoryTable,
			Columns: []string{draft.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Draft{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{draft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
