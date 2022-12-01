// Package sorting reads through data read from a csv file and formatted as a slice of Content,
// each containing three columns, and sorts them based on a certain column using heapsort algorithm,
// then returns the sorted result.
//
// The sorting package deals with a tuple of three elements integers formatted as: int, int, int
// other values such as runes and floats, or other string formats are met with an error message,
// and they won't be taken into consideration while forming the result.

package sorting

import (
	"errors"
	"strconv"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/models"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/logging"
)

// Tuple is a defined type to simulate placeholder for three integer elements,
// a condition taken care of in the associated function.
type Tuple []int

// appendOne function is associated to the defined type Tuple.
// It takes an integer number as an arguement, appends it to the defined type Tuple if it's not already full (with 3 numbers).
// Otherwise, it'll return an error.
func (tuple *Tuple) appendOne(number int) error {
	if len(*tuple) == 3 {
		err := errors.New("warning: cannot add anymore numbers, tuple is already at its maximum capacity")
		logging.LogError(err, "sorting/appendOne")
		return err
	}
	*tuple = append(*tuple, number)
	return nil
}

// stringToInt is a helper function that checks for a specific expression to make sure that the condition of a correct
// string formatting is met or not, then casts the values passed as a string argument into integer values
// to be stored later on in a tuple which will be returned.
func stringToInt(lines []models.Content) ([]Tuple, error) {
	var tuples []Tuple
	for _, line := range lines {
		var tuple Tuple
		for _, value := range []string{line.Column0, line.Column1, line.Column2} {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				logging.LogError(err, "sorting/stringToInt : converting string value to int condition")
				return nil, err
			}
			err = tuple.appendOne(intValue)
			if err != nil {
				return nil, err
			}
		}
		tuples = append(tuples, tuple)
	}
	return tuples, nil
}

// sortList takes a list of tuples as an argument to sort through it based on the second argument column.
// these two values are passed in to create a heap and start the sorting process according the heap sort algorithm,
// then the final sorted list of tuples is returned.
func sortList(tuples []Tuple, column int) ([]Tuple, error) {
	heap := NewHeap(tuples, column)
	heap.heapSort()
	sortedList, err := heap.getSortedList()
	if err != nil {
		return nil, err
	}
	return sortedList, nil
}

// intToString is a helper function which casts the integer values of a tuple passed as an argument
// to a line of type string according to the format: "%d, %d, %d", and then returns it.
func intToString(tuples []Tuple) []models.Content {
	var lines []models.Content
	for _, tuple := range tuples {
		var line models.Content
		columns := []*string{&line.Column0, &line.Column1, &line.Column2}
		for index, value := range tuple {
			strValue := strconv.Itoa(value)
			*columns[index] = strValue
		}
		lines = append(lines, line)
	}
	return lines
}

// Sort performs the full sorting functionality on the data specified based on the passed sorting column.
func Sort(data []models.Content, column int) ([]models.Content, error) {
	tuples, err := stringToInt(data)
	if err != nil {
		return nil, err
	}
	sorted, err := sortList(tuples, column)
	if err != nil {
		return nil, err
	}
	result := intToString(sorted)
	return result, nil
}
