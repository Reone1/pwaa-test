package User

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var Type = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
