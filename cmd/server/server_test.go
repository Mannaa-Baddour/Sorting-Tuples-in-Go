package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

type HandleRequestToSortTestcase struct {
	name           string
	expectedResult ResponseData
	testRequest    *http.Request
}

var handleRequestToSortTestcases = []HandleRequestToSortTestcase{
	{
		name: "Testcase1: proper request and query values",
		expectedResult: ResponseData{
			OperationStatus: "Operation Done Successfully",
			OperationResult: []sorting.Tuple{{3, 2, 1}, {2, 3, 4}, {6, 3, 5}, {1, 4, 6}},
			OperationError:  false,
		},
		testRequest: httptest.NewRequest(http.MethodGet,
			"/?input-file=test.txt&output-file=testResult.txt&column=2", nil),
	},
	{
		name: "Testcase2: non-integer column value",
		expectedResult: ResponseData{
			OperationStatus: "Operation Failed due to Failing in Converting Ascii (column value) to Integer",
			OperationResult: nil,
			OperationError:  true,
		},
		testRequest: httptest.NewRequest(http.MethodGet,
			"/?input-file=test.txt&output-file=testResult.txt&column=a", nil),
	},
}

func TestHandleRequestToSort(t *testing.T) {
	for _, testcase := range handleRequestToSortTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			responseWriter := httptest.NewRecorder()
			handleRequestToSort(responseWriter, testcase.testRequest)
			response := responseWriter.Result()
			var data ResponseData
			err := json.NewDecoder(response.Body).Decode(&data)
			fmt.Println(data)
			if err != nil {
				response.Body.Close()
				t.Errorf("expected: no error, got: %v", err)
			}
			if !reflect.DeepEqual(data, testcase.expectedResult) {
				t.Errorf("expected: %v, got: %v", testcase.expectedResult, data)
			}
			response.Body.Close()
		})
	}
}
