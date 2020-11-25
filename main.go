package main

import (
	"github.com/dinuka-rp/webservice/controllers"
	"net/http"
)

//	this can be run using:
//	go run main.go
//	or
//	go run github.com/dinuka-rp/webservice
//	or
//	go run .
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // handler passed in is a front-controller: handles all the requests received to the server.
	// nil - use default front-controller


	//u := models.User{
	//	ID:        2,
	//	FirstName: "Tricia",
	//	LastName:  "McMillan",
	//}
	//fmt.Println(u)
}
