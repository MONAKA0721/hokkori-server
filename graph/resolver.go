package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
	//rooms  map[int]map[int]chan<- *model.Message
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		//Resolvers: &Resolver{client, map[int]map[int]chan<- *model.Message{}},
		Resolvers: &Resolver{client},
	})
}
