package main

import (
	"log"
	"os"

	"github.com/brunoquindeler/go-with-tests/reading_files/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
