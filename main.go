package main

import (
	"log"
	"net/http"
	"testgqlgen/graph"
	"testgqlgen/graph/generated"
	"testgqlgen/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const port = "8080"

// Mocked Database
var (
	users  []*model.User
	LastID int = 0
	apps   []*model.App
)

func main() {
	resolver := &graph.Resolver{
		UserList: &graph.Users{
			List:   users,
			LastID: &LastID,
		},
		AppList: apps,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
