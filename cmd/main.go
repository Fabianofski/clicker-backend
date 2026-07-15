package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"

	_ "f4b1.dev/clicker-backend/docs"
	"f4b1.dev/clicker-backend/internal/repository"
	"f4b1.dev/clicker-backend/internal/router"
	"f4b1.dev/clicker-backend/internal/service"
)

//	@title			Clicker Backend
//	@version		1.0
//	@description	Clicker Backend in Golang

//	@contact.name	Fabian
//	@contact.email	hello@f4b1.dev

// @BasePath	/api
func main() {

	db, err := sql.Open("sqlite", "./clicker.sqlite")
	if err != nil {
		panic(err)
	}
	sqlUserRepo := repository.NewSQLUserRepository(db)
	userService := service.NewUserService(sqlUserRepo)

	r := router.New(userService)
	http.ListenAndServe(":3000", r)
	fmt.Printf("Server listening on Port 3000")
}
