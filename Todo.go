package main

import "time"

// Todo is a todo object
type Todo struct {
	Id 		  int 		`json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos represents a list of Todo objects
type Todos []Todo
