package usecases

import (
	"github.com/pdoh00/dndAdventuresLeague/domain"
)

// UserRepository defines methods available for user data access
type UserRepository interface {
	Store(user User) error
	FindByEmail(email string) User
}

//User holds data for the user. A user is a use case concern
//Basically a user is need for authentication purposes
type User struct {
	ID      int
	Email   string
	IsAdmin bool
	Player  domain.WotCPlayer //might just need playerID here
}

// CharacterInteractor is the type used for character based use cases
type CharacterInteractor struct {
	UserRepository      UserRepository
	PlayerRepository    domain.PlayerRepository
	CharacterRepository domain.CharacterRepository
}

// Add creates a new character for a given player
func (c *CharacterInteractor) Add(email string, character domain.Character) error {
	user := c.UserRepository.FindByEmail(email)

	character.DCI = user.Player.DCI
	character.PlayerName = user.Player.FirstName + " " + user.Player.LastName

	c.CharacterRepository.Store(character)

	return nil
}

// Character returns the requested character data
func (c *CharacterInteractor) Character(characterID int) domain.Character {
	return c.CharacterRepository.FindByID(characterID)
}

// PlayerInteractor is the type used for player based use cases
type PlayerInteractor struct {
	UserRepository   UserRepository
	PlayerRepository domain.PlayerRepository
}

// UserInteractor is the type used for user based use cases
type UserInteractor struct {
	UserRepository UserRepository
}
