package sorting

import (
	"errors"
	"reflect"
	"testing"
)

// testingDummy is a pointer to a MaxHeap object, which is going to be used in testing.
// var testingDummy *MaxHeap =

// TestNewHeap is a testing function for the NewHeap function, to check whether or not a MaxHeap pointer object
// is being created correctly.
func TestNewHeap(t *testing.T) {
	var expectedResult *MaxHeap = &MaxHeap{
		list:   []Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}},
		column: 0,
		sorted: false,
	}

	testResult := NewHeap([]Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}}, 0)

	if !reflect.DeepEqual(testResult, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, testResult)
	}
}

// heapifyTestcaseBluprint is a blueprint for creating testcases for invoking the method heapify.
type heapifyTestcaseBlueprint struct {
	name           string
	expectedResult []Tuple
	parameters     map[string]int
}

// heapifyTestcases is a slice of objects of heapifyTestcaseBluprint, provides multiple testcase scenarios
// for invoking heapify method.
var heapifyTestcases = []heapifyTestcaseBlueprint{
	{
		name:           "testcase1: column = 0, heapSize = 4, subtreeRootIndex = 0",
		expectedResult: []Tuple{{5, 6, 8}, {1, 2, 3}, {3, 2, 4}, {1, 7, 8}},
		parameters: map[string]int{
			"heapSize":         4,
			"subtreeRootIndex": 0,
		},
	},
	{
		name:           "testcase2: column = 0, heapSize = 4, subtreeRootIndex = 1",
		expectedResult: []Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}},
		parameters: map[string]int{
			"heapSize":         4,
			"subtreeRootIndex": 1,
		},
	},
	{
		name:           "testcase3: column = 1, heapSize = 4, subtreeRootIndex = 0",
		expectedResult: []Tuple{{5, 6, 8}, {1, 7, 8}, {3, 2, 4}, {1, 2, 3}},
		parameters: map[string]int{
			"heapSize":         4,
			"subtreeRootIndex": 0,
		},
	},
}

// TestHeapify is a testing function for the heapify method, to test out three case scenarios on invoking heapify method
// and check whether or not the expected values are met after applying the heapify part of heap sort algorithm.
func TestHeapify(t *testing.T) {
	for index, testcase := range heapifyTestcases {
		var testingHeap *MaxHeap = &MaxHeap{
			list:   []Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}},
			column: 0,
			sorted: false,
		}
		if index == 2 {
			testingHeap.column = 1
		}
		t.Run(testcase.name, func(t *testing.T) {
			var expectedResult []Tuple = testcase.expectedResult
			testingHeap.heapify(testcase.parameters["heapSize"], testcase.parameters["subtreeRootIndex"])
			testResult := testingHeap.list
			if !reflect.DeepEqual(testResult, expectedResult) {
				t.Errorf("expected: %v, got: %v", expectedResult, testResult)
			}
		})
	}
}

// heapSortTestcaseBluprint is a blueprint for creating testcases for invoking the method heapSort.
type heapSortTestcaseBlueprint struct {
	name           string
	expectedResult []Tuple
}

// heapSortTestcases is a slice of objects of heapSortTestcaseBluprint, provides multiple testcase scenarios
// for invoking heapSort method.
var heapSortTestcases = []heapSortTestcaseBlueprint{
	{
		name:           "testcase1: column = 0",
		expectedResult: []Tuple{{1, 2, 3}, {1, 7, 8}, {3, 2, 4}, {5, 6, 8}},
	},
	{
		name:           "testcase2: column = 2",
		expectedResult: []Tuple{{1, 2, 3}, {3, 2, 4}, {1, 7, 8}, {5, 6, 8}},
	},
}

// TestHeapSort is a testing function for the heapSort method, to test out two sorting scenarios
// based on different column values by invoking heapSort method which represents the complete heap sort algorithm,
// and check whether or not the expected final results are met after applying this strategy.
func TestHeapSort(t *testing.T) {
	for index, testcase := range heapSortTestcases {
		var testingHeap *MaxHeap = &MaxHeap{
			list:   []Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}},
			column: 0,
			sorted: false,
		}
		if index == 1 {
			testingHeap.column = 2
		}
		t.Run(testcase.name, func(t *testing.T) {
			testingHeap.heapSort()
			testResult := testingHeap.list
			if !reflect.DeepEqual(testResult, testcase.expectedResult) {
				t.Errorf("expected: %v, got: %v", testcase.expectedResult, testResult)
			}
		})
	}
}

// getSortedListTestcaseBluprint is a blueprint for creating testcases for invoking the method getSortedList.
type getSortedListTestcaseBlueprint struct {
	name           string
	expectedResult []Tuple
	expectedError  error
}

// getSortedListTestcases is a slice of objects of getSortedListTestcaseBluprint, provides multiple testcase scenarios
// for invoking getSortedList method.
var getSortedListTestcases = []getSortedListTestcaseBlueprint{
	{
		name:           "testcase1: sorted = false, error = the passed list is not sorted yet",
		expectedResult: nil,
		expectedError:  errors.New("the passed list is not sorted yet"),
	},
	{
		name:           "testcase2: sorted = true, no errors",
		expectedResult: []Tuple{{1, 2, 3}, {1, 7, 8}, {3, 2, 4}, {5, 6, 8}},
		expectedError:  nil,
	},
}

// TestGetSortedList is a testing function for the getSortedList method, to test out two scenarios on invoking the said method
// one where the list is not sorted, and the other when it's sorted,
// and check whether or not the correct return values are returned and the error is handled correctly.
func TestGetSortedList(t *testing.T) {
	var testingHeap *MaxHeap = &MaxHeap{
		list:   []Tuple{{1, 2, 3}, {5, 6, 8}, {3, 2, 4}, {1, 7, 8}},
		column: 0,
		sorted: false,
	}
	for index, testcase := range getSortedListTestcases {
		if index == 1 {
			testingHeap.heapSort()
		}
		t.Run(testcase.name, func(t *testing.T) {
			testResult, testError := testingHeap.getSortedList()
			if !reflect.DeepEqual(testResult, testcase.expectedResult) && testError.Error() != testcase.expectedError.Error() {
				t.Errorf("expected: (result: %v, error: %v), got: (result: %v, error: %v)", testcase.expectedResult,
					testcase.expectedError.Error(), testResult, testError.Error())
			}
		})
	}
}
