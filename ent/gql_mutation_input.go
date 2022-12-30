// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MONAKA0721/hokkori/ent/post"
)

// CreateDraftInput represents a mutation input for creating drafts.
type CreateDraftInput struct {
	CreateTime    *time.Time
	UpdateTime    *time.Time
	PraiseTitle   string
	LetterTitle   string
	PraiseContent string
	LetterContent string
	PraiseSpoiled bool
	LetterSpoiled bool
	OwnerID       int
	HashtagIDs    []int
	WorkID        *int
	CategoryID    *int
}

// Mutate applies the CreateDraftInput on the DraftMutation builder.
func (i *CreateDraftInput) Mutate(m *DraftMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetPraiseTitle(i.PraiseTitle)
	m.SetLetterTitle(i.LetterTitle)
	m.SetPraiseContent(i.PraiseContent)
	m.SetLetterContent(i.LetterContent)
	m.SetPraiseSpoiled(i.PraiseSpoiled)
	m.SetLetterSpoiled(i.LetterSpoiled)
	m.SetOwnerID(i.OwnerID)
	if v := i.HashtagIDs; len(v) > 0 {
		m.AddHashtagIDs(v...)
	}
	if v := i.WorkID; v != nil {
		m.SetWorkID(*v)
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
}

// SetInput applies the change-set in the CreateDraftInput on the DraftCreate builder.
func (c *DraftCreate) SetInput(i CreateDraftInput) *DraftCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateDraftInput represents a mutation input for updating drafts.
type UpdateDraftInput struct {
	UpdateTime       *time.Time
	PraiseTitle      *string
	LetterTitle      *string
	PraiseContent    *string
	LetterContent    *string
	PraiseSpoiled    *bool
	LetterSpoiled    *bool
	ClearOwner       bool
	OwnerID          *int
	AddHashtagIDs    []int
	RemoveHashtagIDs []int
	ClearWork        bool
	WorkID           *int
	ClearCategory    bool
	CategoryID       *int
}

// Mutate applies the UpdateDraftInput on the DraftMutation builder.
func (i *UpdateDraftInput) Mutate(m *DraftMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.PraiseTitle; v != nil {
		m.SetPraiseTitle(*v)
	}
	if v := i.LetterTitle; v != nil {
		m.SetLetterTitle(*v)
	}
	if v := i.PraiseContent; v != nil {
		m.SetPraiseContent(*v)
	}
	if v := i.LetterContent; v != nil {
		m.SetLetterContent(*v)
	}
	if v := i.PraiseSpoiled; v != nil {
		m.SetPraiseSpoiled(*v)
	}
	if v := i.LetterSpoiled; v != nil {
		m.SetLetterSpoiled(*v)
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
	if i.ClearCategory {
		m.ClearCategory()
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
}

// SetInput applies the change-set in the UpdateDraftInput on the DraftUpdate builder.
func (c *DraftUpdate) SetInput(i UpdateDraftInput) *DraftUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateDraftInput on the DraftUpdateOne builder.
func (c *DraftUpdateOne) SetInput(i UpdateDraftInput) *DraftUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateHashtagInput represents a mutation input for creating hashtags.
type CreateHashtagInput struct {
	Title    string
	PostIDs  []int
	DraftIDs []int
}

// Mutate applies the CreateHashtagInput on the HashtagMutation builder.
func (i *CreateHashtagInput) Mutate(m *HashtagMutation) {
	m.SetTitle(i.Title)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.DraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
}

// SetInput applies the change-set in the CreateHashtagInput on the HashtagCreate builder.
func (c *HashtagCreate) SetInput(i CreateHashtagInput) *HashtagCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateHashtagInput represents a mutation input for updating hashtags.
type UpdateHashtagInput struct {
	Title          *string
	AddPostIDs     []int
	RemovePostIDs  []int
	AddDraftIDs    []int
	RemoveDraftIDs []int
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
	if v := i.AddDraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
	if v := i.RemoveDraftIDs; len(v) > 0 {
		m.RemoveDraftIDs(v...)
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
	CreateTime        *time.Time
	UpdateTime        *time.Time
	Title             string
	Content           string
	Type              post.Type
	Spoiled           bool
	Thumbnail         *string
	OwnerID           int
	HashtagIDs        []int
	WorkID            int
	CategoryID        int
	LikedUserIDs      []int
	BookmarkedUserIDs []int
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
	if v := i.Thumbnail; v != nil {
		m.SetThumbnail(*v)
	}
	m.SetOwnerID(i.OwnerID)
	if v := i.HashtagIDs; len(v) > 0 {
		m.AddHashtagIDs(v...)
	}
	m.SetWorkID(i.WorkID)
	m.SetCategoryID(i.CategoryID)
	if v := i.LikedUserIDs; len(v) > 0 {
		m.AddLikedUserIDs(v...)
	}
	if v := i.BookmarkedUserIDs; len(v) > 0 {
		m.AddBookmarkedUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdateTime              *time.Time
	Title                   *string
	Content                 *string
	Type                    *post.Type
	Spoiled                 *bool
	ClearThumbnail          bool
	Thumbnail               *string
	ClearOwner              bool
	OwnerID                 *int
	AddHashtagIDs           []int
	RemoveHashtagIDs        []int
	ClearWork               bool
	WorkID                  *int
	ClearCategory           bool
	CategoryID              *int
	AddLikedUserIDs         []int
	RemoveLikedUserIDs      []int
	AddBookmarkedUserIDs    []int
	RemoveBookmarkedUserIDs []int
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
	if i.ClearThumbnail {
		m.ClearThumbnail()
	}
	if v := i.Thumbnail; v != nil {
		m.SetThumbnail(*v)
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
	if i.ClearCategory {
		m.ClearCategory()
	}
	if v := i.CategoryID; v != nil {
		m.SetCategoryID(*v)
	}
	if v := i.AddLikedUserIDs; len(v) > 0 {
		m.AddLikedUserIDs(v...)
	}
	if v := i.RemoveLikedUserIDs; len(v) > 0 {
		m.RemoveLikedUserIDs(v...)
	}
	if v := i.AddBookmarkedUserIDs; len(v) > 0 {
		m.AddBookmarkedUserIDs(v...)
	}
	if v := i.RemoveBookmarkedUserIDs; len(v) > 0 {
		m.RemoveBookmarkedUserIDs(v...)
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
	Name              string
	Username          *string
	Profile           *string
	AvatarURL         *string
	PostIDs           []int
	LikedPostIDs      []int
	BookmarkedPostIDs []int
	FollowerIDs       []int
	FollowingIDs      []int
	DraftIDs          []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetName(i.Name)
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.LikedPostIDs; len(v) > 0 {
		m.AddLikedPostIDs(v...)
	}
	if v := i.BookmarkedPostIDs; len(v) > 0 {
		m.AddBookmarkedPostIDs(v...)
	}
	if v := i.FollowerIDs; len(v) > 0 {
		m.AddFollowerIDs(v...)
	}
	if v := i.FollowingIDs; len(v) > 0 {
		m.AddFollowingIDs(v...)
	}
	if v := i.DraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name                    *string
	ClearUsername           bool
	Username                *string
	ClearProfile            bool
	Profile                 *string
	ClearAvatarURL          bool
	AvatarURL               *string
	AddPostIDs              []int
	RemovePostIDs           []int
	AddLikedPostIDs         []int
	RemoveLikedPostIDs      []int
	AddBookmarkedPostIDs    []int
	RemoveBookmarkedPostIDs []int
	AddFollowerIDs          []int
	RemoveFollowerIDs       []int
	AddFollowingIDs         []int
	RemoveFollowingIDs      []int
	AddDraftIDs             []int
	RemoveDraftIDs          []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearUsername {
		m.ClearUsername()
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if i.ClearProfile {
		m.ClearProfile()
	}
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if i.ClearAvatarURL {
		m.ClearAvatarURL()
	}
	if v := i.AvatarURL; v != nil {
		m.SetAvatarURL(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddLikedPostIDs; len(v) > 0 {
		m.AddLikedPostIDs(v...)
	}
	if v := i.RemoveLikedPostIDs; len(v) > 0 {
		m.RemoveLikedPostIDs(v...)
	}
	if v := i.AddBookmarkedPostIDs; len(v) > 0 {
		m.AddBookmarkedPostIDs(v...)
	}
	if v := i.RemoveBookmarkedPostIDs; len(v) > 0 {
		m.RemoveBookmarkedPostIDs(v...)
	}
	if v := i.AddFollowerIDs; len(v) > 0 {
		m.AddFollowerIDs(v...)
	}
	if v := i.RemoveFollowerIDs; len(v) > 0 {
		m.RemoveFollowerIDs(v...)
	}
	if v := i.AddFollowingIDs; len(v) > 0 {
		m.AddFollowingIDs(v...)
	}
	if v := i.RemoveFollowingIDs; len(v) > 0 {
		m.RemoveFollowingIDs(v...)
	}
	if v := i.AddDraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
	if v := i.RemoveDraftIDs; len(v) > 0 {
		m.RemoveDraftIDs(v...)
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
	Title     string
	Thumbnail *string
	PostIDs   []int
	DraftIDs  []int
}

// Mutate applies the CreateWorkInput on the WorkMutation builder.
func (i *CreateWorkInput) Mutate(m *WorkMutation) {
	m.SetTitle(i.Title)
	if v := i.Thumbnail; v != nil {
		m.SetThumbnail(*v)
	}
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.DraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
}

// SetInput applies the change-set in the CreateWorkInput on the WorkCreate builder.
func (c *WorkCreate) SetInput(i CreateWorkInput) *WorkCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateWorkInput represents a mutation input for updating works.
type UpdateWorkInput struct {
	Title          *string
	ClearThumbnail bool
	Thumbnail      *string
	AddPostIDs     []int
	RemovePostIDs  []int
	AddDraftIDs    []int
	RemoveDraftIDs []int
}

// Mutate applies the UpdateWorkInput on the WorkMutation builder.
func (i *UpdateWorkInput) Mutate(m *WorkMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if i.ClearThumbnail {
		m.ClearThumbnail()
	}
	if v := i.Thumbnail; v != nil {
		m.SetThumbnail(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddDraftIDs; len(v) > 0 {
		m.AddDraftIDs(v...)
	}
	if v := i.RemoveDraftIDs; len(v) > 0 {
		m.RemoveDraftIDs(v...)
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
