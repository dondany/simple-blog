package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dondany/simple-blog/storage"
	"github.com/gorilla/mux"
)

type Posts struct {
	logger *log.Logger
}

type KeyPost struct{}

func NewPosts(logger *log.Logger) *Posts {
	return &Posts{logger}
}

func (p *Posts) GetPosts(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handles GET Posts")

	posts := storage.GetPosts()
	err := posts.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
		return
	}
}

func (p *Posts) AddPost(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handles POST Post")

	post := storage.Post{}
	err := post.FromJson(request.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}
	storage.AddPost(&post)
}

func (p *Posts) DeletePost(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handles DELETE Post")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	storage.DeletePost(id)

}
