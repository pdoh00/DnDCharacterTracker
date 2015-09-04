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

// UserInteractor defines methods available on a concrete UserInteractor
type UserInteractor interface {
	Add(email string, password string) error
	// Returns nil if authentication is successful err otherwise
	Authenticate(email string, password string) error
}

// WebServiceHandler is used to handle all http requests
type WebServiceHandler struct {
	CharacterInteractor CharacterInteractor
	UserInteractor      UserInteractor
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

// AuthenticateUserAndRedirect and redirects to appropriate pages
func (handler *WebServiceHandler) AuthenticateUserAndRedirect(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	err := handler.UserInteractor.Authenticate(email, password)
	if err != nil {
		//redirect to login page
		//TODO: add .Error to page and pass into template
		errorData := struct{ Error string }{"Invalid user name or password"}
		handler.Templates.ExecuteTemplate(w, "login.html", errorData)
	} else {
		handler.DisplayPlayerPage(w, r)
	}
}

// DisplaySignUpPage handles a request for the sign up page
func (handler *WebServiceHandler) DisplaySignUpPage(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "signup.html", nil)
}

// CreateNewUser handles a request from the sign up page to create a new user
func (handler *WebServiceHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	//TODO: Checkout using gorilla schema to bind form values to user values
	email := r.FormValue("email")
	password := r.FormValue("password")
	handler.UserInteractor.Add(email, password)
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
