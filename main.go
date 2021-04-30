package main

import (
	"log"
	"os"

	"github.com/dondany/simple-blog/app"
	repository "github.com/dondany/simple-blog/repository/postgresql"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "simple-blog-api", log.LstdFlags)
	logger.Println("simple-blog-api project started")

	postRepo := repository.PostgresqlPost()

	app := app.Http(mux.NewRouter(), logger, &postRepo)
	app.Initialize()
	app.Run(":8080")
}
