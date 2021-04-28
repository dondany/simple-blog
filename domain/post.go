package domain

import "context"

type Post struct {
	Id      int64
	Title   string
	Content string
}

type PostRepository interface {
	FindAll(ctx context.Context) ([]Post, error)
	Find(ctx context.Context, id int64) (Post, error)
	Create(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int64) error
}
