package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/graph/generated"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreatePosts is the resolver for the createPosts field.
func (r *mutationResolver) CreatePosts(ctx context.Context, input ent.CreatePostInput, input2 ent.CreatePostInput) (*ent.Post, error) {
	_, err := r.client.Post.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.Post.Create().SetInput(input2).Save(ctx)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	return r.client.Post.Create().SetInput(input).Save(ctx)
}

// CreateHashtag is the resolver for the createHashtag field.
func (r *mutationResolver) CreateHashtag(ctx context.Context, input ent.CreateHashtagInput) (*ent.Hashtag, error) {
	return r.client.Hashtag.Create().SetInput(input).Save(ctx)
}

// CreateWork is the resolver for the createWork field.
func (r *mutationResolver) CreateWork(ctx context.Context, input ent.CreateWorkInput) (*ent.Work, error) {
	return r.client.Work.Create().SetInput(input).Save(ctx)
}

// LikedPosts is the resolver for the likedPosts field.
func (r *queryResolver) LikedPosts(ctx context.Context, first int, typeArg post.Type) ([]*ent.Post, error) {
	query := fmt.Sprintf(`
SELECT
	p.id
FROM
	posts AS p
LEFT JOIN
	likes AS l on p.id = l.post_id
WHERE
	p.type = '%s'
GROUP BY
	p.id
ORDER BY
	COUNT(l.post_id) DESC
LIMIT %d`, typeArg, first)
	rows, err := r.client.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	postIDs := make([]int, 0)
	for rows.Next() {
		var postID int
		if err := rows.Scan(&postID); err != nil {
			break
		}
		postIDs = append(postIDs, postID)
	}

	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	// Check for errors during rows "Close".
	// This may be more important if multiple statements are executed
	// in a single batch and rows were written as well as read.
	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	// Check for row scan error.
	if err != nil {
		return nil, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	queryPosts, err := r.client.Post.Query().Where(post.IDIn(postIDs...)).All(ctx)
	if err != nil {
		return nil, err
	}

	orderedPosts := make([]*ent.Post, 0)
	for _, postID := range postIDs {
		for _, queryPost := range queryPosts {
			if queryPost.ID == postID {
				orderedPosts = append(orderedPosts, queryPost)
				break
			}
		}
	}

	return orderedPosts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
