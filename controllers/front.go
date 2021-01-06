package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// handles all the routing in the application

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}


func encodeResponseAsJSON(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)			// create json encoder
	enc.Encode(data)
}