package main

import (
	"fmt"
)

var currentID int
var todos Todos

// generate some data for the repo
func init() {

}

// RepoFindTodo returns the Todo object at the given index or an empty Todo
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return an empty Todo
	return Todo{}
}

// RepoCreateTodo creates a Todo object and adds it to the list of todos
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.Id = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyToDo removes the Todo at the given index from the list of todos
func RepoDestroyToDo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to dalete", id)
}
