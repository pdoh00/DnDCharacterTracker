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

// DisplayHomePage handles a request for the home page
func (handler *WebServiceHandler) DisplayHomePage(w http.ResponseWriter, req *http.Request) {
	handler.Templates.ExecuteTemplate(w, "index.html", nil)
}

// DisplayLoginPage handles a request for the login page
func (handler *WebServiceHandler) DisplayLoginPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "login.html", nil)
}

// DisplaySignUpPage handles a request for the sign up page
func (handler *WebServiceHandler) DisplaySignUpPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "signup.html", nil)
}

// DisplayPlayerPage handles a request for the character creation page
func (handler *WebServiceHandler) DisplayPlayerPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "player.html", nil)
}

// DisplayCreateCharacterPage handles a request for the character creation page
func (handler *WebServiceHandler) DisplayCreateCharacterPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "createCharacter.html", nil)
}

// DisplayCharacterPage handles a request to show character data
func (handler *WebServiceHandler) DisplayCharacterPage(w http.ResponseWriter, r *http.Request) {
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
