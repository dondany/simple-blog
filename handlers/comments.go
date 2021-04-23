package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dondany/simple-blog/storage"
	"github.com/gorilla/mux"
)

type Comments struct {
	logger *log.Logger
}

type KeyComment struct{}

func NewComments(logger *log.Logger) *Comments {
	return &Comments{logger}
}

func (c *Comments) GetComments(rw http.ResponseWriter, request *http.Request) {
	c.logger.Println("Handles GET Comments")

	comments := storage.GetComments()
	err := comments.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
		return
	}
}

func (c *Comments) AddComment(rw http.ResponseWriter, request *http.Request) {
	c.logger.Println("Handles POST comment")

	comment := storage.Comment{}
	err := comment.FromJson(request.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}

	storage.AddComment(&comment)
}

func (c *Comments) DeleteComment(rw http.ResponseWriter, request *http.Request) {
	c.logger.Println("Handles DELETE comment")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	storage.DeleteComment(id)
}
