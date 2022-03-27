package user

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
}

func TestUserTestSuite(t *testing.T) {
	db, err := OpenSqlLiteDB("sqllite-user-test.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	initUserRepository(db.db)
	suite.Run(t, new(UserTestSuite))
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
