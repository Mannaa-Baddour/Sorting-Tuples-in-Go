package main

import (
	"fmt"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/sorting"
)

func main() {
	var inputFile string
	fmt.Print("Please provide a path to the input file: ")
	fmt.Scanln(&inputFile)

	var list []sorting.Tuple
	list = sorting.ReadingTuplesFromFile(inputFile)
	var column int
	fmt.Print("Please provide a column to sort by: ")
	fmt.Scanln(&column)
	list = sorting.SortList(list, column)

	var outputFile string
	fmt.Print("Please provide a path to the output file: ")
	fmt.Scanln(&outputFile)
	sorting.SaveResultToFile(outputFile, list)
}
