package router

import (
	"net/http"

	"f4b1.dev/clicker-backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Route("/api", func(r chi.Router) {

			r.Get("/swagger/*", httpSwagger.WrapHandler)
		})
	})

	return r
}
