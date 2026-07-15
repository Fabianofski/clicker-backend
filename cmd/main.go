package main

import (
	"fmt"
	"net/http"

	_ "f4b1.dev/clicker-backend/docs"
	"f4b1.dev/clicker-backend/internal/router"
)

//	@title			Clicker Backend
//	@version		1.0
//	@description	Clicker Backend in Golang

//	@contact.name	Fabian
//	@contact.email	hello@f4b1.dev

// @host		f4b1.dev/clicker-backend
// @BasePath	/v2
func main() {
	r := router.New()
	http.ListenAndServe(":3000", r)
	fmt.Printf("Server listening on Port 3000")
}
