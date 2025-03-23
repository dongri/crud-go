package interfaces

import (
	"net/http"

	"github.com/dongri/crud-go/backend/app/infrastructures"
	"github.com/dongri/crud-go/backend/app/interfaces/handlers"
	"github.com/dongri/crud-go/backend/app/usecases"
	"github.com/dongri/crud-go/backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-gorp/gorp/v3"
)

var defaultMiddlewares = []func(next http.Handler) http.Handler{
	middleware.Logger,
	middleware.Recoverer,
	middleware.StripSlashes,
}

func NewRootHandler(env config.Env, dbmap *gorp.DbMap) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://localhost:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(defaultMiddlewares...)

	r.Get("/", handlers.HomeHandler)
	r.Get("/status", handlers.StatusHandler)

	r.Mount("/api", apiRouter(dbmap))

	return r
}

func apiRouter(dbmap *gorp.DbMap) http.Handler {
	infra := infrastructures.NewInfra(dbmap)
	useCase := usecases.NewUseeCase(infra)
	handler := handlers.NewHandler(useCase)

	ph := handler.PostHandler
	ch := handler.CommentHandler

	r := chi.NewRouter()

	r.Get("/posts", ph.ListHandler)
	r.Post("/posts", ph.NewHandler)
	r.Get("/posts/{id}", ph.ShowHandler)
	r.Post("/posts/{id}/update", ph.UpdateHandler)
	r.Post("/posts/{id}/delete", ph.DeleteHandler)

	r.Get("/posts/{id}/comments", ch.ListHandler)
	r.Post("/posts/{id}/comments", ch.NewHandler)

	return r
}
