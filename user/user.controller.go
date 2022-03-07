package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type userController struct {
	service *UserService
}

func SetupUserController(service *UserService) {
	controller := userController{service: service}
	http.HandleFunc("/users", controller.userHandler)
	http.HandleFunc("/users/", controller.userWithIdHandler)
}

func (c *userController) handleGetUsers(response http.ResponseWriter, request *http.Request) {
	usersJson, err := json.Marshal(c.service.GetUsers())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(usersJson)
}

func (c *userController) handleGetUser(response http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/users/")
	user := c.service.GetUser(id)
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

func (c *userController) handlePostUser(response http.ResponseWriter, request *http.Request) {
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
	addedUser := c.service.AddUser(&newUser)
	addedUserJson, err := json.Marshal(addedUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Write(addedUserJson)
}

func (c *userController) handleDeleteUser(_ http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/users/")
	c.service.RemoveUser(id)
}

func (c *userController) userHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		c.handleGetUsers(response, request)
	case http.MethodPost:
		c.handlePostUser(response, request)
	}
}

func (c *userController) userWithIdHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		c.handleGetUser(response, request)
	case http.MethodDelete:
		c.handleDeleteUser(response, request)
	}
}
