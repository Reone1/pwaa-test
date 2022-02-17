package main

import (
	"net/http"

	"pwaa-test.com/graphql"
)
func main (){
	http.Handle("/graphql", graphql.SetGraphqlHandler())

	http.ListenAndServe(":8080", nil)
}