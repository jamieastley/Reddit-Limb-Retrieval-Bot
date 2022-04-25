package main

import (
	"api/graph"
	"api/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/jamieastley/limbretrievalbot/repository"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo, err := repository.NewRepository(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			BannedSubreddit: repo.BannedSubreddit,
		},
	}))

	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
