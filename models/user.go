package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // (collection of) slice of users with pointers to User objects
	nextID = 1
)
