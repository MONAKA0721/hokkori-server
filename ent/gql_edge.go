// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (h *Hashtag) Posts(ctx context.Context) ([]*Post, error) {
	result, err := h.Edges.PostsOrErr()
	if IsNotLoaded(err) {
		result, err = h.QueryPosts().All(ctx)
	}
	return result, err
}

func (po *Post) Owner(ctx context.Context) (*User, error) {
	result, err := po.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryOwner().Only(ctx)
	}
	return result, err
}

func (po *Post) Hashtags(ctx context.Context) ([]*Hashtag, error) {
	result, err := po.Edges.HashtagsOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryHashtags().All(ctx)
	}
	return result, err
}

func (u *User) Posts(ctx context.Context) ([]*Post, error) {
	result, err := u.Edges.PostsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}
