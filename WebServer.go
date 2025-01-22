/*
   This program is a simple yet complete example of a web server in Go.
   It is typical to use Go for such applications due to its efficient handling
   of concurrency and its straightforward, powerful standard library for networking tasks.
*/

// Package declaration: main is the default package for creating executable Go applications.
package main

// Importing necessary packages:
// - fmt: Provides formatted I/O with functions analogous to C's printf and scanf.
// - net/http: Provides HTTP client and server implementations.
import (
	"fmt"
	"net/http"
)

// helloWorld is a handler function that writes "Hello, World!" to the HTTP response.
// It matches the signature required by http.HandlerFunc.
// - w http.ResponseWriter: An interface used to send HTTP responses.
// - r *http.Request: A data structure that represents the client HTTP request.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Fprintf is a function from the fmt package that formats and writes to an io.Writer object.
	// Here, it's writing "Hello, World!" to the HTTP response writer (w).
	fmt.Fprintf(w, "Hello, World!")
}

// main is the entry point for the Go application.
func main() {
	// HandleFunc registers the handler function for the given URL pattern.
	// "/" indicates the root URL path.
	// helloWorld is the function that will be called when a request is made to the root URL.
	http.HandleFunc("/", helloWorld)

	// ListenAndServe starts an HTTP server with a given address and handler.
	// Here, ":8080" specifies that the server should listen on port 8080.
	// nil indicates that the default ServeMux is used as the handler.
	http.ListenAndServe(":8080", nil)
}
