package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atharvamahamuni/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb project")

	fmt.Println("Making server up and running on Port:4000...")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Running server on port:4000")

}
