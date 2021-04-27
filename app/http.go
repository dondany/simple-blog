package app

import (
	"log"
	"net/http"

	"github.com/dondany/simple-blog/domain"
	"github.com/dondany/simple-blog/handlers"
	"github.com/gorilla/mux"
)

type app struct {
	Router      *mux.Router
	Logger      *log.Logger
	PostRepo    *domain.PostRepository
	CommentRepo *domain.CommentRepository
}

func Http(r *mux.Router, l *log.Logger, postRepo *domain.PostRepository, commentRepo *domain.CommentRepository) app {
	return app{
		Router:      r,
		Logger:      l,
		PostRepo:    postRepo,
		CommentRepo: commentRepo,
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
	app.Router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.Get).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postsHandler.GetAll).Methods(http.MethodGet)

	commentsHandler := handlers.NewComments(app.Logger, app.CommentRepo)
	app.Router.HandleFunc("/comments", commentsHandler.GetAll).Methods(http.MethodGet)
	app.Router.HandleFunc("/comments/{id:[0-9]+}", commentsHandler.Get).Methods(http.MethodGet)
}
