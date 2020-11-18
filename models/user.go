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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u) // &u grabs a reference to the user object that was provided
	// (because pointers to User objects are only stored in the users slice)

	return u, nil
}
