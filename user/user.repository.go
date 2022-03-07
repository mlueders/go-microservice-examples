package user

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateTables() {
	createUserTableSQL := `CREATE TABLE user (
		"id" uuid NOT NULL PRIMARY KEY,
		"first_name" TEXT,
		"last_name" TEXT,
		"address_city" TEXT,
		"address_state" TEXT
	  );`

	statement, err := r.db.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (r *UserRepository) InsertUser(user *User) {
	statement, err := r.db.Prepare(`INSERT INTO user(id, first_name, last_name, address_city, address_state) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(user.Id, user.FirstName, user.LastName, user.Address.City, user.Address.State)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (r *UserRepository) DeleteUser(id string) {
	statement, err := r.db.Prepare(`DELETE FROM user WHERE id = ?`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (r *UserRepository) FindUserById(id string) *User {
	statement, err := r.db.Prepare("SELECT id, first_name, last_name, address_city, address_state FROM user where id = ?")
	if err != nil {
		log.Fatalln(err.Error())
	}
	result, err := statement.Query(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer result.Close()

	if result.Next() == false {
		return nil
	}
	user := User{}
	result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Address.City, &user.Address.State)
	return &user
}

func (r *UserRepository) FindAllUsers() []*User {
	row, err := r.db.Query("SELECT id, first_name, last_name, address_city, address_state FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var users []*User
	for row.Next() {
		user := User{}
		users = append(users, &user)
		row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Address.City, &user.Address.State)
	}
	return users
}
