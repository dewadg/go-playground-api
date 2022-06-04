package gql

import (
	_ "embed"

	"github.com/dewadg/go-playground-api/internal/app/gql/resolver"
	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//go:embed schema.graphql
var schemaString string

func Register(router chi.Router) error {
	r := resolver.New()

	schema, err := graphql.ParseSchema(schemaString, r, graphql.UseFieldResolvers())
	if err != nil {
		return err
	}

	router.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})

	return nil
}
