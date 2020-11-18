package controllers

import "net/http"

type userController struct{}

// method (basically another function)
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
