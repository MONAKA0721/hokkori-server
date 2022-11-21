package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/ent/bookmark"
	"github.com/MONAKA0721/hokkori/ent/like"
	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/ent/user"
	"github.com/MONAKA0721/hokkori/ent/work"
	"github.com/MONAKA0721/hokkori/graph/generated"
	"github.com/MONAKA0721/hokkori/graph/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput, image *graphql.Upload) (*ent.User, error) {
	if image != nil {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		output, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(r.config.S3.BucketName),
			Key:         aws.String(uuid.String()),
			Body:        image.File,
			ContentType: aws.String(image.ContentType),
			ACL:         types.ObjectCannedACLPublicRead,
		})
		if err != nil {
			return nil, err
		}
		input.AvatarURL = &output.Location
	}

	return r.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreatePosts is the resolver for the createPosts field.
func (r *mutationResolver) CreatePosts(ctx context.Context, input ent.CreatePostInput, input2 ent.CreatePostInput, hashtagTitles []*string) (*ent.Post, error) {
	createHashtags := make([]*ent.HashtagCreate, 0)
	for _, hashtagTitle := range hashtagTitles {
		createHashtags = append(createHashtags, r.client.Hashtag.Create().SetTitle(*hashtagTitle))
	}

	hashtags, err := r.client.Hashtag.CreateBulk(createHashtags...).Save(ctx)
	if err != nil {
		return nil, err
	}

	newHashtagIDs := make([]int, len(hashtags))
	for i, hashtag := range hashtags {
		newHashtagIDs[i] = hashtag.ID
	}

	input.HashtagIDs = append(input.HashtagIDs, newHashtagIDs...)
	_, err = r.client.Post.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	input2.HashtagIDs = append(input2.HashtagIDs, newHashtagIDs...)
	return r.client.Post.Create().SetInput(input2).Save(ctx)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput, hashtagTitles []*string) (*ent.Post, error) {
	createHashtags := make([]*ent.HashtagCreate, 0)
	for _, hashtagTitle := range hashtagTitles {
		createHashtags = append(createHashtags, r.client.Hashtag.Create().SetTitle(*hashtagTitle))
	}

	hashtags, err := r.client.Hashtag.CreateBulk(createHashtags...).Save(ctx)
	if err != nil {
		return nil, err
	}

	newHashtagIDs := make([]int, len(hashtags))
	for i, hashtag := range hashtags {
		newHashtagIDs[i] = hashtag.ID
	}

	input.HashtagIDs = append(input.HashtagIDs, newHashtagIDs...)

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

// LikePost is the resolver for the LikePost field.
func (r *mutationResolver) LikePost(ctx context.Context, input model.LikePostInput) (*model.LikePostPayload, error) {
	_, err := r.client.Like.Create().SetUserID(input.UserID).SetPostID(input.PostID).Save(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Post.Get(ctx, input.PostID)
	if err != nil {
		return nil, err
	}

	return &model.LikePostPayload{
		ClientMutationID: input.ClientMutationID,
		Post:             p,
	}, nil
}

// UnlikePost is the resolver for the UnlikePost field.
func (r *mutationResolver) UnlikePost(ctx context.Context, input model.UnlikePostInput) (*model.UnlikePostPayload, error) {
	_, err := r.client.Like.Delete().Where(like.HasUserWith(user.IDEQ(input.UserID)), like.HasPostWith(post.IDEQ(input.PostID))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Post.Get(ctx, input.PostID)
	if err != nil {
		return nil, err
	}

	return &model.UnlikePostPayload{
		ClientMutationID: input.ClientMutationID,
		Post:             p,
	}, nil
}

// BookmarkPost is the resolver for the bookmarkPost field.
func (r *mutationResolver) BookmarkPost(ctx context.Context, input model.BookmarkPostInput) (*model.BookmarkPostPayload, error) {
	_, err := r.client.Bookmark.Create().SetUserID(input.UserID).SetPostID(input.PostID).Save(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Post.Get(ctx, input.PostID)
	if err != nil {
		return nil, err
	}

	return &model.BookmarkPostPayload{
		ClientMutationID: input.ClientMutationID,
		Post:             p,
	}, nil
}

// UnbookmarkPost is the resolver for the unbookmarkPost field.
func (r *mutationResolver) UnbookmarkPost(ctx context.Context, input model.UnbookmarkPostInput) (*model.UnbookmarkPostPayload, error) {
	_, err := r.client.Bookmark.Delete().Where(bookmark.HasUserWith(user.IDEQ(input.UserID)), bookmark.HasPostWith(post.IDEQ(input.PostID))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Post.Get(ctx, input.PostID)
	if err != nil {
		return nil, err
	}

	return &model.UnbookmarkPostPayload{
		ClientMutationID: input.ClientMutationID,
		Post:             p,
	}, nil
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, input model.FollowUserInput) (*model.FollowUserPayload, error) {
	_, err := r.client.User.UpdateOneID(input.UserID).AddFollowerIDs(input.FollowerID).Save(ctx)
	if err != nil {
		return nil, err
	}

	u, err := r.client.User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return &model.FollowUserPayload{
		ClientMutationID: input.ClientMutationID,
		User:             u,
	}, nil
}

// UnfollowUser is the resolver for the unfollowUser field.
func (r *mutationResolver) UnfollowUser(ctx context.Context, input model.UnfollowUserInput) (*model.UnfollowUserPayload, error) {
	_, err := r.client.User.UpdateOneID(input.UserID).RemoveFollowerIDs(input.FollowerID).Save(ctx)
	if err != nil {
		return nil, err
	}

	u, err := r.client.User.Get(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return &model.UnfollowUserPayload{
		ClientMutationID: input.ClientMutationID,
		User:             u,
	}, nil
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

// TopicWorks is the resolver for the topicWorks field.
func (r *queryResolver) TopicWorks(ctx context.Context, first int) ([]*ent.Work, error) {
	query := fmt.Sprintf(`
SELECT
	w.id
FROM
	works AS w
LEFT JOIN
	posts AS p on w.id = p.work_posts
GROUP BY
	w.id
ORDER BY
	COUNT(p.id) DESC
LIMIT %d`, first)
	rows, err := r.client.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	workIDs := make([]int, 0)
	for rows.Next() {
		var workID int
		if err := rows.Scan(&workID); err != nil {
			break
		}
		workIDs = append(workIDs, workID)
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

	queryWorks, err := r.client.Work.Query().Where(work.IDIn(workIDs...)).All(ctx)
	if err != nil {
		return nil, err
	}

	orderedWorks := make([]*ent.Work, 0)
	for _, workID := range workIDs {
		for _, queryWork := range queryWorks {
			if queryWork.ID == workID {
				orderedWorks = append(orderedWorks, queryWork)
				break
			}
		}
	}

	return orderedWorks, nil
}

// WorkCategories is the resolver for the workCategories field.
func (r *queryResolver) WorkCategories(ctx context.Context, workID int) ([]*model.WorkCategory, error) {
	query := fmt.Sprintf(`
SELECT
	c.id AS categoryID,
	c.name AS categoryName,
	COUNT(p.id) AS postCount
FROM
	posts AS p
INNER JOIN
	categories AS c on p.post_category = c.id
WHERE
	p.work_posts = %d
GROUP BY
	c.id
ORDER BY
	COUNT(p.id) DESC
`, workID)
	rows, err := r.client.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	workCategories := make([]*model.WorkCategory, 0)
	for rows.Next() {
		var categoryID int
		var categoryName string
		var postCount int
		if err := rows.Scan(&categoryID, &categoryName, &postCount); err != nil {
			fmt.Println(err.Error())
			break
		}
		workCategories = append(workCategories, &model.WorkCategory{
			CategoryID:   categoryID,
			CategoryName: categoryName,
			PostCount:    postCount,
		})
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

	return workCategories, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
