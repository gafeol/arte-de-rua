package models

import "github.com/graphql-go/graphql"

type Art struct {
	ID    int64  `json:"id"`
	Frase string `json:"frase"`
}

var ArtType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Arte",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"frase": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"test": &graphql.Field{
				Type: ArtType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Art{
						ID:    1,
						Frase: "Frase de teste",
					}, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: QueryType,
	},
)
