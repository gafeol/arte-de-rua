package schemas

import (
	"errors"
	"strconv"

	"github.com/gafeol/arte-de-rua/orm"
	"github.com/graphql-go/graphql"
)

type Art struct {
	ID       uint64
	Phrase   string
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
			"phrase": &graphql.Field{
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
	Name string
}

var artistType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Artista",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
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
	artistType.AddFieldConfig("nArts", &graphql.Field{
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			artist := p.Source.(orm.Artist)
			arts, err := orm.FindArtByArtist(artist.ID)
			return len(arts), err
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
					stringID, ok := p.Args["id"].(string)

					if ok {
						id, err := strconv.ParseUint(stringID, 10, 64)
						art, err := orm.FindArt(id)
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
					idString := p.Args["id"].(string)
					id, err := strconv.ParseUint(idString, 10, 64)
					if err != nil {
						return nil, err
					}
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
					"phrase": &graphql.ArgumentConfig{
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
					phraseString := p.Args["phrase"].(string)
					imgURLString := p.Args["imgURL"].(string)
					artistIDString := p.Args["artistID"].(string)
					artistID, err := strconv.ParseUint(artistIDString, 10, 64)
					if err != nil {
						panic(err)
					}
					art := &orm.Art{Phrase: phraseString, ImgURL: imgURLString, ArtistID: artistID}
					if err := art.Create(); err != nil {
						return nil, err
					}
					return art, nil
				},
			},
			"addArtist": &graphql.Field{
				Type: artistType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					nameString, _ := p.Args["name"].(string)
					artist := &orm.Artist{Name: nameString}
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
