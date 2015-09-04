package interfaces

import (
	"fmt"

	"github.com/pdoh00/dndAdventuresLeague/useCases"
)

const (
	//DBUserRepoID is the identifier for the DBUserRepo
	DBUserRepoID = "DBUserRepo"
)

// DBUserRepo is used to query and persist User data
type DBUserRepo DBRepo

// NewDBUserRepo creates a new user repository
func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
	dbUserRepo := new(DBUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers[DBUserRepoID]
	return dbUserRepo
}

// Store persists a usecases.User to the data store
func (repo *DBUserRepo) Store(user usecases.User) error {
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO users (id, email, password, isAdmin) VALUES (%d, '%s','%s', %d)",
		user.ID, user.Email, user.Password, btoi(user.IsAdmin)))
	//playerRepo
	//playerRepo
	return nil
}

// FindByEmail a user by email
func (repo *DBUserRepo) FindByEmail(email string) usecases.User {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT id, password, isAdmin FROM users WHERE email = '%s'", email))
	var isAdmin bool
	var id int
	var password string
	row.Next()
	row.Scan(&id, &password, &isAdmin)
	fmt.Printf("pw in repo is %s\n", password)
	return usecases.User{ID: id, Email: email, Password: password, IsAdmin: isAdmin}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
