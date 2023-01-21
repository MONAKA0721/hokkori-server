package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/MONAKA0721/hokkori/env"
	"github.com/MONAKA0721/hokkori/s3"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/ent/migrate"
	"github.com/MONAKA0721/hokkori/graph"
	"github.com/MONAKA0721/hokkori/middleware"

	_ "github.com/lib/pq"
)

func main() {
	environment := os.Getenv("HOKKORI_ENV")
	if environment == "" {
		log.Fatal("HOKKORI_ENV is not set")
	}

	cfg := env.Config{
		S3: s3.Config{
			EndpointURL: "https://s3.ap-northeast-1.amazonaws.com",
			BucketName:  "hokkori-" + environment,
		},
	}

	// Create ent.Client and run the schema migration.
	var host string
	switch environment {
	case "dev":
		host = "db"
	default:
		host = "host.docker.internal"
	}
	client, err := ent.Open(dialect.Postgres, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		"5432",
		"postgres",
		"hokkori",
		os.Getenv("HOKKORI_DB_PASSWORD"),
	))
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s3Uploader, err := s3.NewUploader(ctx, cfg.S3)
	if err != nil {
		log.Fatal("initializing s3 client")
	}

	srv := handler.NewDefaultServer(graph.NewSchema(&cfg, client, s3Uploader))
	// クエリを吐く設定
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		fmt.Printf("around: %s %s %s", oc.OperationName, oc.RawQuery, oc.Variables)
		return next(ctx)
	})

	http.Handle("/",
		playground.Handler("Hokkori", "/query"),
	)
	if environment == "dev" {
		http.Handle("/query", srv)
	} else {
		http.Handle("/query", middleware.EnsureValidToken(environment)(srv))
	}
	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
