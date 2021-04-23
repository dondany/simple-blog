package main

import (
	"log"
	"os"

	"github.com/dondany/simple-blog/app"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "simple-blog-api", log.LstdFlags)
	logger.Println("simple-blog-api project started")

	app := app.App{Router: mux.NewRouter(), Logger: logger}
	app.Initialize()
	app.Run(":8080")

}
