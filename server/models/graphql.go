package schemas

import "fmt"
import "errors"
import "github.com/graphql-go/graphql"
import "github.com/gafeol/arte-de-rua/server/orm"

type Art struct {
	ID     string `json:"id"`
	Frase  string `json:"frase"`
	ImgURL string `json:"imgURL"`
}

var dummyArts = []Art{
	Art{
		ID:     "1",
		Frase:  "Arte de id 1",
		ImgURL: "lkasjdlkasjdl",
	},
	Art{
		ID:     "2",
		Frase:  "Arte de id 2",
		ImgURL: "blboboblbobl",
	},
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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"test": &graphql.Field{
				Type: artType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					stringID, ok := p.Args["id"].(string)
					if ok {
						fmt.Println("Ok searching for id ", stringID)
						art, err := classes.FindArt(stringID)
						return art, err
					} else {
						return nil, errors.New("ID was not a string")
					}
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
