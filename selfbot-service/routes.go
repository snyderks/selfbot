package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
		"Edit config",
		"PATCH",
		"/edit",
		Edit,
	},
	Route{
		"Edit config",
		"PATCH",
		"/restart",
		Restart,
	},
}
