// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/predicate"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// WorkDelete is the builder for deleting a Work entity.
type WorkDelete struct {
	config
	hooks    []Hook
	mutation *WorkMutation
}

// Where appends a list predicates to the WorkDelete builder.
func (wd *WorkDelete) Where(ps ...predicate.Work) *WorkDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wd.hooks) == 0 {
		affected, err = wd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wd.mutation = mutation
			affected, err = wd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wd.hooks) - 1; i >= 0; i-- {
			if wd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: work.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: work.FieldID,
			},
		},
	}
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// WorkDeleteOne is the builder for deleting a single Work entity.
type WorkDeleteOne struct {
	wd *WorkDelete
}

// Exec executes the deletion query.
func (wdo *WorkDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{work.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkDeleteOne) ExecX(ctx context.Context) {
	wdo.wd.ExecX(ctx)
}
