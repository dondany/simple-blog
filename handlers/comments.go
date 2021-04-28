package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dondany/simple-blog/domain"
	"github.com/gorilla/mux"
)

type Comments struct {
	logger      *log.Logger
	commentRepo domain.CommentRepository
}

type KeyComment struct{}

func NewComments(logger *log.Logger, commentRepo *domain.CommentRepository) *Comments {
	return &Comments{logger, *commentRepo}
}

func (c *Comments) Find(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	result, err := c.commentRepo.Find(request.Context(), int64(id))
	if err != nil {
		http.Error(rw, "Could not fetch the comment", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

func (p *Comments) FindAll(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling GetAll posts")
	result, err := p.commentRepo.FindAll(request.Context())
	if err != nil {
		http.Error(rw, "Unable to fetch all comments", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

// func (c *Comments) GetComments(rw http.ResponseWriter, request *http.Request) {
// 	c.logger.Println("Handles GET Comments")

// 	comments := storage.GetComments()
// 	err := comments.ToJson(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (c *Comments) AddComment(rw http.ResponseWriter, request *http.Request) {
// 	c.logger.Println("Handles POST comment")

// 	comment := storage.Comment{}
// 	err := comment.FromJson(request.Body)
// 	if err != nil {
// 		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
// 		return
// 	}

// 	storage.AddComment(&comment)
// }

// func (c *Comments) DeleteComment(rw http.ResponseWriter, request *http.Request) {
// 	c.logger.Println("Handles DELETE comment")

// 	vars := mux.Vars(request)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
// 		return
// 	}

// 	storage.DeleteComment(id)
// }
