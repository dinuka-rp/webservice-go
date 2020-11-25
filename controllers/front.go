package controllers

import "net/http"

// handles all the routing in the application

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
