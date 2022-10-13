package main

import (
	"fmt"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

func main() {
	var inputFile string
	fmt.Print("Please provide a path to the input file: ")
	fmt.Scanln(&inputFile)

	// var list []sorting.Tuple
	list, err := sorting.ReadingTuplesFromFile(inputFile)
	if err != nil {
		fmt.Println("Error in Reading Tuples From File:", err.Error())
		return
	}
	var column int
	fmt.Print("Please provide a column to sort by: ")
	fmt.Scanln(&column)
	list, err = sorting.SortList(list, column)
	if err != nil {
		fmt.Println("Error in Sorting Tuples:", err.Error())
	}

	var outputFile string
	fmt.Print("Please provide a path to the output file: ")
	fmt.Scanln(&outputFile)
	err = sorting.SaveResultToFile(outputFile, list)
	if err != nil {
		fmt.Println("Error in Saving Result to Output File", err.Error())
	}
}
