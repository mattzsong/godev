package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env
	godotenv.Load()

	// Grab port from environment
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in .env")
	} else {
		fmt.Println(portString)
	}

	// create main router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://", "http://"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create RIPv1 router for testing
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	// Create server and handle requests
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
