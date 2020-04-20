package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ataleksand/gql/storage"

	"github.com/ataleksand/gql/models"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/graphql-go/graphql"
)

type Api struct {
	queryType *graphql.Object
	userType  *graphql.Object
	schema    graphql.Schema
}

func NewApi(storage storage.Storage) *Api {
	api := &Api{}

	api.userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	api.queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: api.userType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(string)
						if ok {
							return models.Users(qm.Where("id=?", id)).One(context.Background(), storage.Get())
						}
						return nil, nil
					},
				},
			},
		},
	)

	api.schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: api.queryType,
		},
	)
	return api
}

type GqlHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (gf GqlHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gf(w, r)
}

func (api *Api) Gqlhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	res := graphql.Do(graphql.Params{
		Schema:        api.schema,
		RequestString: query,
	})

	if len(res.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", res.Errors)
	}

	_ = json.NewEncoder(w).Encode(res)
}
