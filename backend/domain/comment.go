package domain

type Comment struct {
	Id      int64
	PostId  int64
	Author  string
	Content string
}
