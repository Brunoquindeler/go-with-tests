package main

import (
	"log"
	"net/http"

	"github.com/brunoquindeler/go-with-tests/templates/blogrenderer"
)

func main() {

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		log.Fatal(err)
	}

	var aPost1 = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	var aPost2 = blogrenderer.Post{
		Title:       "hello world 2",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	var posts = []blogrenderer.Post{aPost1, aPost2}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		postRenderer.Render(w, aPost1)
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		postRenderer.RenderIndex(w, posts)
	})

	http.ListenAndServe(":8080", nil)
}
