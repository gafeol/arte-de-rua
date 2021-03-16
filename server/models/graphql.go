package schemas

import (
	"errors"

	"github.com/gafeol/arte-de-rua/orm"
	"github.com/graphql-go/graphql"
)

type Art struct {
	ID     string
	Frase  string
	ImgURL string
}

var artType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Arte",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"frase": &graphql.Field{
				Type: graphql.String,
			},
			"imgURL": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

type Artist struct {
	ID   string
	Nome string
}

var artistType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Artista",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"nome": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"arts": &graphql.Field{
				Type: graphql.NewList(artType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					art, err := classes.AllArts()
					return art, err
				},
			},
			"art": &graphql.Field{
				Type: artType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					stringID, ok := p.Args["id"].(string)
					if ok {
						art, err := classes.FindArt(stringID)
						return art, err
					} else {
						return nil, errors.New("ID was not a string")
					}
				},
			},
			"artist": &graphql.Field{
				Type: artistType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idString, _ := p.Args["id"].(string)
					return classes.FindArtist(idString)
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addArt": &graphql.Field{
				Type: artType,
				Args: graphql.FieldConfigArgument{
					"frase": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"imgURL": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fraseString, _ := p.Args["frase"].(string)
					imgURLString, _ := p.Args["imgURL"].(string)
					art := &classes.Art{Frase: fraseString, ImgURL: imgURLString}
					if err := art.Create(); err != nil {
						return nil, err
					}
					return art, nil
				},
			},
			"addArtist": &graphql.Field{
				Type: artistType,
				Args: graphql.FieldConfigArgument{
					"nome": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					nomeString, _ := p.Args["nome"].(string)
					artist := &classes.Artist{Nome: nomeString}
					if err := artist.Create(); err != nil {
						return nil, err
					}
					return artist, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
