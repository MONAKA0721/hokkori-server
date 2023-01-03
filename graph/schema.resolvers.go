package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/ent/bookmark"
	"github.com/MONAKA0721/hokkori/ent/category"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
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
		randomUUID, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		output, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(r.config.S3.BucketName),
			Key:         aws.String(randomUUID.String()),
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
func (r *mutationResolver) CreatePosts(ctx context.Context, input ent.CreatePostInput, input2 ent.CreatePostInput, hashtagTitles []*string, image *graphql.Upload) (*ent.Post, error) {
	newHashtagIDs, err := GetNewHashtagIDs(ctx, r.client, hashtagTitles)
	if err != nil {
		return nil, err
	}

	input.HashtagIDs = append(input.HashtagIDs, newHashtagIDs...)
	_, err = r.client.Post.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	input2.HashtagIDs = append(input2.HashtagIDs, newHashtagIDs...)

	if image != nil {
		randomUUID, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		output, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(r.config.S3.BucketName),
			Key:         aws.String(randomUUID.String()),
			Body:        image.File,
			ContentType: aws.String(image.ContentType),
			ACL:         types.ObjectCannedACLPublicRead,
		})
		if err != nil {
			return nil, err
		}
		input2.Thumbnail = &output.Location
	}

	return r.client.Post.Create().SetInput(input2).Save(ctx)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput, hashtagTitles []*string) (*ent.Post, error) {
	newHashtagIDs, err := GetNewHashtagIDs(ctx, r.client, hashtagTitles)
	if err != nil {
		return nil, err
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

// CreateDraft is the resolver for the createDraft field.
func (r *mutationResolver) CreateDraft(ctx context.Context, input ent.CreateDraftInput, hashtagTitles []*string) (*ent.Draft, error) {
	newHashtagIDs, err := GetNewHashtagIDs(ctx, r.client, hashtagTitles)
	if err != nil {
		return nil, err
	}

	input.HashtagIDs = append(input.HashtagIDs, newHashtagIDs...)

	return r.client.Draft.Create().SetInput(input).Save(ctx)
}

// UpdateDraft is the resolver for the updateDraft field.
func (r *mutationResolver) UpdateDraft(ctx context.Context, id int, input ent.UpdateDraftInput, hashtagTitles []*string) (*ent.Draft, error) {
	newHashtagIDs, err := GetNewHashtagIDs(ctx, r.client, hashtagTitles)
	if err != nil {
		return nil, err
	}

	input.AddHashtagIDs = append(input.AddHashtagIDs, newHashtagIDs...)

	return r.client.Draft.UpdateOneID(id).ClearHashtags().SetInput(input).Save(ctx)
}

// DeleteDraft is the resolver for the deleteDraft field.
func (r *mutationResolver) DeleteDraft(ctx context.Context, input model.DeleteDraftInput) (*model.DeleteDraftPayload, error) {
	d, err := r.client.Draft.Get(ctx, input.DraftID)
	if err != nil {
		return nil, err
	}

	owner, err := d.QueryOwner().Only(ctx)
	if err != nil {
		return nil, err
	}
	if owner.ID != input.UserID {
		return nil, fmt.Errorf("onwer does not match draftID: %v, userID: %v", input.DraftID, input.UserID)
	}

	if err := r.client.Draft.DeleteOne(d).Exec(ctx); err != nil {
		return nil, err
	}

	return &model.DeleteDraftPayload{
		ClientMutationID: input.ClientMutationID,
		DraftID:          &input.DraftID,
		UserID:           &input.UserID,
	}, nil
}

// LikedPosts is the resolver for the likedPosts field.
func (r *queryResolver) LikedPosts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().Unique(false).Order(
		func(s *sql.Selector) {
			t := sql.Table(like.Table)
			s.LeftJoin(t).On(s.C(post.FieldID), t.C(like.PostColumn))
			s.GroupBy(s.C(post.FieldID))
			s.OrderBy(sql.Desc(sql.Count(t.C(like.PostColumn))))
		}).Paginate(ctx, after, first, before, last, ent.WithPostFilter(where.Filter))
}

// TopicWorks is the resolver for the topicWorks field.
func (r *queryResolver) TopicWorks(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.WorkWhereInput) (*ent.WorkConnection, error) {
	return r.client.Work.Query().Unique(false).Order(
		func(s *sql.Selector) {
			t := sql.Table(post.Table)
			s.LeftJoin(t).On(s.C(work.FieldID), t.C(post.WorkColumn))
			s.GroupBy(s.C(work.FieldID))
			s.OrderBy(sql.Desc(sql.Count(t.C(post.FieldID))))
		}).Paginate(ctx, after, first, before, last, ent.WithWorkFilter(where.Filter))
}

// TopicHashtags is the resolver for the topicHashtags field.
func (r *queryResolver) TopicHashtags(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.HashtagWhereInput) (*ent.HashtagConnection, error) {
	return r.client.Hashtag.Query().Unique(false).Order(
		func(s *sql.Selector) {
			t := sql.Table(post.HashtagsTable)
			s.LeftJoin(t).On(s.C(hashtag.FieldID), t.C("hashtag_id"))
			s.GroupBy(s.C(hashtag.FieldID))
			s.OrderBy(sql.Desc(sql.Count(t.C("post_id"))))
		}).Paginate(ctx, after, first, before, last, ent.WithHashtagFilter(where.Filter))
}

// WorkCategories is the resolver for the workCategories field.
func (r *queryResolver) WorkCategories(ctx context.Context, workID int) ([]*model.WorkCategory, error) {
	var workCategories []*model.WorkCategory
	err := r.client.Debug().Category.Query().
		GroupBy(category.FieldID, category.FieldName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(post.Table)
			s.Join(t).On(s.C(category.FieldID), t.C(post.CategoryColumn))
			s.Where(sql.EQ(t.C(post.WorkColumn), workID))
			s.OrderBy(sql.Desc(sql.Count(t.C(post.FieldID))))
			return sql.As(sql.Count(t.C(post.FieldID)), "count")
		}).
		Scan(ctx, &workCategories)
	if err != nil {
		return nil, err
	}

	return workCategories, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
