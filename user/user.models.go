package user

type AddUserRequest struct {
	FirstName string
	LastName  string
	Address   Address
}

// User model info
// @Description User account information
// @Description with user id, name, and address
type User struct {
	Id           string
	FirstName    string
	LastName     string
	Address      Address
	FavoriteJoke string
} //@name User

type Address struct {
	City  string
	State string
} //@name Address
