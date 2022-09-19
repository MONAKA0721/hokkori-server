// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/category"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/ent/user"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PostCreate) SetCreateTime(t time.Time) *PostCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PostCreate) SetUpdateTime(t time.Time) *PostCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetType sets the "type" field.
func (pc *PostCreate) SetType(po post.Type) *PostCreate {
	pc.mutation.SetType(po)
	return pc
}

// SetSpoiled sets the "spoiled" field.
func (pc *PostCreate) SetSpoiled(b bool) *PostCreate {
	pc.mutation.SetSpoiled(b)
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PostCreate) SetOwnerID(id int) *PostCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PostCreate) SetOwner(u *User) *PostCreate {
	return pc.SetOwnerID(u.ID)
}

// AddHashtagIDs adds the "hashtags" edge to the Hashtag entity by IDs.
func (pc *PostCreate) AddHashtagIDs(ids ...int) *PostCreate {
	pc.mutation.AddHashtagIDs(ids...)
	return pc
}

// AddHashtags adds the "hashtags" edges to the Hashtag entity.
func (pc *PostCreate) AddHashtags(h ...*Hashtag) *PostCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pc.AddHashtagIDs(ids...)
}

// SetWorkID sets the "work" edge to the Work entity by ID.
func (pc *PostCreate) SetWorkID(id int) *PostCreate {
	pc.mutation.SetWorkID(id)
	return pc
}

// SetWork sets the "work" edge to the Work entity.
func (pc *PostCreate) SetWork(w *Work) *PostCreate {
	return pc.SetWorkID(w.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (pc *PostCreate) SetCategoryID(id int) *PostCreate {
	pc.mutation.SetCategoryID(id)
	return pc
}

// SetCategory sets the "category" edge to the Category entity.
func (pc *PostCreate) SetCategory(c *Category) *PostCreate {
	return pc.SetCategoryID(c.ID)
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (pc *PostCreate) AddLikedUserIDs(ids ...int) *PostCreate {
	pc.mutation.AddLikedUserIDs(ids...)
	return pc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (pc *PostCreate) AddLikedUsers(u ...*User) *PostCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddLikedUserIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := post.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := post.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Post.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Post.update_time"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if v, ok := pc.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Post.type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := post.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Post.type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Spoiled(); !ok {
		return &ValidationError{Name: "spoiled", err: errors.New(`ent: missing required field "Post.spoiled"`)}
	}
	if _, ok := pc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Post.owner"`)}
	}
	if _, ok := pc.mutation.WorkID(); !ok {
		return &ValidationError{Name: "work", err: errors.New(`ent: missing required edge "Post.work"`)}
	}
	if _, ok := pc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "Post.category"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: post.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: post.FieldType,
		})
		_node.Type = value
	}
	if value, ok := pc.mutation.Spoiled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: post.FieldSpoiled,
		})
		_node.Spoiled = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.OwnerTable,
			Columns: []string{post.OwnerColumn},
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
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.HashtagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.HashtagsTable,
			Columns: post.HashtagsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.WorkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.WorkTable,
			Columns: []string{post.WorkColumn},
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
		_node.work_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.CategoryTable,
			Columns: []string{post.CategoryColumn},
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
		_node.post_category = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LikedUsersTable,
			Columns: post.LikedUsersPrimaryKey,
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
		createE := &LikeCreate{config: pc.config, mutation: newLikeMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
