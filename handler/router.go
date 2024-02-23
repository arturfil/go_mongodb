package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


type Response struct {
    Msg  string
    Code int
}

func CreateRouter() http.Handler {

	router := chi.NewRouter()

    router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/healthcheck", HealthCheck)

	router.Get("/todos", GetTodos)
	router.Post("/todos/create", CreateTodo)

	return router
}
