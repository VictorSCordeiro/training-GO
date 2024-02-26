package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Extract path from the request URL
	path := r.URL.Path[1:] // Removing the leading '/'

	// Write the requested path to the response
	fmt.Fprintf(w, path)
}

func main() {
	//var name string
	fmt.Println("Hello World")
	//fmt.Scan(&name)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Attach handleRequest function to the root endpoint
	mux.HandleFunc("/", handleRequest)

	// Start the HTTP server on port 3333
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
