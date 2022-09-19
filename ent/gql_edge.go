// Code generated by ent, DO NOT EDIT.

package ent

import "context"

func (c *Category) Post(ctx context.Context) ([]*Post, error) {
	result, err := c.Edges.PostOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryPost().All(ctx)
	}
	return result, err
}

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

func (po *Post) Work(ctx context.Context) (*Work, error) {
	result, err := po.Edges.WorkOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryWork().Only(ctx)
	}
	return result, err
}

func (po *Post) Category(ctx context.Context) (*Category, error) {
	result, err := po.Edges.CategoryOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryCategory().Only(ctx)
	}
	return result, err
}

func (po *Post) LikedUsers(ctx context.Context) ([]*User, error) {
	result, err := po.Edges.LikedUsersOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryLikedUsers().All(ctx)
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

func (u *User) LikedPosts(ctx context.Context) ([]*Post, error) {
	result, err := u.Edges.LikedPostsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryLikedPosts().All(ctx)
	}
	return result, err
}

func (w *Work) Posts(ctx context.Context) ([]*Post, error) {
	result, err := w.Edges.PostsOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryPosts().All(ctx)
	}
	return result, err
}
