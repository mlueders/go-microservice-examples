package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
)

func (suite *UserTestSuite) TestHttpHappyPath() {
	service := NewService(repository)
	controller := &controller{service: service}

	addUserRequest := AddUserRequest{
		FirstName: "first",
		LastName:  "last",
		Address:   Address{
			City:  "the city",
			State: "the state",
		},
	}
	addUserResponse := User{}

	suite.Run("should POST user", func() {
		requestBody, _ := json.Marshal(addUserRequest)

		request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()
		controller.handlePostUser(response, request)
		json.Unmarshal(response.Body.Bytes(), &addUserResponse)

		suite.NotEmpty(addUserResponse.Id,
			"AddUser returned user with no id")
		suite.Equal(addUserResponse.FirstName, addUserRequest.FirstName,
			"FirstName == %q, want %q", addUserResponse.FirstName, addUserRequest.FirstName)
	})

	suite.Run("should GET user", func() {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		controller.handleGetUser(response, request)
		getUserResponse := User{}
		json.Unmarshal(response.Body.Bytes(), &getUserResponse)

		suite.True(reflect.DeepEqual(addUserResponse, getUserResponse),
			"User == %+v, want %+v", getUserResponse, addUserResponse)
	})

	suite.Run("should DELETE user", func() {
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		controller.handleDeleteUser(response, request)

		request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response = httptest.NewRecorder()
		controller.handleGetUser(response, request)

		suite.Equal(response.Code, http.StatusNotFound,
			"Response == %q, want %q", response.Code, http.StatusNotFound)
	})
}
