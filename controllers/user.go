package controllers

import (
	"net/http"
	"regexp"
)

type userController struct{
	userIDPattern *regexp.Regexp

}

// method (basically another function)
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// constructor
func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}