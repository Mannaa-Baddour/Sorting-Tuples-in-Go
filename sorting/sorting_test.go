package sorting

import (
	"errors"
	"reflect"
	"testing"
)

// appednOneTestcaseBlueprint is a blueprint for creating testcases for invoking the method appendOne.
type appendOneTestcaseBlueprint struct {
	name string
	// tuple         Tuple
	expectedError error
}

// appendOneTestcases is a slice of objects of appendOneTestcaseBluprint, provides multiple testcase scenarios
// for invoking appendOne method.
var appendOneTestcases = []appendOneTestcaseBlueprint{
	{
		name: "testcase1: appending an integer to a tuple with empty space",
		// tuple:         Tuple([]int{1, 2}),
		expectedError: nil,
	},
	{
		name: "testcase2: appending an integer to a full tuple (with three elements)",
		// tuple:         Tuple([]int{1, 2, 3}),
		expectedError: errors.New("warning: cannot add anymore numbers, tuple is already at its maximum capacity"),
	},
}

// TestAppendOne is a testing function for when appendOne method is invoked.
// Checks if the appending operation is done correctly, taking into consideration that the tuple
// should only have three integers, otherwise it raises an error.
func TestAppendOne(t *testing.T) {
	var tuple Tuple = Tuple([]int{1, 2})
	// var tuple *Tuple = &Tuple{}
	// *tuple = []int{1, 2}
	for _, testcase := range appendOneTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			testError := tuple.appendOne(3)
			// testError := testcase.tuple.appendOne(3)
			// testError := tuple.appendOne(3)
			if testError.Error() != testcase.expectedError.Error() {
				t.Errorf("expected error: %v, got error: %v", testcase.expectedError.Error(), testError.Error())
			}
		})
	}
}

// type handleErrorsTestcaseBlueprint struct {
// 	name string
// 	expectedResult
// }

// stringToIntTestcaseBlueprint is a blueprint for creating testcases
// for invoking the function stringToInt.
type stringToIntTestcaseBlueprint struct {
	name           string
	expectedResult Tuple
	inputLine      string
}

// stringToIntTestcases is a slice of objects of stringToIntTestcaseBlueprint,
// provides multiple testcase scenarios for invoking stringToInt function.
var stringToIntTestcases = []stringToIntTestcaseBlueprint{
	{
		name:           "testcase1: correct line format and values type",
		expectedResult: Tuple([]int{1, 2, 3}),
		inputLine:      "1,2, 3",
	},
	{
		name:           "testcase2: incorrect line format with correct values type",
		expectedResult: nil,
		inputLine:      "1; 2, 3",
	},
	{
		name:           "testcase3: correct line format with incorrect values type",
		expectedResult: nil,
		inputLine:      "1 ,b, 3.0",
	},
}

// TestStringToInt is a testing function to check whether or not the invokation of stringToInt function
// checks for the correct input formatting before it get assigned to a tuple.
func TestStringToInt(t *testing.T) {
	for _, testcase := range stringToIntTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			testResult := stringToInt(testcase.inputLine)
			if !reflect.DeepEqual(testResult, testcase.expectedResult) {
				t.Errorf("expected: %v, got: %v", testcase.expectedResult, testResult)
			}
		})
	}
}

// readingInputFromFileTestcaseBlueprint is a blueprint for creating testcases
// for invoking the function ReadingTuplesFromFile.
type readingTuplesFromFileTestcaseBlueprint struct {
	name           string
	inputFilePath  string
	expectedResult []Tuple
}

// readingInputFromFileTestcases is a slice of objects of readingInputFromFileTestcaseBluprint,
// provides multiple testcase scenarios for invoking ReadingTuplesFromFile function.
var readingInputFromFileTestcases = []readingTuplesFromFileTestcaseBlueprint{
	{
		name:           "testcase1: reading tuples with correct format and values type",
		inputFilePath:  "./ForTesting/inputtest1.txt",
		expectedResult: []Tuple{{1, 2, 3}, {5, 6, 8}, {1, 6, 7}, {4, 3, 1}},
	},
	{
		name:           "testcase2: reading tuples with incorrect format and/or values type",
		inputFilePath:  "./ForTesting/inputtest2.txt",
		expectedResult: []Tuple{{4, 2, 1}, {2, 1, 3}},
	},
}

// TestReadingTuplesFromFile is a testing function to test the invokation of ReadingTuplesFromFile function,
// it checks if the said function is taking into consideration for the correctness of the format and the values type
// of the tuples before adding them to the list the will be returned.
func TestReadingTuplesFromFile(t *testing.T) {
	for _, testcase := range readingInputFromFileTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			testResult := ReadingTuplesFromFile(testcase.inputFilePath)
			if !reflect.DeepEqual(testResult, testcase.expectedResult) {
				t.Errorf("expected: %v, got: %v", testcase.expectedResult, testResult)
			}
		})
	}
}

// sortListTestcaseBlueprint is a blueprint for creating testcases for invoking the function SortList.
type sortListTestcaseBlueprint struct {
	name           string
	expectedResult []Tuple
	column         int
}

// sortListTestcases is a slice of objects of sortListTestcaseBluprint,
// provides multiple testcase scenarios for invoking SortList function.
var sortListTestcases = []sortListTestcaseBlueprint{
	{
		name:           "testcase1: sorting list by first column",
		expectedResult: []Tuple{{1, 2, 3}, {4, 5, 8}, {6, 1, 7}, {8, 4, 3}},
		column:         0,
	},
	{
		name:           "testcase1: sorting list by last column",
		expectedResult: []Tuple{{1, 2, 3}, {8, 4, 3}, {6, 1, 7}, {4, 5, 8}},
		column:         2,
	},
}

// TestSortList is a testing function that checks if the invokation of SortList is being done correctly
// with proper arguments being fed to the heap sort algorithm, and if it gets the proper result back.
func TestSortList(t *testing.T) {
	for _, testcase := range sortListTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			var listOfTuples = []Tuple{{4, 5, 8}, {1, 2, 3}, {8, 4, 3}, {6, 1, 7}}
			testResult := SortList(listOfTuples, testcase.column)
			if !reflect.DeepEqual(testResult, testcase.expectedResult) {
				t.Errorf("expected: %v, got: %v", testcase.expectedResult, testResult)
			}
		})
	}
}

// TestIntToString is a testing function to check whether or not a correct string format of casting
// tuple values is met properly when the function intToString gets invoked.
func TestIntToString(t *testing.T) {
	var expectedResult string = "1, 2, 3\n"
	testResult := intToString([]int{1, 2, 3})
	if testResult != expectedResult {
		t.Errorf("expected: %s, got: %s", expectedResult, testResult)
	}
}

// func TestSaveResultToFile(t *testing.T) {

// }
