package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("tuples.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	sort.Strings(list)

	file, err = os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, value := range list {
		file.WriteString(value + "\n")
	}

}
