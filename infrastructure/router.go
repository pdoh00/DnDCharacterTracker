package infrastructure

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pdoh00/dndAdventuresLeague/interfaces"
)

// NewRouter creates a new mux router
func NewRouter(routes []interfaces.Route) *mux.Router {
	rootdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	//TODO Figure out how to handle these in the Route object
	router.
		PathPrefix("/static/img/").
		Handler(WebLogger(prefixHandler(rootdir, "/static/img/"), "/static/img/"))

	router.
		PathPrefix("/static/css/").
		Handler(WebLogger(prefixHandler(rootdir, "/static/css/"), "/static/css/"))

	router.
		PathPrefix("/static/js/").
		Handler(WebLogger(prefixHandler(rootdir, "/static/js/"), "/static/js/"))

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

func prefixHandler(rootdir string, routePrefix string) http.Handler {
	return http.
		StripPrefix(routePrefix,
		http.FileServer(http.Dir(rootdir+routePrefix)))
}
