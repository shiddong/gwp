package main

import "fmt"

// Test 测试CRUD
func Test() {
	post := Post{
		Content: "Hello Golang",
		Author:  "shiddong",
	}
	fmt.Printf("post: %+v\n", post)

	post.Create()
	fmt.Printf("post: %+v\n", post)

	readPost, _ := GetPost(post.Id)
	fmt.Printf("readPost: %+v\n", readPost)

	readPost.Content = "Bonjour Monde"
	readPost.Author = "Pierre"
	err := readPost.Update()
	if err != nil {
		panic(err)
	}

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
	posts, _ = Posts(10)
	fmt.Println(posts)
}
