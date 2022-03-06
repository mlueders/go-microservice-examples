package main

import (
	"go-microservice-examples/baseline/user"
	"net/http"
)

func main() {
	user.SetupUserController()
	http.ListenAndServe(":8090", nil)
}

