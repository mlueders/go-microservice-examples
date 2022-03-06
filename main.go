package main

import (
	"go-microservice-examples/user"
	"net/http"
)

func main() {
	user.SetupUserController()
	http.ListenAndServe(":8090", nil)
}

