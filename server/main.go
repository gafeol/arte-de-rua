package main

import (
	"fmt"
	"net/http"

	"github.com/gafeol/arte-de-rua/models"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schemas.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	port := 8000
	panic(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
