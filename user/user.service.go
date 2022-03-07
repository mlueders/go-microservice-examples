package user

import (
	"github.com/google/uuid"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id string) *User {
	return s.repository.FindUserById(id)
}

func (s *UserService) GetUsers() []*User {
	return s.repository.FindAllUsers()
}

func (s *UserService) AddUser(userToAdd *AddUserRequest) *User {
	user := User{
		Id:        uuid.New().String(),
		FirstName: userToAdd.FirstName,
		LastName:  userToAdd.LastName,
		Address:   userToAdd.Address,
	}
	s.repository.InsertUser(&user)
	return &user
}

func (s *UserService) RemoveUser(id string) {
	s.repository.DeleteUser(id)
}
