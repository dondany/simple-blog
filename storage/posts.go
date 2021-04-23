package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type Post struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Posts struct {
	Posts []*Post
}

var postList = []*Post{
	&Post{
		1,
		"The First Post",
		"Content of the very first post",
	},
	&Post{
		2,
		"The Second Post",
		"Wow, another post!",
	},
}

func (p *Posts) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)
}

func (p *Post) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)
}

func GetPosts() Posts {
	return Posts{postList}
}

func AddPost(p *Post) {
	p.Id = getNextId()
	postList = append(postList, p)
}

func DeletePost(id int) {
	for index, post := range postList {
		fmt.Println(id, post.Id)
		if post.Id == id {
			postList = append(postList[:index], postList[index+1:]...)
			return
		}
	}
}

func getNextId() int {
	post := postList[len(postList)-1]
	return post.Id + 1
}
