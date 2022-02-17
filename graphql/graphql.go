package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	Content "pwaa-test.com/content"
	User "pwaa-test.com/user"
)

func init() {
	// set graphql field & type
	// bind resolver
}

func SetGraphqlHandler() *handler.Handler {
  
	fields := graphql.Fields{
		"user": &graphql.Field{
			Type: graphql.NewList(User.Type),
			Resolve: func(p graphql.ResolveParams) (interface {} , error) {
				return User.GetUsers(), nil
			},
		},
		"content" :&graphql.Field{
			Type: Content.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Content.Content{
					ID: "content-id",
					Title: "content-title",
					Author :"content-author",
					Text:"content-text",
				}, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})

	config := &handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: false,
		Playground: true,
	}

	h := handler.New(config)
	return h
}