package main

import (
	"fmt"
)

func Test() {
	post := Post{Content: "Hello Golang", Author: "shiddong"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Lau", Post: &post}
	comment.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Printf("post with comments: %+v", readPost)
}
