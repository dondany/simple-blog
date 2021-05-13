package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dondany/simple-blog/backend/domain"
	"github.com/gorilla/mux"
)

type Posts struct {
	logger   *log.Logger
	postRepo domain.PostRepository
}

type PostResponse struct {
	Id            int64
	Title         string
	Content       string
	Likes         int64
	CommentsCount int
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
	post, err := p.postRepo.Get(request.Context(), int64(id))
	if err != nil {
		http.Error(rw, "Could not fetch the post", http.StatusInternalServerError)
		return
	}

	result, err := p.mapPostResponse(request.Context(), post)
	if err != nil {
		http.Error(rw, "Unable to map post", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

func (p *Posts) FindAll(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling GetAll posts")
	posts, err := p.postRepo.GetAll(request.Context())
	if err != nil {
		http.Error(rw, "Unable to fetch all posts", http.StatusBadRequest)
		return
	}

	var enrichedPosts []*PostResponse
	for _, post := range posts {
		enrichedPost, err := p.mapPostResponse(request.Context(), post)
		if err != nil {
			http.Error(rw, "Unable to map post", http.StatusInternalServerError)
			return
		}
		enrichedPosts = append(enrichedPosts, enrichedPost)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(enrichedPosts)
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

func (p *Posts) GetComments(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling Get comments of the post")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	result, err := p.postRepo.GetComments(request.Context(), int64(id))
	if err != nil {
		http.Error(rw, "Unable to fetch all comments for the post", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

func (p *Posts) AddComment(rw http.ResponseWriter, request *http.Request) {
	setupCors(&rw, request)
	if request.Method == "OPTIONS" {
		return
	}
	p.logger.Println("Handling Add comment for the post")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	result := domain.Comment{}
	err = json.NewDecoder(request.Body).Decode(&result)
	if err != nil {
		http.Error(rw, "Unable to unmarshal request body", http.StatusBadRequest)
		return
	}
	result.PostId = int64(id)
	p.postRepo.CreateComment(request.Context(), result)
	rw.WriteHeader(http.StatusCreated)
}

func (p *Posts) AddLike(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling Add Like for post")
	p.handleTogglLike(rw, request, true)
}

func (p *Posts) DeleteLike(rw http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handling Delete Like for post")
	p.handleTogglLike(rw, request, false)
}

func (p *Posts) handleTogglLike(rw http.ResponseWriter, request *http.Request, toggle bool) {
	setupCors(&rw, request)
	if request.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.postRepo.LikePost(request.Context(), int64(id), toggle)

	rw.WriteHeader(http.StatusCreated)
}

func setupCors(rw *http.ResponseWriter, request *http.Request) {
	(*rw).Header().Set("Access-Control-Allow-Origin", "*")
	(*rw).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*rw).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (p *Posts) mapPostResponse(ctx context.Context, post domain.Post) (*PostResponse, error) {
	count, err := p.postRepo.GetCommentsCount(ctx, post.Id)
	if err != nil {
		return nil, err
	}
	return &PostResponse{Id: post.Id, Title: post.Title, Content: post.Content, Likes: post.Likes, CommentsCount: count}, nil
}
