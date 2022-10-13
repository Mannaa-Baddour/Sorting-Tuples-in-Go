package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	srv "github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/server"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

func TestSendRequestToSort(t *testing.T) {
	expectedResult := srv.ResponseData{
		OperationStatus: "Operation Done Successfully",
		OperationResult: []sorting.Tuple{{3, 2, 1}, {2, 3, 4}, {6, 3, 5}, {1, 4, 6}},
		OperationError:  false,
	}
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		srv.HandleRequestToSort(responseWriter, request)
	}))
	defer server.Close()
	client := http.Client{
		Timeout: time.Second * 5,
	}
	var response srv.ResponseData
	var err error
	response, err = sendRequestToSort(client, server.URL, "", "../server/test.txt", "../server/result.txt", "2")
	if err != nil {
		t.Errorf("expected: no error, got: %v", err)
	}
	if !reflect.DeepEqual(response, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, response)
	}
}
