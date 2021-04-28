package domain

import "context"

type Comment struct {
	Id      int64
	PostId  int64
	Author  string
	Content string
}

type CommentRepository interface {
	FindAll(ctx context.Context) ([]Comment, error)
	Find(ctx context.Context, id int64) (Comment, error)
}
