package interfaces

import "net/http"

// Route is the definition of an http route
type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

// Routes is a container for all defined routes on the webservice
type Routes []Route
