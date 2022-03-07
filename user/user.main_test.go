package user

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// os.Exit skips defer calls, so we need to call another function
	code, err := initDbAndRun(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func initDbAndRun(m *testing.M) (code int, err error) {
	dbName := "sqllite-user-test.db"
	os.Remove(dbName)
	file, err := os.Create(dbName)
	if err != nil {
		return -1, fmt.Errorf("could not create db file: %w", err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", file.Name())
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}

	defer db.Close()
	defer os.Remove(dbName)

	initUserRepository(db)

	return m.Run(), nil
}
