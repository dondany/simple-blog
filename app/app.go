package app

import (
	"log"
	"net/http"

	"github.com/dondany/simple-blog/domain"
	"github.com/dondany/simple-blog/handlers"
	"github.com/dondany/simple-blog/storage/repository/postgresql"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Logger   *log.Logger
	PostRepo *domain.PostRepository
}

func (app *App) Initialize() {
	app.initializeRoutes()
}

func (app *App) Run(host string) {
	app.Logger.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) initializeRoutes() {
	postRepo := postgresql.NewPostresqlPostRepository()
	postsHandler := handlers.NewPosts(app.Logger, &postRepo)
	app.Router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.Get).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postsHandler.GetAll).Methods(http.MethodGet)

	commentsRepo := postgresql.NewPostresqlCommentRepository()
	commentsHandler := handlers.NewComments(app.Logger, &commentsRepo)
	app.Router.HandleFunc("/comments", commentsHandler.GetAll).Methods(http.MethodGet)
	app.Router.HandleFunc("/comments/{id:[0-9]+}", commentsHandler.Get).Methods(http.MethodGet)
}
