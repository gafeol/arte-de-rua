package schemas

import (
	"errors"
	"strconv"

	"github.com/gafeol/arte-de-rua/orm"
	"github.com/graphql-go/graphql"
)

type Art struct {
	ID       uint64
	Frase    string
	ImgURL   string
	ArtistID uint64
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
			"artist": &graphql.Field{
				Type: artistType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					art := p.Source.(orm.Art)
					return orm.FindArtist(art.ArtistID)
				},
			},
		},
	},
)

type Artist struct {
	ID   uint64
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

func init() {
	artistType.AddFieldConfig("arts", &graphql.Field{
		Type: graphql.NewList(artType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			artist := p.Source.(orm.Artist)
			return orm.FindArtByArtist(artist.ID)
		},
	})
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"arts": &graphql.Field{
				Type: graphql.NewList(artType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					art, err := orm.AllArts()
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
					stringID, ok := p.Args["id"].(uint64)
					if ok {
						art, err := orm.FindArt(stringID)
						return art, err
					} else {
						return nil, errors.New("ID was not a string")
					}
				},
			},
			"artists": &graphql.Field{
				Type: graphql.NewList(artistType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return orm.AllArtists()
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
					id := p.Args["id"].(uint64)
					return orm.FindArtist(id)
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
					"artistID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fraseString := p.Args["frase"].(string)
					imgURLString := p.Args["imgURL"].(string)
					artistIDString := p.Args["artistID"].(string)
					artistID, err := strconv.ParseUint(artistIDString, 10, 64)
					if err != nil {
						panic(err)
					}
					art := &orm.Art{Frase: fraseString, ImgURL: imgURLString, ArtistID: artistID}
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
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					nomeString, _ := p.Args["nome"].(string)
					artist := &orm.Artist{Nome: nomeString}
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
