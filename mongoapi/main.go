package main

import (
	"fmt"
	"github.com/prabhat8991/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
