package main

import "net/http"

// Route defines the structure for defining routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds an array of Routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Create new bot",
		"POST",
		"/create",
		Create,
	},
	Route{
		"Deletes a bot",
		"DELETE",
		"/delete",
		Delete,
	},
	Route{
		"Edit config",
		"PATCH",
		"/edit",
		Edit,
	},
	Route{
		"Restart bot",
		"GET",
		"/restart",
		Restart,
	},
	Route{
		"Start bot",
		"GET",
		"/start",
		Start,
	},
	Route{
		"Stop bot",
		"GET",
		"/stop",
		Stop,
	},
}
