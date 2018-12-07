package graphql

import (
	"api/db"
	"api/models"

	"fmt"
	"log"
	"strconv"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"hobby": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, err := strconv.ParseUint(p.Args["id"].(string), 10, 64)
					if err != nil {
						return nil, nil
					}

					db := db.ConnectGORM()
					user := models.User{}
					user.ID = uint(idQuery)
					db.First(&user)
					log.Print(idQuery)
					return &user, nil
				},
			},
		},
	},
)

func ExecuteQuery(query string) *graphql.Result {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("unexpected errors: %v", result.Errors)
	}

	return result
}
