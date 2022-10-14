package main

import (
	"compada/person-service/graph"
	"compada/person-service/graph/generated"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// https://gqlgen.com/recipes/cors/
func main() {
	router := chi.NewRouter()
	server_port := os.Getenv("PORT")


	if cors_origins, present := os.LookupEnv("CORS_ORIGINS"); present {
		// Add CORS middleware around every request
		// See https://github.com/rs/cors for full option listing
		router.Use(cors.New(cors.Options{
			AllowedOrigins:   strings.Fields(cors_origins),
			AllowCredentials: true,
			Debug:            true,
		}).Handler)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/query", srv)

	log.Println("listening on port", server_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server_port), router))
}
