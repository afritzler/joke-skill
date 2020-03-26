package main

import (
	"log"
	"net/http"

	"github.com/afritzler/joke-skill"
)

func main() {
	http.HandleFunc("/random_joke", joke.RandomJoke)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
