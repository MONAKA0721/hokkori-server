package main

import (
	"context"
	"fmt"
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

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("HOKKORI_DB_HOST"),
		os.Getenv("HOKKORI_DB_PORT"),
		os.Getenv("HOKKORI_DB_USER"),
		os.Getenv("HOKKORI_DB_DATABASE"),
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

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/",
		playground.Handler("Hokkori", "/query"),
	)
	if os.Getenv("HOKKORI_ENV") == "local" {
		http.Handle("/query", srv)
	} else {
		http.Handle("/query", middleware.EnsureValidToken()(srv))
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
