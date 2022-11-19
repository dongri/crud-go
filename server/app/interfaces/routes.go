package interfaces

import (
	"net/http"

	"github.com/dongri/crud-go/server/app/infrastructures"
	"github.com/dongri/crud-go/server/app/interfaces/request"
	"github.com/dongri/crud-go/server/app/usecases"
	"github.com/dongri/crud-go/server/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"gopkg.in/gorp.v1"
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

	r.Get("/", request.HomeHandler)
	r.Get("/status", request.StatusHandler)

	r.Mount("/api", apiRouter(dbmap))

	return r
}

func apiRouter(dbmap *gorp.DbMap) http.Handler {
	postInfra := infrastructures.NewPost(dbmap)
	postUseCase := usecases.NewPost(postInfra)
	postReq := request.NewPostRequest(postUseCase)

	r := chi.NewRouter()
	r.Get("/posts", postReq.ListHandler)
	r.Get("/posts/{id}", postReq.ShowHandler)
	r.Post("/posts/new", postReq.NewHandler)
	r.Post("/posts/{id}/update", postReq.UpdateHandler)
	r.Post("/posts/{id}/delete", postReq.DeleteHandler)

	return r
}
