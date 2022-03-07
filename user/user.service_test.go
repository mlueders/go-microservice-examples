package user

import (
	"testing"
)

var service = NewUserService()

func TestShouldAddUser(t *testing.T) {
	userToAdd := AddUserRequest{
		FirstName: "first",
		LastName:  "last",
		Address: Address{
			City:  "Austin",
			State: "Texas",
		},
	}

	addedUser := service.AddUser(&userToAdd)
	if addedUser == nil {
		t.Errorf("AddUser returned nil")
	}
	if addedUser.Id == "" {
		t.Error("AddUser returned user with no id")
	}
	if addedUser.FirstName != userToAdd.FirstName {
		t.Errorf("FirstName == %q, want %q", addedUser.FirstName, userToAdd.FirstName)
	}
	if addedUser.LastName != userToAdd.LastName {
		t.Errorf("LastName == %q, want %q", addedUser.LastName, userToAdd.LastName)
	}
	if addedUser.Address.State != userToAdd.Address.State {
		t.Errorf("Address.State == %q, want %q", addedUser.Address.State, userToAdd.Address.State)
	}
	if addedUser.Address.City != userToAdd.Address.City {
		t.Errorf("Address.City mismatch == %q, want %q", addedUser.Address.City, userToAdd.Address.City)
	}
}

func TestShouldRetrieveAddedUser(t *testing.T) {
	userToAdd := AddUserRequest{
		FirstName: "first",
		LastName:  "last",
		Address: Address{
			City:  "Austin",
			State: "Texas",
		},
	}

	addedUser := service.AddUser(&userToAdd)
	retrievedUser := service.GetUser(addedUser.Id)

	if retrievedUser == nil {
		t.Errorf("GetUser returned nil")
	}
	if retrievedUser.FirstName != userToAdd.FirstName {
		t.Errorf("FirstName == %q, want %q", retrievedUser.FirstName, userToAdd.FirstName)
	}
	if retrievedUser.LastName != userToAdd.LastName {
		t.Errorf("LastName == %q, want %q", retrievedUser.LastName, userToAdd.LastName)
	}
	if retrievedUser.Address.State != userToAdd.Address.State {
		t.Errorf("Address.State == %q, want %q", retrievedUser.Address.State, userToAdd.Address.State)
	}
	if retrievedUser.Address.City != userToAdd.Address.City {
		t.Errorf("Address.City == %q, want %q", retrievedUser.Address.City, userToAdd.Address.City)
	}	
}

func TestShouldDeleteUser(t *testing.T) {
	userToAdd := AddUserRequest{
		FirstName: "first",
		LastName:  "last",
		Address: Address{
			City:  "Austin",
			State: "Texas",
		},
	}

	addedUser := service.AddUser(&userToAdd)
	if addedUser == nil {
		t.Errorf("AddUser returned nil")
	}

	service.RemoveUser(addedUser.Id)

	retrievedUser := service.GetUser(addedUser.Id)
	if retrievedUser != nil {
		t.Errorf("GetUser returned user after removal")
	}
}
