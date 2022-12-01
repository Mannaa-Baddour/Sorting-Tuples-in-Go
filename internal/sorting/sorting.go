// Package sorting reads through a text file containing a three elements tuple, sorts them based on a certain column
// and outputs the result to a new file.
//
// The sorting package deals with a tuple of three elements integers formatted as: int, int, int
// other values such as runes and floats, or other string formats are met with an error message,
// and they won't be taken into consideration while forming the result.

package sorting

import (
	"log"
	"os"
)

// Tuple is a defined type to simulate placeholder for three integer elements,
// a condition taken care of in the associated function.
//type Tuple []int

// appendOne function is associated to the defined type Tuple.
// It takes an integer number as an arguement, appends it to the defined type Tuple if it's not already full (with 3 numbers).
// Otherwise it'll return an error.
//func (tuple *Tuple) appendOne(number int) error {
//	if len(*tuple) == 3 {
//		return errors.New("warning: cannot add anymore numbers, tuple is already at its maximum capacity")
//	}
//	*tuple = append(*tuple, number)
//	return nil
//}

// ReadingTuplesFromFile opens and reads the input from a file specified by its name in the argument fileName.
// Checks for errors in the process of opening the file, then reads it line by line,
// and uses a helper function to cast the strings into ints for the tuple.
// If the tuple consists of 3 elements (which is the main condition), it adds it to the list of tuples,
// otherwise, it logs an error and skips the current tuple.
//func ReadingTuplesFromFile(fileName string) ([]Tuple, error) {
//	fmt.Printf("Reading input from file %s\n", fileName)
//	var listOfTuples []Tuple
//	var file *os.File
//	var err error
//	file, err = os.Open(fileName)
//	// HandleErrors(err, file, true)
//	err = HandleErrors(err, file, false)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//	var scanner *bufio.Scanner = bufio.NewScanner(file)
//	for scanner.Scan() {
//		var line string = scanner.Text()
//		// tuple := stringToInt(line)
//		tuple, err := stringToInt(line)
//		if err != nil {
//			return nil, err
//		}
//		if len(tuple) == 3 {
//			listOfTuples = append(listOfTuples, tuple)
//		} else {
//			err = errors.New("warning: cannot append tuple, length of tuple is not correct")
//			// HandleErrors(err, nil, false)
//			returnedErr := HandleErrors(err, nil, false)
//			if returnedErr != nil {
//				return nil, returnedErr
//			}
//		}
//	}
//	return listOfTuples, nil
//}

// HandleErrors logs the error passed as an argument to it, checks whether or not the execution of some commands is ok.
// Logging will be output to a file log.txt.
// The second argument fileToCheck, specifies whether there is a file that needs to be closed due to an error or not.
// The last argument fatal, specifies whether the execution needs to be terminated or not.
func HandleErrors(err error, fileToCheck *os.File, fatal bool) error {
	var logFile *os.File
	var logFileErr error
	var logFileOptions int = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	logFile, logFileErr = os.OpenFile("logs.txt", logFileOptions, 0666)
	if logFileErr != nil {
		log.Println(logFileErr)
		logFile.Close()
		return logFileErr
		// os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	if err != nil {
		log.Println(err)
		if fileToCheck != nil {
			fileToCheck.Close()
		}
		if fatal {
			os.Exit(1)
		}
		return err
	}
	return nil
}

// stringToInt is a helper function that checks for a specific expression to make sure that the condition of a correct
// string formatting is met or not, then casts the values passed as a string argument into integer values
// to be stored later on in a tuple which will be returned.
//func stringToInt(line string) (Tuple, error) {
//	var match bool
//	var err error
//	var tuple Tuple
//	match, err = regexp.MatchString(`^\d,\s?\d,\s?\d$`, line)
//	// HandleErrors(err, nil, true)
//	err = HandleErrors(err, nil, false)
//	if err != nil {
//		return nil, err
//	}
//	if match {
//		var stringParts []string = strings.Split(line, ",")
//		for _, value := range stringParts {
//			value = strings.TrimSpace(value)
//			var intValue int
//			intValue, err = strconv.Atoi(value)
//			// HandleErrors(err, nil, true)
//			err = HandleErrors(err, nil, false)
//			if err != nil {
//				return nil, err
//			}
//			err = tuple.appendOne(intValue)
//			// HandleErrors(err, nil, false)
//			err = HandleErrors(err, nil, false)
//			if err != nil {
//				return nil, err
//			}
//		}
//	} else {
//		fmt.Printf("Current line \"%s\" doesn't match the correct format (int, int, int)\n", line)
//	}
//	return tuple, nil
//}

// SortList takes a list of tuples as an argument to sort through it based on the second argument column.
// these two values are passed in to create a heap and start the sorting process according the heap sort algorithm,
// then the final sorted list of tuples is returned.
//func SortList(list []Tuple, column int) ([]Tuple, error) {
//	var heap *MaxHeap = NewHeap(list, column)
//	heap.heapSort()
//	var sortedList []Tuple
//	var err error
//	sortedList, err = heap.getSortedList()
//	// HandleErrors(err, nil, false)
//	err = HandleErrors(err, nil, false)
//	if err != nil {
//		return nil, err
//	}
//	return sortedList, nil
//}

// intToString is a helper function which casts the integer values of a tuple passed as an argument
// to a line of type string according to the format: "%d, %d, %d", and then returns it.
//func intToString(tuple Tuple) string {
//	var line string
//	for index, value := range tuple {
//		if index == 2 {
//			line += strconv.Itoa(value) + "\n"
//		} else {
//			line += strconv.Itoa(value) + ", "
//		}
//	}
//	return line
//}

// SaveResultToFile Saves the desired final result of list of tuples passed as an argument,
// to a file specified by the argument fileName which represents its path.
//func SaveResultToFile(fileName string, result []Tuple) error {
//	fmt.Printf("Saving final result to %s\n", fileName)
//	var file *os.File
//	var err error
//	file, err = os.Create(fileName)
//	// HandleErrors(err, file, true)
//	err = HandleErrors(err, file, false)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	for _, tuple := range result {
//		var line string = intToString(tuple)
//		file.WriteString(line)
//	}
//	fmt.Println("Operation completed successfully")
//	return nil
//}
