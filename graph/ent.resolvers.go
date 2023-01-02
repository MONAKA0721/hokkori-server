package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.CategoryWhereInput) (*ent.CategoryConnection, error) {
	return r.client.Category.Query().Paginate(ctx, after, first, before, last,
		ent.WithCategoryFilter(where.Filter))
}

// Drafts is the resolver for the drafts field.
func (r *queryResolver) Drafts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DraftOrder, where *ent.DraftWhereInput) (*ent.DraftConnection, error) {
	return r.client.Draft.Query().Paginate(ctx, after, first, before, last,
		ent.WithDraftOrder(orderBy),
		ent.WithDraftFilter(where.Filter))
}

// Hashtags is the resolver for the hashtags field.
func (r *queryResolver) Hashtags(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.HashtagWhereInput) (*ent.HashtagConnection, error) {
	return r.client.Hashtag.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithHashtagFilter(where.Filter),
		)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
			ent.WithPostFilter(where.Filter),
		)
}

// Works is the resolver for the works field.
func (r *queryResolver) Works(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.WorkWhereInput) (*ent.WorkConnection, error) {
	return r.client.Work.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWorkFilter(where.Filter),
		)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
