package main

import (
	"go-microservice-examples/user"
	"net/http"
)

func main() {
	user.SetupUserController(user.NewUserService())
	http.ListenAndServe(":8090", nil)
}

