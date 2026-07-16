package main

import (
	"context"
	"log/slog"
	"net/http"

	_ "f4b1.dev/clicker-backend/docs"
	"f4b1.dev/clicker-backend/ent"
	"f4b1.dev/clicker-backend/internal/repository"
	"f4b1.dev/clicker-backend/internal/router"
	"f4b1.dev/clicker-backend/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

//	@title			Clicker Backend
//	@version		1.0
//	@description	Clicker Backend in Golang

// @BasePath	/api
func main() {
	client, err := ent.Open("sqlite3", "./clicker.sqlite?cache=shared&_fk=1")
	if err != nil {
		slog.Error("failed opening connection to sqlite:", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		slog.Error("failed creating schema resources:", err)
	}

	entUserRepo := repository.NewEntUserRepository(client)
	userService := service.NewUserService(entUserRepo)
	buildingService, err := service.NewBuildingService()
	if err != nil {
		panic(err)
	}
	slog.Info("Buildings from csv: ", buildingService.Buildings)

	r := router.New(userService)
	http.ListenAndServe(":3000", r)
	slog.Info("Server listening on Port 3000")
}
