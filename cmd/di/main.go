package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brunoquindeler/go-with-tests/dependency_injection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependency_injection.Greet(w, "World")
}

func main() {
	dependency_injection.Greet(os.Stdout, "Elodie")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
