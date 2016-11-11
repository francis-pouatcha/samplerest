package main

import (
	"log"
	"net/http"
)

func main() {
	RepoCreateTodo(Todo{Name: "Write a presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreateTodo(Todo{Name: "Francis at home"})
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
