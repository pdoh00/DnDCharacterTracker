package domain

import (
	"fmt"
)

//Character holds data for the character
type Character struct {
	ID            int
	CharacterName string
	Class         string
	Level         int
	Background    string
	PlayerName    string
	Faction       string
	Race          string
	Alignment     string
	XP            int
	DCI           string
	Strength      int
	Dexterity     int
	Constitution  int
	Intelligence  int
	Wisdom        int
	Charisma      int
}

// ToString returns a string with Character data
func (c *Character) ToString() string {
	return fmt.Sprintf("ID: %d\n"+
		"CharacterName: %s\n"+
		"Class: %s\n"+
		"Level: %d\n"+
		"BackGround: %s\n"+
		"PlayerName: %s\n"+
		"Faction: %s\n"+
		"Race: %s\n"+
		"Alignment: %s\n"+
		"DCI: %s\n"+
		"Strength: %d\n"+
		"Dexterity: %d\n"+
		"Constitution: %d\n"+
		"Intelligence: %d\n"+
		"Wisdom: %d\n"+
		"Charisma: %d\n",
		c.ID, c.CharacterName, c.Class, c.Level, c.Background,
		c.PlayerName, c.Faction, c.Race, c.Alignment, c.DCI,
		c.Strength, c.Dexterity, c.Constitution, c.Intelligence,
		c.Wisdom, c.Charisma)
}

// WotCPlayer holds data regarding a registered player
type WotCPlayer struct {
	DCI       string
	FirstName string
	LastName  string
}
