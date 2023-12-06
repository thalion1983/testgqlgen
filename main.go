package main

import (
	"log"
	"net/http"
	"testgqlgen/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const port = "8080"

// Mocked Database
var (
	users []*graph.User
	maxID int = 0
)

func main() {
	resolver := &graph.Resolver{
		UserList: users,
		MaxID: &maxID,
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
