package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gafeol/arte-de-rua/graphql"
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
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Printf("Listening on port %v\n", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
