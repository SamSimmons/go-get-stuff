package main

import "net/http"

// Route some comment here?
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes something
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MatchInfo",
		"GET",
		"/match/{id}",
		MatchInfo,
	},
	Route{
		"MatchStats",
		"GET",
		"/match/{id}/stats",
		MatchStats,
	},
	Route{
		"ExtraInfo",
		"GET",
		"/match/{id}/info",
		ExtraInfo,
	},
}
