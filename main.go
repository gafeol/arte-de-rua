package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gafeol/arte-de-rua/graphql"
	"github.com/graphql-go/handler"
)

func addHeaderWrapper(h *handler.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		h.ServeHTTP(w, r)
	})
}

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schemas.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", addHeaderWrapper(h))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write([]byte("OK"))
	})
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Printf("Listening on port %v\n", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
