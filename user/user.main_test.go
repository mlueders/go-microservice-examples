package user

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	db, err := OpenSqlLiteDB("sqllite-user-test.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	initUserRepository(db.db)

	os.Exit(m.Run())
}

type SqlLiteDB struct {
	db *sql.DB
	file *os.File
}

func OpenSqlLiteDB(name string) (*SqlLiteDB, error) {
	os.Remove(name)
	file, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("could not create db file: %w", err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", file.Name())
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return &SqlLiteDB{
		db: db,
		file: file,
	}, nil
}

func (testDB *SqlLiteDB) Close() {
	defer testDB.db.Close()
	defer os.Remove(testDB.file.Name())
}
