package main

import (
	"fmt"
	"packages/sorting"
)

/* First test is run normally to show that the sorting is working
The Second test introduces runes and floats which get disregarded by regex
The Third test sorts the tuples based on the second column even though the form of the spaces isn't quite correct */

func main() {
	var index int = 1
	for index <= 3 {
		var column int = 0
		if index == 3 {
			column = 1
		}
		var inputFile string = fmt.Sprintf("../sorting/testfile%d.txt", index)
		var list []sorting.Tuple

		list = sorting.ReadingTuplesFromFile(inputFile)
		list = sorting.SortList(list, column)

		var outputFile string = fmt.Sprintf("../sorting/result%d.txt", index)
		sorting.SaveResultToFile(outputFile, list)
		index++
	}
}
