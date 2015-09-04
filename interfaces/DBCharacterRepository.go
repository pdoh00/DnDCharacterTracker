package interfaces

import (
	"fmt"

	"github.com/pdoh00/dndAdventuresLeague/domain"
)

// DBCharacterRepo is used to query and persist Character data
type DBCharacterRepo DBRepo

const (
	//DBCharacterRepoID is the identifier for the DBCharacterRepo
	DBCharacterRepoID = "DBCharacterRepo"
)

// NewDBCharacterRepo creates a new character repository
func NewDBCharacterRepo(dbHandlers map[string]DBHandler) *DBCharacterRepo {
	dbCharacterRepo := new(DBCharacterRepo)
	dbCharacterRepo.dbHandlers = dbHandlers
	dbCharacterRepo.dbHandler = dbHandlers[DBCharacterRepoID]
	return dbCharacterRepo
}

// Store persists a usecases.User to the data store
func (repo *DBCharacterRepo) Store(character domain.Character) error {

	sqlStmt :=
		fmt.Sprintf("INSERT INTO characters "+
			"(characterName, class, level, background, playerName, faction, race, alignment, xp, dci, strength, dexterity, constitution, intelligence, wisdom, charisma)"+
			"VALUES ('%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', %d, '%s', %d, %d, %d, %d, %d, %d)",
			character.CharacterName,
			character.Class,
			character.Level,
			character.Background,
			character.PlayerName,
			character.Faction,
			character.Race,
			character.Alignment,
			character.XP,
			character.DCI,
			character.Strength,
			character.Dexterity,
			character.Constitution,
			character.Intelligence,
			character.Wisdom,
			character.Charisma)

	_, err := repo.dbHandler.Execute(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

// FindByDCI finds a character by DCI
func (repo *DBCharacterRepo) FindByDCI(dci string) []domain.Character {
	sqlStmt := fmt.Sprintf("SELECT * FROM characters WHERE dci = %s", dci)
	row := repo.dbHandler.Query(sqlStmt)
	defer row.Close()
	var characters []domain.Character
	for row.Next() {
		characters = append(characters, parseCharacter(row))
	}
	return characters
}

// FindByID finds a character by ID
func (repo *DBCharacterRepo) FindByID(id int) domain.Character {
	sqlStmt := fmt.Sprintf("SELECT id,characterName,class,level,background,playerName,faction,race,alignment,xp,dci,strength,dexterity,constitution,intelligence,wisdom,charisma FROM characters WHERE id = %d", id)
	row := repo.dbHandler.Query(sqlStmt)
	defer row.Close()
	row.Next()
	return parseCharacter(row)
}

// Retire retires character
func (repo *DBCharacterRepo) Retire(characterID int) error {
	return nil
}

func parseCharacter(row Row) domain.Character {
	var id int
	var characterName string
	var class string
	var level int
	var background string
	var playerName string
	var faction string
	var race string
	var alignment string
	var xp int
	var dci string
	var strength int
	var dexterity int
	var constitution int
	var intelligence int
	var wisdom int
	var charisma int

	row.Scan(&id,
		&characterName,
		&class,
		&level,
		&background,
		&playerName,
		&faction,
		&race,
		&alignment,
		&xp,
		&dci,
		&strength,
		&dexterity,
		&constitution,
		&intelligence,
		&wisdom,
		&charisma)

	character := domain.Character{ID: id,
		CharacterName: characterName,
		Class:         class,
		Level:         level,
		Background:    background,
		PlayerName:    playerName,
		Faction:       faction,
		Race:          race,
		Alignment:     alignment,
		XP:            xp,
		DCI:           dci,
		Strength:      strength,
		Dexterity:     dexterity,
		Constitution:  constitution,
		Intelligence:  intelligence,
		Wisdom:        wisdom,
		Charisma:      charisma}

	return character
}
