package user

import (
	"gotest.tools/assert"
	"testing"
)

func TestService(t *testing.T) {
	service := NewService(repository)

	t.Run("should add user", func(t *testing.T) {
		userToAdd := AddUserRequest{
			FirstName: "first",
			LastName:  "last",
			Address: Address{
				City:  "Austin",
				State: "Texas",
			},
		}

		addedUser := service.AddUser(&userToAdd)
		assert.Assert(t, addedUser != nil, "AddUser returned nil")
		assert.Assert(t, addedUser.Id != "", "AddUser returned user with no id")
		assert.Equal(t, userToAdd.FirstName, addedUser.FirstName)
		assert.Equal(t, userToAdd.LastName, addedUser.LastName)
		assert.Equal(t, userToAdd.Address, addedUser.Address)
	})

	t.Run("should retrieve added user", func (t *testing.T) {
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

		assert.Assert(t, retrievedUser != nil, "GetUser returned nil")
		assert.Equal(t, userToAdd.FirstName, retrievedUser.FirstName)
		assert.Equal(t, userToAdd.LastName, retrievedUser.LastName)
		assert.Equal(t, userToAdd.Address, retrievedUser.Address)
	})

	t.Run("should delete user", func(t *testing.T) {
		userToAdd := AddUserRequest{
			FirstName: "first",
			LastName:  "last",
			Address: Address{
				City:  "Austin",
				State: "Texas",
			},
		}

		addedUser := service.AddUser(&userToAdd)
		assert.Assert(t, addedUser != nil, "AddUser returned nil")

		service.RemoveUser(addedUser.Id)

		retrievedUser := service.GetUser(addedUser.Id)
		assert.Assert(t, retrievedUser, "GetUser returned user after removal")
	})
}
