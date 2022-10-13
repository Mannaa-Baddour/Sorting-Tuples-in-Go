package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

// ResponseData is a defined struct that will hold the server response data that will be sent
// back to the client as a json file.
type ResponseData struct {
	OperationStatus string          `json:"operation-status"`
	OperationResult []sorting.Tuple `json:"operation-result"`
	OperationError  bool            `json:"operation-error"`
}

// HandleRequestToSort is a function that receives an http get request from the client,
// extracts three parameters from the url query that specify the input file, the output file,
// and the column to sort by. It replies with a message indicating a successful operation.
func HandleRequestToSort(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("incoming connection from:", request.RemoteAddr)
	values := request.URL.Query()
	params := map[string]interface{}{}
	var err error
	var operationStatus string
	var operationResult []sorting.Tuple = nil
	operationError := false
	for key, value := range values {
		if strings.Compare(key, "column") == 0 {
			params[key], err = strconv.Atoi(value[0])
			if err != nil {
				log.Println(err)
				operationStatus = "Operation Failed due to Failing in Converting Ascii (column value) to Integer"
				operationError = true
				break
			}
		} else {
			params[key] = value[0]
		}
	}

	if !operationError {
		list := sorting.ReadingTuplesFromFile(params["input-file"].(string))
		list = sorting.SortList(list, params["column"].(int))
		sorting.SaveResultToFile(params["output-file"].(string), list)
		if list != nil {
			operationStatus = "Operation Done Successfully"
			operationResult = list
		} else {
			operationStatus = "Operation Could not be Completed, due to Errors in Sorting Part"
			operationError = true
		}
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseData := ResponseData{
		OperationStatus: operationStatus,
		OperationResult: operationResult,
		OperationError:  operationError,
	}
	err = json.NewEncoder(responseWriter).Encode(responseData)
	sorting.HandleErrors(err, nil, true)
}
