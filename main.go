package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger/v2"

	_ "f4b1.dev/clicker-backend/docs"
)

//	@title			Clicker Backend
//	@version		1.0
//	@description	Clicker Backend in Golang

//	@contact.name	Fabian
//	@contact.email	hello@f4b1.dev

// @host		f4b1.dev/clicker-backend
// @BasePath	/v2
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":3000", r)
}
