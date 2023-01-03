// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/MONAKA0721/hokkori/ent"
)

type BookmarkPostInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	PostID           int     `json:"postID"`
}

type BookmarkPostPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	Post             *ent.Post `json:"post"`
}

type DeleteDraftInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	DraftID          int     `json:"draftId"`
	UserID           int     `json:"userID"`
}

type DeleteDraftPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	DraftID          *int    `json:"draftId"`
	UserID           *int    `json:"userID"`
}

type FollowUserInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	FollowerID       int     `json:"followerID"`
}

type FollowUserPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	User             *ent.User `json:"user"`
}

type LikePostInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	PostID           int     `json:"postID"`
}

type LikePostPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	Post             *ent.Post `json:"post"`
}

type UnbookmarkPostInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	PostID           int     `json:"postID"`
}

type UnbookmarkPostPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	Post             *ent.Post `json:"post"`
}

type UnfollowUserInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	FollowerID       int     `json:"followerID"`
}

type UnfollowUserPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	User             *ent.User `json:"user"`
}

type UnlikePostInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	UserID           int     `json:"userID"`
	PostID           int     `json:"postID"`
}

type UnlikePostPayload struct {
	ClientMutationID *string   `json:"clientMutationId"`
	Post             *ent.Post `json:"post"`
}

type WorkCategory struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
