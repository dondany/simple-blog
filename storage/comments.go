package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type Comment struct {
	Id      int    `json:"id"`
	PostId  int    `json:"postId"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Comments struct {
	Comments []*Comment
}

var commentList = []*Comment{
	&Comment{
		1,
		1,
		"Nice article",
		"Marc",
	},
	&Comment{
		2,
		1,
		"Wow, finally something",
		"Ana",
	},
	&Comment{
		3,
		2,
		"Getting somewhere",
		"Bob",
	},
}

func (c *Comments) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(c)
}

func (c *Comment) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(c)
}

func GetComments() Comments {
	return Comments{commentList}
}

func AddComment(c *Comment) {
	c.Id = getNextCommentId()
	commentList = append(commentList, c)
}

func DeleteComment(id int) {
	for index, comment := range commentList {
		fmt.Println(id, comment.Id)
		if comment.Id == id {
			commentList = append(commentList[:index], commentList[index+1:]...)
			return
		}
	}
}

func getNextCommentId() int {
	comment := commentList[len(commentList)-1]
	return comment.Id + 1
}
