package infrastructure

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pdoh00/dndAdventuresLeague/interfaces"
)

// NewRouter creates a new mux router
func NewRouter(routes []interfaces.Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//TODO Figure out how to handle these in the Route object
	router.
		PathPrefix("/static/").
		Handler(WebLogger(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))), "Static Resource"))

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//decorate the handler with a logger
		handler = WebLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
