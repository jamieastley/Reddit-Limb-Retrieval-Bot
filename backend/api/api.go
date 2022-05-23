package main

import (
	"api/graph"
	"api/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamieastley/limbretrievalbot/repository"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"
const graphQLEndpoint = "/graphql"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r, err := repository.NewRepository(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	s := CreateNewServer(&r)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			BannedSubredditHandler: s.BannedSubreddit,
			IgnoredUserHandler:     s.IgnoredUser,
		},
	}))
	srv.Use(extension.Introspection{})
	s.MountHandlers(srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))
}

type Server struct {
	Router *chi.Mux
	*repository.Repository
}

func CreateNewServer(r *repository.Repository) *Server {
	s := &Server{
		Router:     chi.NewRouter(),
		Repository: r,
	}
	return s
}

func (s *Server) MountHandlers(gql *handler.Server) {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", playground.Handler("GraphQL playground", graphQLEndpoint))
	s.Router.Post(graphQLEndpoint, gql.ServeHTTP)
}
