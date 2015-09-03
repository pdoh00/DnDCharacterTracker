// Package interfaces -
// The code does what it takes to make the fact that an HTTP call arrived
// unrecognizable for the use cases layer.
package interfaces

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pdoh00/dndAdventuresLeague/domain"
)

// CharacterInteractor defines methods available on the CharacterInteractor
type CharacterInteractor interface {
	Add(email string, character domain.Character) error
	RetrieveCharacter(characterID int) domain.Character
}

// WebServiceHandler is used to handle all http requests
type WebServiceHandler struct {
	CharacterInteractor CharacterInteractor
	Templates           *template.Template
}

// DisplayLoginPage handles a request for the login page
func (handler *WebServiceHandler) DisplayLoginPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "index.html", nil)
}

// DisplayCharacter handles a request to show character data
func (handler *WebServiceHandler) DisplayCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	characterID, err := strconv.Atoi(vars["characterID"])
	fmt.Printf("Display character with id %d\n", characterID)
	if err != nil {
		panic(err)
	}
	character := handler.CharacterInteractor.RetrieveCharacter(characterID)
	fmt.Printf("Character retrieved:\n%s\n", character.ToString())
	handler.Templates.ExecuteTemplate(w, "character.html", character)
}
