package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dondany/simple-blog/domain"
	"github.com/gorilla/mux"
)

type Posts struct {
	logger   *log.Logger
	postRepo domain.PostRepository
}

type KeyPost struct{}

func NewPosts(logger *log.Logger, postRepo *domain.PostRepository) *Posts {
	return &Posts{logger, *postRepo}
}

func (p *Posts) Get(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	result, err := p.postRepo.Get(request.Context(), int64(id))
	if err != nil {
		http.Error(rw, "Could not fetch the post", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

func (p *Posts) GetAll(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling GetAll posts")
	result, err := p.postRepo.GetAll(request.Context())
	if err != nil {
		http.Error(rw, "Unable to fetch all posts", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

// func (p *Posts) GetPosts(rw http.ResponseWriter, request *http.Request) {
// 	p.logger.Println("Handles GET Posts")

// 	posts := storage.GetPosts()
// 	err := posts.ToJson(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (p *Posts) AddPost(rw http.ResponseWriter, request *http.Request) {
// 	p.logger.Println("Handles POST Post")

// 	post := storage.Post{}
// 	err := post.FromJson(request.Body)
// 	if err != nil {
// 		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
// 		return
// 	}
// 	storage.AddPost(&post)
// }

// func (p *Posts) DeletePost(rw http.ResponseWriter, request *http.Request) {
// 	p.logger.Println("Handles DELETE Post")

// 	vars := mux.Vars(request)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
// 		return
// 	}

// 	storage.DeletePost(id)

//}
