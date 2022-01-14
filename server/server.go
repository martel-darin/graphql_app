package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/martel-darin/graphql_app/db"
	"github.com/martel-darin/graphql_app/graph"
	"github.com/martel-darin/graphql_app/graph/generated"
)

const defaultPort = "5679"

func main() {
	if err := db.InitDB(); err != nil {
		panic(err)
	} else {
		defer db.ConnectionPool.Close(context.Background())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
