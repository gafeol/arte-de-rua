package main

import (
	"fmt"
	"net/http"

	"github.com/gafeol/arte-de-rua/server/models"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &models.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	port := 8000
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
