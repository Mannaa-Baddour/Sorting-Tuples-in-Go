package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

// main is a starting point function that takes a URL and a port as cmd string argument to run the server at that address
// while handling any errors that might occur related to the URL or the port in use.
func main() {
	url := flag.String("host", "localhost", "URL to where the server starts")
	port := flag.String("port", "8080", "Port used by the server")
	flag.Parse()
	address := fmt.Sprintf("%s:%s", *url, *port)
	http.HandleFunc("/", server.HandleRequestToSort)
	err := http.ListenAndServe(address, nil)
	sorting.HandleErrors(err, nil, true)
	fmt.Println("Server started at address:", address)
}
