package user

import (
	"github.com/google/uuid"
)

var userMap = map[string]*User{}

type AddUserRequest struct {
	FirstName string
	LastName  string
	Address   *Address
}

type User struct {
	Id        string
	FirstName string
	LastName  string
	Address   *Address
}

type Address struct {
	City  string
	State string
}

func init() {
	user := User{
		Id:        uuid.New().String(),
		FirstName: "Bo",
		LastName:  "Jangles",
		Address: &Address{
			City:  "New Orleans",
			State: "Louisiana",
		},
	}
	userMap[user.Id] = &user
}

func GetUser(id string) *User {
	return userMap[id]
}

func GetUsers() []*User {
	userList := make([]*User, 0, len(userMap))
	for _, val := range userMap {
		userList = append(userList, val)
	}
	return userList
}

func AddUser(userToAdd *AddUserRequest) *User {
	user := User{
		Id:        uuid.New().String(),
		FirstName: userToAdd.FirstName,
		LastName:  userToAdd.LastName,
		Address:   userToAdd.Address,
	}
	userMap[user.Id] = &user
	return &user
}

func RemoveUser(id string) {
	delete(userMap, id)
}
