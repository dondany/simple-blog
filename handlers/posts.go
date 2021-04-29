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

func (p *Posts) Find(rw http.ResponseWriter, request *http.Request) {
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

func (p *Posts) FindAll(rw http.ResponseWriter, request *http.Request) {
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

func (p *Posts) Create(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling Create post")
	result := domain.Post{}
	err := json.NewDecoder(request.Body).Decode(&result)
	if err != nil {
		http.Error(rw, "Unable to unmarshal request body", http.StatusBadRequest)
		return
	}
	p.postRepo.Create(request.Context(), &result)
	rw.WriteHeader(http.StatusCreated)
}

func (p *Posts) Delete(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling Delete post")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.postRepo.Delete(request.Context(), int64(id))
	rw.WriteHeader(http.StatusOK)
}
