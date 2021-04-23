package app

import (
	"log"
	"net/http"

	"github.com/dondany/simple-blog/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Logger *log.Logger
}

func (app *App) Initialize() {
	app.initializeRoutes()
}

func (app *App) Run(host string) {
	app.Logger.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) initializeRoutes() {
	postsHandler := handlers.NewPosts(app.Logger)
	app.Router.HandleFunc("/posts", postsHandler.GetPosts).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postsHandler.AddPost).Methods(http.MethodPost)
	app.Router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.DeletePost).Methods(http.MethodDelete)

	commentsHandler := handlers.NewComments(app.Logger)
	app.Router.HandleFunc("/comments", commentsHandler.GetComments).Methods(http.MethodGet)
	app.Router.HandleFunc("/comments", commentsHandler.AddComment).Methods(http.MethodPost)
	app.Router.HandleFunc("/comments/{id:[0-9]+}", commentsHandler.DeleteComment).Methods(http.MethodDelete)
}
