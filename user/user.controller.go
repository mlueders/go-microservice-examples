package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func SetupUserController() {
	http.HandleFunc("/users", UserHandler)
	http.HandleFunc("/users/", UserWithIdHandler)
}

func handleGetUsers(response http.ResponseWriter, request *http.Request) {
	usersJson, err := json.Marshal(GetUsers())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(usersJson)
}

func handleGetUser(response http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/users/")
	user := GetUser(id)
	if user == nil {
		response.WriteHeader(http.StatusNotFound)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(userJson)
}

func handlePostUser(response http.ResponseWriter, request *http.Request) {
	var newUser AddUserRequest
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &newUser)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	addedUser := AddUser(&newUser)
	addedUserJson, err := json.Marshal(addedUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Write(addedUserJson)
}

func handleDeleteUser(_ http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/users/")
	RemoveUser(id)
}

func UserHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetUsers(response, request)
	case http.MethodPost:
		handlePostUser(response, request)
	}
}

func UserWithIdHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handleGetUser(response, request)
	case http.MethodDelete:
		handleDeleteUser(response, request)
	}
}
