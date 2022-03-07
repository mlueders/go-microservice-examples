package user

import (
	"github.com/google/uuid"
)

type UserService struct {
	userMap map[string]*User
}

func NewUserService() *UserService {
	return &UserService{
		userMap: map[string]*User{},
	}
}

func (s *UserService) GetUser(id string) *User {
	return s.userMap[id]
}

func (s *UserService) GetUsers() []*User {
	userList := make([]*User, 0, len(s.userMap))
	for _, val := range s.userMap {
		userList = append(userList, val)
	}
	return userList
}

func (s *UserService) AddUser(userToAdd *AddUserRequest) *User {
	user := User{
		Id:        uuid.New().String(),
		FirstName: userToAdd.FirstName,
		LastName:  userToAdd.LastName,
		Address:   userToAdd.Address,
	}
	s.userMap[user.Id] = &user
	return &user
}

func (s *UserService) RemoveUser(id string) {
	delete(s.userMap, id)
}
