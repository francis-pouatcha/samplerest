package main

import "net/http"

// Route is the structure for a route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines a list of routes.
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
	Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
}
