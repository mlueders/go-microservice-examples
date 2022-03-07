package main

import (
	"database/sql"
	"go-microservice-examples/user"
	"log"
	"net/http"
	"os"
)

func main() {
	db := initDb()
	defer db.Close()

	userRepository := user.NewUserRepository(db)
	userRepository.CreateTables()

	user.SetupUserController(user.NewUserService(userRepository))
	http.ListenAndServe(":8090", nil)
}

func initDb() *sql.DB {
	dbName := "sqllite.db"
	os.Remove(dbName)
	file, err := os.Create(dbName)
	if err != nil {
	}
	file.Close()

	if err != nil {
		log.Fatal("could not create db file: %w", err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", file.Name())
	if err != nil {
		log.Fatal("could not connect to database: %w", err)
	}
	return db
}
