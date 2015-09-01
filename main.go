package main

//Using http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
//to see how I like the layout of the project

import (
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/pdoh00/dndAdventuresLeague/interfaces"

	"github.com/pdoh00/dndAdventuresLeague/useCases"
)

var (
	cachedTemplates *template.Template
)

func init() {
	var err error
	cachedTemplates, err = template.ParseGlob("/templates/")
	if err != nil {
		panic(err)
	}
}

func main() {
	rootdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// This is the only way I have found to be able to serve images requested in the templates
	http.Handle("/static/css/", http.StripPrefix("/static/css/",
		http.FileServer(http.Dir(path.Join(rootdir, "/static/css/")))))

	http.Handle("/static/img/", http.StripPrefix("/static/img/",
		http.FileServer(http.Dir(path.Join(rootdir, "/static/img/")))))

	http.Handle("/static/js/", http.StripPrefix("/static/js/",
		http.FileServer(http.Dir(path.Join(rootdir, "/static/js/")))))

	characterInteractor := new(usecases.CharacterInteractor)
	// characterInteractor.CharacterRepository =
	// characterInteractor.PlayerRepository =
	// characterInteractor.UserRepository =

	webserviceHandler := interfaces.WebServiceHandler{}
	webserviceHandler.CharacterInteractor = characterInteractor

	http.HandleFunc("/character/{characterID}", func(w http.ResponseWriter, r *http.Request) {
		webserviceHandler.ShowCharacter(w, r)
	})
}
