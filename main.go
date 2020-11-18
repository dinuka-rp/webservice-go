package main

import (
	"fmt"
	"github.com/dinuka-rp/webservice/models"
)

//	this can be run using:
//	go run main.go
//	or
//	go run github.com/dinuka-rp/webservice
//	or
//	go run .
func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
