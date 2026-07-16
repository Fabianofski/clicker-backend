package router

import (
	"net/http"
	"os"

	"f4b1.dev/clicker-backend/internal/handlers"
	"f4b1.dev/clicker-backend/internal/middleware"
	"f4b1.dev/clicker-backend/internal/service"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func New(userService *service.UserService) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	jwt_secret := os.Getenv("JWT_SECRET")
	userHandler := handlers.NewUserHandler(userService, jwt_secret)

	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", userHandler.SignUp)
				r.Post("/login", userHandler.Login)
				r.Group(func(r chi.Router) {
					r.Use(middleware.JWTMiddleware(jwt_secret))
					r.Delete("/delete", userHandler.DeleteAccount)
				})
			})
			r.Route("/leaderboard", func(r chi.Router) {
				leaderboardHandler := handlers.NewLeaderboardHandler()
				r.Get("/", leaderboardHandler.GetLeaderboard)
			})

			r.Group(func(r chi.Router) {
				r.Use(middleware.JWTMiddleware(jwt_secret))

				r.Route("/game", func(r chi.Router) {
					gameHandler := handlers.NewGameHandler()
					r.Get("/state", gameHandler.GetGameState)
				})

				r.Route("/store", func(r chi.Router) {
					storeHandler := handlers.NewStoreHandler()
					r.Post("/purchase", storeHandler.Purchase)
				})
			})
		})
	})

	return r
}
