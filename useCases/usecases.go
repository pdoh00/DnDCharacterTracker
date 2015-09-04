package usecases

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdoh00/dndAdventuresLeague/domain"
)

const (
	bcryptCost int = 10 //cost is 2^n
)

// UserRepository defines methods available for user data access
type UserRepository interface {
	Store(user User) error
	FindByEmail(email string) User
}

//User holds data for the user. A user is a use case concern
//Basically a user is need for authentication purposes
type User struct {
	ID       int
	Email    string
	Password string
	IsAdmin  bool
	Player   domain.WotCPlayer //might just need playerID here
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

// RetrieveCharacter returns the requested character data
func (c *CharacterInteractor) RetrieveCharacter(characterID int) domain.Character {
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

// Add creates a new user
func (u *UserInteractor) Add(email string, password string) error {
	//TODO: add password protection
	pwBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwBytes, bcryptCost)

	if err != nil {
		panic(err)
	}

	user := User{Email: email, Password: string(hashedPassword[:]), IsAdmin: false}
	u.UserRepository.Store(user)
	return nil
}

// Authenticate compares the stored user password with the passed in password
func (u *UserInteractor) Authenticate(email string, password string) error {
	user := u.UserRepository.FindByEmail(email)
	fmt.Printf("UserID %d\n Email %s\n PW %s\n Admin %v\n", user.ID, user.Email, user.Password, user.IsAdmin)
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
