package app

import (
	"log"
	"net/http"

	"github.com/dondany/simple-blog/domain"
	"github.com/dondany/simple-blog/handlers"
	"github.com/gorilla/mux"
)

type app struct {
	Router   *mux.Router
	Logger   *log.Logger
	PostRepo *domain.PostRepository
}

func Http(r *mux.Router, l *log.Logger, postRepo *domain.PostRepository) app {
	return app{
		Router:   r,
		Logger:   l,
		PostRepo: postRepo,
	}
}

func (app *app) Initialize() {
	app.initializeRoutes()
}

func (app *app) Run(host string) {
	app.Logger.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *app) initializeRoutes() {
	postsHandler := handlers.NewPosts(app.Logger, app.PostRepo)
	app.Router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.Find).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.Delete).Methods(http.MethodDelete)
	app.Router.HandleFunc("/posts", postsHandler.FindAll).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postsHandler.Create).Methods(http.MethodPost)
	app.Router.HandleFunc("/posts/{id:[0-9]+}/comments", postsHandler.GetComments).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts/{id:[0-9]+}/comments", postsHandler.AddComment).Methods(http.MethodPost)
}
