package user

import (
	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetUser(id string) *User {
	return s.repository.FindUserById(id)
}

func (s *Service) GetUsers() []*User {
	return s.repository.FindAllUsers()
}

func (s *Service) AddUser(userToAdd *AddUserRequest) *User {
	user := User{
		Id:        uuid.New().String(),
		FirstName: userToAdd.FirstName,
		LastName:  userToAdd.LastName,
		Address:   userToAdd.Address,
	}
	s.repository.InsertUser(&user)
	return &user
}

func (s *Service) RemoveUser(id string) {
	s.repository.DeleteUser(id)
}
