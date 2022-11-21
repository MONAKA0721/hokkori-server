package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/env"
	"github.com/MONAKA0721/hokkori/graph/generated"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	config   *env.Config
	client   *ent.Client
	uploader *manager.Uploader
	//rooms  map[int]map[int]chan<- *model.Message
}

// NewSchema creates a graphql executable schema.
func NewSchema(config *env.Config, client *ent.Client, uploader *manager.Uploader) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		//Resolvers: &Resolver{client, map[int]map[int]chan<- *model.Message{}},
		Resolvers: &Resolver{config, client, uploader},
	})
}
