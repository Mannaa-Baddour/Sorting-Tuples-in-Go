package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/config"
	"github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/internal/api/routes"
)

// main is a starting point function that takes a URL and a port as cmd string argument to run the server at that address
// while handling any errors that might occur related to the URL or the port in use.
func main() {
	fmt.Println("Enter Host Name (localhost): ")
	fmt.Scanln(&config.Host)
	if config.Host == "" {
		config.Host = "localhost"
	}
	fmt.Println("Enter Port Number (5432): ")
	fmt.Scanln(&config.Port)
	if config.Port == "" {
		config.Port = "5432"
	}
	fmt.Println("Enter User Name (postgres): ")
	fmt.Scanln(&config.User)
	if config.User == "" {
		config.User = "postgres"
	}
	fmt.Println("Enter Password: ")
	fmt.Scanln(&config.Password)
	fmt.Println("Enter Database Name (users_files_db): ")
	fmt.Scanln(&config.DBname)
	if config.DBname == "" {
		config.DBname = "users_files_db"
	}
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		err := http.ListenAndServe(":30010", mux)
		if err != nil {
			log.Fatal(err)
		}
		waitGroup.Done()
	}()
	fmt.Println("Starting Server at: localhost:30010")
	waitGroup.Wait()
}
