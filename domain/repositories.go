package domain

// PlayerRepository defines methods available for player data access
type PlayerRepository interface {
	Store(player WotCPlayer) error
	FindByDCI(dci string) WotCPlayer
}

// CharacterRepository defines methods available for character data access
type CharacterRepository interface {
	Store(character Character) error
	Retire(characterID int) error
	FindByPlayerID(playerID int) []Character
	FindByID(characterID int) Character
	FindByDCI(dci string) []Character
}
