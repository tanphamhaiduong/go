package modules

import (
	"github.com/graphql-go/graphql"
)

var (
	// rootQuery ...
	rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: graphql.Fields{},
	})
	// rootMutation ...
	rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: graphql.Fields{},
	})
)

// MakeSchema ...
func MakeSchema(resolver Resolver) (graphql.Schema, error) {
	addToSchema(resolver)
	addCustomToSchema(resolver)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		return graphql.Schema{}, err
	}
	return schema, nil
}
