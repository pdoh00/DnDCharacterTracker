package interfaces

import (
	"fmt"

	"github.com/pdoh00/dndAdventuresLeague/domain"
)

const (
	//DBPlayerRepoID is the identifier for the DBPlayerRepo
	DBPlayerRepoID = "DBPlayerRepo"
)

// DBWotCPlayerRepo is used to query and persist WotC Player data
type DBWotCPlayerRepo DBRepo

// NewDBWotCPlayerRepo creates a new player repository
func NewDBWotCPlayerRepo(dbHandlers map[string]DBHandler) *DBWotCPlayerRepo {
	dbUserRepo := new(DBWotCPlayerRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers[DBPlayerRepoID]
	return dbUserRepo
}

// Store persists a usecases.User to the data store
func (repo *DBWotCPlayerRepo) Store(player domain.WotCPlayer) error {
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO wotcplayer (first_name, last_name, dci) VALUES ('%s', '%s', '%s')",
		player.FirstName, player.LastName, player.DCI))
	return nil
}

// FindByDCI finds a player by DCI
func (repo *DBWotCPlayerRepo) FindByDCI(dci string) domain.WotCPlayer {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT first_name, last_name FROM players WHERE dci = %s", dci))
	var firstName string
	var lastName string
	row.Scan(&firstName, &lastName)
	return domain.WotCPlayer{FirstName: firstName, LastName: lastName, DCI: dci}
}
