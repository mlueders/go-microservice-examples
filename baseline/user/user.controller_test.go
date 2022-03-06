package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHttpHappyPath(t *testing.T) {
	addUserRequest := AddUserRequest{
		FirstName: "first",
		LastName:  "last",
		Address:   &Address{
			City:  "the city",
			State: "the state",
		},
	}
	addUserResponse := User{}

	t.Run("should POST user", func(t *testing.T) {
		requestBody, _ := json.Marshal(addUserRequest)

		request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()
		handlePostUser(response, request)
		json.Unmarshal(response.Body.Bytes(), &addUserResponse)

		if addUserResponse.Id == "" {
			t.Error("AddUser returned user with no id")
		}
		if addUserResponse.FirstName != addUserRequest.FirstName {
			t.Errorf("FirstName == %q, want %q", addUserResponse.FirstName, addUserRequest.FirstName)
		}
	})

	t.Run("should GET user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		handleGetUser(response, request)
		getUserResponse := User{}
		json.Unmarshal(response.Body.Bytes(), &getUserResponse)

		if reflect.DeepEqual(addUserResponse, getUserResponse) == false {
			t.Errorf("User == %+v, want %+v", getUserResponse, addUserResponse)
		}
	})

	t.Run("should DELETE user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response := httptest.NewRecorder()
		handleDeleteUser(response, request)

		request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", addUserResponse.Id), nil)
		response = httptest.NewRecorder()
		handleGetUser(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("Response == %q, want %q", response.Code, http.StatusNotFound)
		}
	})
}
