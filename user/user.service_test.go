package user

import (
	"github.com/stretchr/testify/assert"
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
		assert.NotNil(t, addedUser, "AddUser returned nil")
		assert.NotEmpty(t, addedUser.Id)
		assert.Equal(t, userToAdd.FirstName, addedUser.FirstName,
			"FirstName == %q, want %q", addedUser.FirstName, userToAdd.FirstName)
		assert.Equal(t, userToAdd.LastName, addedUser.LastName,
			"LastName == %q, want %q", addedUser.LastName, userToAdd.LastName)
		assert.Equal(t, userToAdd.Address.State, addedUser.Address.State,
			"Address.State == %q, want %q", addedUser.Address.State, userToAdd.Address.State)
		assert.Equal(t, userToAdd.Address.City, addedUser.Address.City,
			"Address.City mismatch == %q, want %q", addedUser.Address.City, userToAdd.Address.City)
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

		assert.NotNil(t, retrievedUser, "GetUser returned nil")
		assert.Equal(t, userToAdd.FirstName, retrievedUser.FirstName,
			"FirstName == %q, want %q", retrievedUser.FirstName, userToAdd.FirstName)
		assert.Equal(t, userToAdd.LastName, retrievedUser.LastName,
			"LastName == %q, want %q", retrievedUser.LastName, userToAdd.LastName)
		assert.Equal(t, userToAdd.Address.State, retrievedUser.Address.State,
			"Address.State == %q, want %q", retrievedUser.Address.State, userToAdd.Address.State)
		assert.Equal(t, userToAdd.Address.City, retrievedUser.Address.City,
			"Address.City == %q, want %q", retrievedUser.Address.City, userToAdd.Address.City)
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
		assert.NotNil(t, addedUser, "AddUser returned nil")

		service.RemoveUser(addedUser.Id)

		retrievedUser := service.GetUser(addedUser.Id)
		assert.Nil(t, retrievedUser, "GetUser returned user after removal")
	})
}
