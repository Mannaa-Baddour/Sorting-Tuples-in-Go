package logging

import (
	"fmt"
	"log"
	"os"
)

// logFilePath is the path for the log file where all logs are printed in
var logFilePath = "../../internal/logs/logs.txt"

// LogError is a function that formats the error and where it occurred, and it adds it to the logs file
// specified in the variable logFilePath.
func LogError(executionError error, location string) {
	fullError := fmt.Sprintf("%v, Occurred at: %s", executionError, location)
	options := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	logFile, err := os.OpenFile(logFilePath, options, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println(fullError)
}
