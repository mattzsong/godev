package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mattzsong/godev-backend/config"
	"github.com/mattzsong/godev-backend/handlers"
	"github.com/mattzsong/godev-backend/pkg/shutdown"
	"github.com/mattzsong/godev-backend/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type apiClient struct {
	db *mongo.Database
}

func main() {
	var exitCode int
	defer os.Exit(exitCode)

	//load config
	env, err := config.LoadConfig()
	if err != nil {
		log.Printf("Config failed to load: %s\n", err)
		exitCode = 1
		return
	}

	//run server with db and handler
	cleanup, err := run(env)
	if err != nil {
		log.Printf("Server failed to start: %s", err)
		exitCode = 1
		return
	}
	defer cleanup()

	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	router, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}
	srv := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: router,
	}

	go func() {
		log.Printf("Serving on port: %s\n", env.PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return func() {
		cleanup()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			cancel()
		}()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown failed: %s", err)
		}
		log.Print("Server Exited Properly")
	}, nil

}

func buildServer(env config.EnvVars) (chi.Router, func(), error) {
	mongodb, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}

	router := initUserRouter()
	// router.Mount("/admin", adminRouter())

	return router, func() {
		storage.CloseMongo(mongodb)
	}, nil
}

func initUserRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Timeout(60 * time.Second))

	//test routes for server
	router.Get("/health", handlers.HandlerReadiness)
	router.Get("/err", handlers.HandlerErr)

	//swagger handling
	// router.Get("/swagger/*", swagger.HandlerDefault)

	//routes for recipe lists

	return router
}

// func adminRouter() http.Handler {
// 	router := chi.NewRouter()
// 	// router.Use(AdminOnly)
// }

// func adminOnly(r *http.Request) http.Handler {

// }
