// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MONAKA0721/hokkori/ent/post"
)

// CreateHashtagInput represents a mutation input for creating hashtags.
type CreateHashtagInput struct {
	Title   string
	PostIDs []int
}

// Mutate applies the CreateHashtagInput on the HashtagMutation builder.
func (i *CreateHashtagInput) Mutate(m *HashtagMutation) {
	m.SetTitle(i.Title)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
}

// SetInput applies the change-set in the CreateHashtagInput on the HashtagCreate builder.
func (c *HashtagCreate) SetInput(i CreateHashtagInput) *HashtagCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateHashtagInput represents a mutation input for updating hashtags.
type UpdateHashtagInput struct {
	Title         *string
	AddPostIDs    []int
	RemovePostIDs []int
}

// Mutate applies the UpdateHashtagInput on the HashtagMutation builder.
func (i *UpdateHashtagInput) Mutate(m *HashtagMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateHashtagInput on the HashtagUpdate builder.
func (c *HashtagUpdate) SetInput(i UpdateHashtagInput) *HashtagUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateHashtagInput on the HashtagUpdateOne builder.
func (c *HashtagUpdateOne) SetInput(i UpdateHashtagInput) *HashtagUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreateTime *time.Time
	UpdateTime *time.Time
	Title      string
	Content    string
	Type       post.Type
	Spoiled    bool
	OwnerID    int
	HashtagIDs []int
	WorkID     int
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetTitle(i.Title)
	m.SetContent(i.Content)
	m.SetType(i.Type)
	m.SetSpoiled(i.Spoiled)
	m.SetOwnerID(i.OwnerID)
	if v := i.HashtagIDs; len(v) > 0 {
		m.AddHashtagIDs(v...)
	}
	m.SetWorkID(i.WorkID)
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdateTime       *time.Time
	Title            *string
	Content          *string
	Type             *post.Type
	Spoiled          *bool
	ClearOwner       bool
	OwnerID          *int
	AddHashtagIDs    []int
	RemoveHashtagIDs []int
	ClearWork        bool
	WorkID           *int
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
	if v := i.Spoiled; v != nil {
		m.SetSpoiled(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
	if v := i.AddHashtagIDs; len(v) > 0 {
		m.AddHashtagIDs(v...)
	}
	if v := i.RemoveHashtagIDs; len(v) > 0 {
		m.RemoveHashtagIDs(v...)
	}
	if i.ClearWork {
		m.ClearWork()
	}
	if v := i.WorkID; v != nil {
		m.SetWorkID(*v)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdate builder.
func (c *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdateOne builder.
func (c *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name    string
	PostIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
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
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateWorkInput represents a mutation input for creating works.
type CreateWorkInput struct {
	Title   string
	PostIDs []int
}

// Mutate applies the CreateWorkInput on the WorkMutation builder.
func (i *CreateWorkInput) Mutate(m *WorkMutation) {
	m.SetTitle(i.Title)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
}

// SetInput applies the change-set in the CreateWorkInput on the WorkCreate builder.
func (c *WorkCreate) SetInput(i CreateWorkInput) *WorkCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateWorkInput represents a mutation input for updating works.
type UpdateWorkInput struct {
	Title         *string
	AddPostIDs    []int
	RemovePostIDs []int
}

// Mutate applies the UpdateWorkInput on the WorkMutation builder.
func (i *UpdateWorkInput) Mutate(m *WorkMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateWorkInput on the WorkUpdate builder.
func (c *WorkUpdate) SetInput(i UpdateWorkInput) *WorkUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateWorkInput on the WorkUpdateOne builder.
func (c *WorkUpdateOne) SetInput(i UpdateWorkInput) *WorkUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
