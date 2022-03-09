package user

import "database/sql"

var repository *Repository

func initUserRepository(db *sql.DB) {
	repository = NewRepository(db)
	repository.CreateTables()
}
