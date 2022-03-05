package main

import (
	"encoding/json"
	"fmt"
	"go-microservice-examples/baseline/user"
)

func main() {
	newUser := user.AddUserRequest{
		FirstName: "Bob",
		LastName:  "Builder",
	}
	user.AddUser(&newUser)

	for _, u := range user.GetUsers() {
		json, _ := json.Marshal(u)
		fmt.Println("User -> ", string(json))
	}
}
