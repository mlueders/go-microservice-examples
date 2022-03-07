package user

type AddUserRequest struct {
	FirstName string
	LastName  string
	Address   Address
}

type User struct {
	Id        string
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	City  string
	State string
}

