package user

func (suite *UserTestSuite) TestService() {
	service := NewService(repository)

	suite.Run("should add user", func() {
		userToAdd := AddUserRequest{
			FirstName: "first",
			LastName:  "last",
			Address: Address{
				City:  "Austin",
				State: "Texas",
			},
		}

		addedUser := service.AddUser(&userToAdd)
		suite.NotNil(addedUser, "AddUser returned nil")
		suite.NotEmpty(addedUser.Id, "AddUser returned user with no id")
		suite.Equal(userToAdd.FirstName, addedUser.FirstName,
			"FirstName == %q, want %q", addedUser.FirstName, userToAdd.FirstName)
		suite.Equal(userToAdd.LastName, addedUser.LastName,
			"LastName == %q, want %q", addedUser.LastName, userToAdd.LastName)
		suite.Equal(userToAdd.Address.State, addedUser.Address.State,
			"Address.State == %q, want %q", addedUser.Address.State, userToAdd.Address.State)
		suite.Equal(userToAdd.Address.City, addedUser.Address.City,
			"Address.City mismatch == %q, want %q", addedUser.Address.City, userToAdd.Address.City)
	})

	suite.Run("should retrieve added user", func () {
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

		suite.NotNil(retrievedUser, "GetUser returned nil")
		suite.Equal(userToAdd.FirstName, retrievedUser.FirstName,
			"FirstName == %q, want %q", retrievedUser.FirstName, userToAdd.FirstName)
		suite.Equal(userToAdd.LastName, retrievedUser.LastName,
			"LastName == %q, want %q", retrievedUser.LastName, userToAdd.LastName)
		suite.Equal(userToAdd.Address.State, retrievedUser.Address.State,
			"Address.State == %q, want %q", retrievedUser.Address.State, userToAdd.Address.State)
		suite.Equal(userToAdd.Address.City, retrievedUser.Address.City,
			"Address.City == %q, want %q", retrievedUser.Address.City, userToAdd.Address.City)
	})

	suite.Run("should delete user", func() {
		userToAdd := AddUserRequest{
			FirstName: "first",
			LastName:  "last",
			Address: Address{
				City:  "Austin",
				State: "Texas",
			},
		}

		addedUser := service.AddUser(&userToAdd)
		suite.NotNil(addedUser, "AddUser returned nil")

		service.RemoveUser(addedUser.Id)

		retrievedUser := service.GetUser(addedUser.Id)
		suite.Nil(retrievedUser, "GetUser returned user after removal")
	})
}
