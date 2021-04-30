# simple-blog
Simple blog Rest API. 

## Currently implemented endpoints

`GET /posts` - returns all the posts

`GET /posts/{id}` - returns the post with the given id

`POST /posts/` - create a post from request body

`DELETE /posts/{id}` - delete the post with the given id

`GET /posts/{id}/comments` - returns all comments under the given post

`POST /posts/{id}/comments` - create a comment under the given post