package domain

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

// WotCPlayer holds data regarding a registered player
type WotCPlayer struct {
	DCI       string
	FirstName string
	LastName  string
}
