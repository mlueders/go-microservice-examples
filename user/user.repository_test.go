package user

import "database/sql"

var repository *UserRepository

func initUserRepository(db *sql.DB) {
	repository = NewUserRepository(db)
	repository.CreateTables()
}
