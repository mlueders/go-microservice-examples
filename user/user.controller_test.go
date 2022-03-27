package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHappyPath(t *testing.T) {
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

	t.Run("should POST user", func(t *testing.T) {
		requestBody, _ := json.Marshal(addUserRequest)

		request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()
		controller.handlePostUser(response, request)
		json.Unmarshal(response.Body.Bytes(), &addUserResponse)

		assert.Assert(t, addUserResponse.Id == "")
		assert.Equal(t, addUserResponse.FirstName, addUserRequest.FirstName)
	})

	t.Run("should GET user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		controller.handleGetUser(response, request)
		getUserResponse := User{}
		json.Unmarshal(response.Body.Bytes(), &getUserResponse)

		assert.DeepEqual(t, addUserResponse, getUserResponse)
	})

	t.Run("should DELETE user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		controller.handleDeleteUser(response, request)

		request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response = httptest.NewRecorder()
		controller.handleGetUser(response, request)

		assert.Equal(t, response.Code, http.StatusNotFound)
	})
}
