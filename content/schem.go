package Content

import (
	"github.com/graphql-go/graphql"
)

type Content struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Text string `json:"text"`
}

var Type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Content",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Description: "contents author",
			Type: graphql.String,
		},
		"text" :&graphql.Field{
			Type: graphql.String,
		},
	},
})