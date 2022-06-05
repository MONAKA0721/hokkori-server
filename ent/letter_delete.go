// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/letter"
	"github.com/MONAKA0721/hokkori/ent/predicate"
)

// LetterDelete is the builder for deleting a Letter entity.
type LetterDelete struct {
	config
	hooks    []Hook
	mutation *LetterMutation
}

// Where appends a list predicates to the LetterDelete builder.
func (ld *LetterDelete) Where(ps ...predicate.Letter) *LetterDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LetterDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ld.hooks) == 0 {
		affected, err = ld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LetterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ld.mutation = mutation
			affected, err = ld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ld.hooks) - 1; i >= 0; i-- {
			if ld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LetterDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LetterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: letter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: letter.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
}

// LetterDeleteOne is the builder for deleting a single Letter entity.
type LetterDeleteOne struct {
	ld *LetterDelete
}

// Exec executes the deletion query.
func (ldo *LetterDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{letter.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LetterDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}