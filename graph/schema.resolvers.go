package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/graph/generated"
	"github.com/MONAKA0721/hokkori/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*ent.Todo, error) {
	return r.client.Todo.Create().
		SetText(input.Text).
		Save(ctx)
}

func (r *mutationResolver) CreateLetter(ctx context.Context, input model.CreateLetterInput) (*ent.Letter, error) {
	return r.client.Letter.Create().
		SetContent(input.Content).
		Save(ctx)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Todo.Query().All(ctx)
}

func (r *queryResolver) Letters(ctx context.Context) ([]*ent.Letter, error) {
	return r.client.Letter.Query().All(ctx)
}

func (r *todoResolver) Done(ctx context.Context, obj *ent.Todo) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *ent.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Text(ctx context.Context, obj *ent.Todo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
