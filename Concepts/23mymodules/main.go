package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// https://go.dev/blog/using-go-modules

func main() {
	fmt.Println("Hello, Let's see modules.")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To First Golang Server</h1>"))
}
