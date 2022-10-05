package sorting

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// a defined type to simulate a tuple with three elements as sepecified in the associated function
type Tuple []int

func (tuple *Tuple) appendOne(number int) error {
	/* This function appends one number to the defined type Tuple if it's not already full (with 3 numbers)
	otherwise it'll return an error that will be logged using the function HandleErrors */
	if len(*tuple) > 3 {
		return errors.New("cannot add anymore numbers")
	}
	*tuple = append(*tuple, number)
	return nil
}

func ReadingTuplesFromFile(fileName string) []Tuple {
	/* Reads the input from the given files specified in the argument fileName, checks for errors in the process of
	opening the file, the reads it line by line, and uses a helper function to cast the strings into ints for the tuple
	then if the tuple consists of 3 elements (the main condition), it adds it to the list */
	var listOfTuples []Tuple
	var file *os.File
	var err error
	file, err = os.Open(fileName)
	handleErrors(err)
	defer file.Close()
	var scanner *bufio.Scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		var tuple Tuple = stringToInt(line)
		if len(tuple) == 3 {
			listOfTuples = append(listOfTuples, tuple)
		}
	}
	return listOfTuples
}

func handleErrors(err error) {
	// Logging the error passed as an argument to it, checks whether or not everything's ok
	if err != nil {
		log.Fatal(err)
	}
}

func stringToInt(line string) Tuple {
	/* a helper function that checks for a specific expression to make sure that the condition is met
	then casts the string values into integer values to be stored in a tuple */
	var regExp *regexp.Regexp
	var err error
	regExp, err = regexp.Compile(`^\d,\s?\d,\s?\d$`)
	handleErrors(err)
	var tuple Tuple
	if regExp.MatchString(line) {
		var stringParts []string = strings.Split(line, ",")
		for _, value := range stringParts {
			value = strings.TrimSpace(value)
			var intValue int
			intValue, err = strconv.Atoi(value)
			handleErrors(err)
			tuple.appendOne(intValue)
		}
	}
	return tuple
}

func SortList(list []Tuple, column int) []Tuple {
	/* Sorts the list of tuples based on the columns specified to be sorted by
	the column values and the tuples are bound with a map, with respect to the condition of non-unique keys */
	var mapping map[int]Tuple = make(map[int]Tuple)
	for _, segment := range list {
		if _, ok := mapping[segment[column]]; ok {
			mapping[segment[column]+1] = segment
		} else {
			mapping[segment[column]] = segment
		}
	}
	var keys []int = make([]int, 0, len(mapping))
	for key := range mapping {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	var final []Tuple
	for _, key := range keys {
		final = append(final, mapping[key])
	}
	return final
}

func intToString(tuple Tuple) string {
	// a helper function which casts the integer values of a tuple to a line of type string according to certain format
	var str string
	for index, value := range tuple {
		if index == 2 {
			str += strconv.Itoa(value) + "\n"
		} else {
			str += strconv.Itoa(value) + ", "
		}
	}
	return str
}

func SaveResultToFile(fileName string, result []Tuple) {
	// Saves the desired final result of list of tuples to a file specified as an argument
	var file *os.File
	var err error
	file, err = os.Create(fileName)
	handleErrors(err)
	defer file.Close()
	for _, tuple := range result {
		var line string = intToString(tuple)
		file.WriteString(line)
	}
}
