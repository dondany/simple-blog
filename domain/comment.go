package domain

import "context"

type Comment struct {
	Id      int64
	PostId  int64
	Author  string
	Content string
}

type CommentRepository interface {
	GetAll(ctx context.Context) ([]Comment, error)
	Get(ctx context.Context, id int64) (Comment, error)
}
