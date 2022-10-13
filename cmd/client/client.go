package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	// "reflect"
	"time"

	srv "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

// sendRequestToSort is a function that takes a client argument, along with server and port to connect to a remote
// server, and parameters required to perform sorting operations on that server.
// Returns a struct from the decoding of the json response, and error when things go wrong.
func sendRequestToSort(client http.Client, server, port, infile, outfile, sortColumn string) (srv.ResponseData, error) {
	// Setting up query parameters and server address.
	inputFile := url.QueryEscape(infile)
	outputFile := url.QueryEscape(outfile)
	column := url.QueryEscape(sortColumn)
	if port != "" {
		server = fmt.Sprintf("%s:%s", server, port)
	}

	// Preparing the request.
	params := fmt.Sprintf("input-file=%s&output-file=%s&column=%s", inputFile, outputFile, column)
	request := fmt.Sprintf("%s/get?%s", server, params)

	// Sending request to server.
	fmt.Println("Sending request to:", server)
	response, err := client.Get(request)
	if err != nil {
		log.Println(err)
		response.Body.Close()
		//os.Exit(1)
		return srv.ResponseData{}, err
	}
	defer response.Body.Close()

	// Getting the response and decoding its json form into a server.ResponseData struct.
	var data srv.ResponseData
	fmt.Println("Decoding")
	err = json.NewDecoder(response.Body).Decode(&data)
	sorting.HandleErrors(err, nil, false)
	fmt.Println(data)
	return data, nil
}

// main is a function that represents a client forming a request by specifying the cmd parameters
// infile, outfile, and sort-column, which represent the query values that'll be added to the request
// alongside the specified server and port, then receiving a proper response from the said server.
func main() {
	// Setting up a new client with a timeout to terminate connection of 5 seconds.
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// Setting up cmd arguments using flag package.
	server := flag.String("server", "http://127.0.0.1", "Server url to connect to for sorting service")
	port := flag.String("port", "8080", "Port number that is used by the server")
	infile := flag.String("infile", "", "Input file with unsorted tuples")
	outfile := flag.String("outfile", "", "Output file to save the sorted tuples")
	sortColumn := flag.String("sort-column", "0", "Column number to sort tuples by")
	flag.Parse()

	data, err := sendRequestToSort(client, *server, *port, *infile, *outfile, *sortColumn)
	sorting.HandleErrors(err, nil, false)
	fmt.Println("Operation Completed\n", data)
}
