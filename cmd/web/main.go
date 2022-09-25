package main

import (
	"fmt"
	"github.com/nambroa/go-mock-project/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// Code to handle incoming requests.
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello, world!")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("Bytes written:", n)
	//})

	fmt.Println("Starting application on port", portNumber)
	// Start a webserver and listen to a specific port.
	_ = http.ListenAndServe(portNumber, nil)
}
