package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	// srv "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

func TestSendRequestToSort(t *testing.T) {
	type Res struct {
		OperationStatus string          `json:"operation-status"`
		OperationResult []sorting.Tuple `json:"operation-result"`
		OperationError  bool            `json:"operation-error"`
	}
	expectedResult := Res{
		OperationError:  false,
		OperationResult: []sorting.Tuple{{3, 2, 1}, {2, 3, 4}, {6, 3, 5}, {1, 4, 6}},
		OperationStatus: "Operation Done Successfully",
	}
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		// srv.HandleRequestToSort(...)
		// Here I was trying to use handleRequestToSort function that exist in server.go
		// instead, I had to mimic its functionality.
		responseWriter.Header().Set("Content-Type", "application/json")
		data := Res{
			OperationStatus: "Operation Done Successfully",
			OperationResult: []sorting.Tuple{{3, 2, 1}, {2, 3, 4}, {6, 3, 5}, {1, 4, 6}},
			OperationError:  false,
		}
		json.NewEncoder(responseWriter).Encode(data)
	}))
	defer server.Close()
	client := http.Client{
		Timeout: time.Second * 5,
	}
	var response map[string]interface{}
	var err error
	response, err = sendRequestToSort(client, server.URL, "", "../server/test.txt", "../server/result.txt", "2")
	if err != nil {
		t.Errorf("expected: no error, got: %v", err)
	}
	// Please Check Code Here
	fmt.Println(reflect.TypeOf(response) == reflect.TypeOf(expectedResult))
	fmt.Println(reflect.TypeOf(response["operation-error"]) == reflect.TypeOf(expectedResult.OperationError))
	fmt.Println(reflect.TypeOf(response["operation-status"]) == reflect.TypeOf(expectedResult.OperationStatus))
	fmt.Println(reflect.TypeOf(response["operation-result"]) == reflect.TypeOf(expectedResult.OperationResult))
	fmt.Println(reflect.TypeOf(response["operation-result"]))

	responseBytes, _ := json.Marshal(response)
	var typedResponse Res
	json.Unmarshal(responseBytes, &typedResponse)
	// Here, the condition fails because json file is returning a slice of interface instead of tuples.
	if !reflect.DeepEqual(typedResponse, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, typedResponse)
	}
}
