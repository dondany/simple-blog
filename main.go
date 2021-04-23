package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dondany/simple-blog/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "simple-blog-api", log.LstdFlags)
	logger.Println("simple-blog-api project started")

	router := mux.NewRouter()

	postsHandler := handlers.NewPosts(logger)
	router.HandleFunc("/posts", postsHandler.GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", postsHandler.AddPost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id:[0-9]+}", postsHandler.DeletePost).Methods(http.MethodDelete)

	commentsHandler := handlers.NewComments(logger)
	router.HandleFunc("/comments", commentsHandler.GetComments).Methods(http.MethodGet)
	router.HandleFunc("/comments", commentsHandler.AddComment).Methods(http.MethodPost)
	router.HandleFunc("/comments/{id:[0-9]+}", commentsHandler.DeleteComment).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))

}
