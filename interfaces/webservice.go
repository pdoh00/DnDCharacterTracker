package interfaces

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/pdoh00/dndAdventuresLeague/domain"
)

// CharacterInteractor defines methods available on the CharacterInteractor
type CharacterInteractor interface {
	Add(email string, character domain.Character) error
	Character(characterID int) domain.Character
}

// WebServiceHandler is used to handle all http requests
type WebServiceHandler struct {
	CharacterInteractor CharacterInteractor
	Templates           *template.Template
}

// ShowCharacter handles a request to show character data
func (handler WebServiceHandler) ShowCharacter(w http.ResponseWriter, r *http.Request) {
	characterID, err := strconv.Atoi(r.FormValue("characterID"))
	if err != nil {
		panic(err)
	}
	character := handler.CharacterInteractor.Character(characterID)
	handler.Templates.ExecuteTemplate(w, "/templates/character.html", character)
}
