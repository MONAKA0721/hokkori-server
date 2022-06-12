// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MONAKA0721/hokkori/ent/post"
)

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreateTime *time.Time
	UpdateTime *time.Time
	Title      string
	Content    string
	Type       post.Type
	OwnerID    int
}

// Mutate applies the CreatePostInput on the PostCreate builder.
func (i *CreatePostInput) Mutate(m *PostCreate) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetTitle(i.Title)
	m.SetContent(i.Content)
	m.SetType(i.Type)
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the CreatePostInput on the create builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c)
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdateTime *time.Time
	Title      *string
	Content    *string
	Type       *post.Type
	OwnerID    *int
	ClearOwner bool
}

// Mutate applies the UpdatePostInput on the PostMutation builder.
func (i *UpdatePostInput) Mutate(m *PostMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the update builder.
func (u *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdatePostInput on the update-one builder.
func (u *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name    string
	PostIDs []int
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetName(i.Name)
	if ids := i.PostIDs; len(ids) > 0 {
		m.AddPostIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name          *string
	AddPostIDs    []int
	RemovePostIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if ids := i.AddPostIDs; len(ids) > 0 {
		m.AddPostIDs(ids...)
	}
	if ids := i.RemovePostIDs; len(ids) > 0 {
		m.RemovePostIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
