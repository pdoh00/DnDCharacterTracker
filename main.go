package main

//Using http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/
//to see how I like the layout of the project

import (
	"crypto/tls"
	"html/template"
	"log"
	"net/http"

	"github.com/pdoh00/dndAdventuresLeague/domain"
	"github.com/pdoh00/dndAdventuresLeague/infrastructure"
	"github.com/pdoh00/dndAdventuresLeague/interfaces"
	"github.com/pdoh00/dndAdventuresLeague/useCases"
)

var (
	cachedTemplates *template.Template
)

func init() {
	cachedTemplates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// rootdir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	dbHandler := infrastructure.NewSqliteHandler("./dndAdventureLeague.sqlite")
	dbHandler.Execute("create table IF NOT EXISTS characters (id integer not null primary key," +
		"characterName text," +
		"class text," +
		"level integer," +
		"background text," +
		"playerName text," +
		"faction text," +
		"race text," +
		"alignment text," +
		"xp integer," +
		"dci text," +
		"strength int," +
		"dexterity int," +
		"constitution int," +
		"intelligence int," +
		"wisdom int," +
		"charisma int)")

	dbHandler.Execute("create table IF NOT EXISTS players (dci text not null primary key," +
		"first_name text," +
		"last_name text)")

	dbHandler.Execute("create table IF NOT EXISTS users (id integer not null primary key," +
		"email text not null," +
		"password text not null," +
		"isAdmin integer)")

	handlers := make(map[string]interfaces.DBHandler)
	handlers[interfaces.DBUserRepoID] = dbHandler
	handlers[interfaces.DBPlayerRepoID] = dbHandler
	handlers[interfaces.DBCharacterRepoID] = dbHandler

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)

	characterInteractor := new(usecases.CharacterInteractor)
	characterInteractor.CharacterRepository = interfaces.NewDBCharacterRepo(handlers)
	characterInteractor.CharacterRepository.Store(domain.Character{
		CharacterName: "LionHeart",
		Class:         "Fighter",
		Level:         1,
		Background:    "Folk Hero",
		PlayerName:    "Phil",
		Faction:       "Harpers",
		Race:          "Human",
		Alignment:     "Lawful Good",
		XP:            0,
		DCI:           "12345",
		Strength:      15,
		Dexterity:     13,
		Constitution:  14,
		Intelligence:  8,
		Wisdom:        12,
		Charisma:      10})

	characterInteractor.PlayerRepository = interfaces.NewDBWotCPlayerRepo(handlers)
	characterInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)

	webserviceHandler := interfaces.WebServiceHandler{}
	webserviceHandler.UserInteractor = userInteractor
	webserviceHandler.CharacterInteractor = characterInteractor
	webserviceHandler.Templates = cachedTemplates

	routes := []interfaces.Route{
		interfaces.Route{
			Name:    "Index",
			Pattern: "/",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplayHomePage(w, r)
			},
		},
		interfaces.Route{
			Name:    "LoginPage",
			Pattern: "/profile/login",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplayLoginPage(w, r)
			},
		},
		interfaces.Route{
			Name:    "Login",
			Pattern: "/profile/login",
			Method:  "POST",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.AuthenticateUserAndRedirect(w, r)
			},
		},
		interfaces.Route{
			Name:    "SignUpPage",
			Pattern: "/profile/signup",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplaySignUpPage(w, r)
			},
		},
		interfaces.Route{
			Name:    "SignUp",
			Pattern: "/profile/signup",
			Method:  "POST",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.CreateNewUser(w, r)
			},
		},
		interfaces.Route{
			Name:    "Player",
			Pattern: "/player/{playerID}",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplayPlayerPage(w, r)
			},
		},
		interfaces.Route{
			Name:    "Character",
			Pattern: "/character/create",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplayCreateCharacterPage(w, r)
			},
		},
		interfaces.Route{
			Name:    "Character",
			Pattern: "/character/{characterID}",
			Method:  "GET",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				webserviceHandler.DisplayCharacterPage(w, r)
			},
		},
	}

	router := infrastructure.NewRouter(routes)
	config := &tls.Config{MinVersion: tls.VersionTLS10}
	server := &http.Server{
		Addr:      ":8081",
		Handler:   router,
		TLSConfig: config,
	}

	//TODO: Change to TLS
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
