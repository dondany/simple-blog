package domain

import "context"

type Post struct {
	Id      int64
	Title   string
	Content string
}

type PostRepository interface {
	GetAll(ctx context.Context) ([]Post, error)
	Get(ctx context.Context, id int64) (Post, error)
	Create(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int64) error
	GetComments(ctx context.Context, id int64) ([]Comment, error)
}
