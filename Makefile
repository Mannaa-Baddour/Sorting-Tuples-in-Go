# Build all
build: build-sort build-server

# Run sort
run-sort:
	@echo "-- Running target: sort"
	go run cmd/sort/sort.go

# Build sort
build-sort:
	@echo "-- Building target: sort"
	go build -o bin/sort cmd/sort/sort.go

# Run server
run-server:
	@echo "-- Running server with specified parameters"
	go run cmd/server/server.go -host localhost -port 30010

# Build server
build-server:
	@echo "-- Building target: server"
	go build -o bin/server cmd/server/server.go

# Clean build
clean:
	rm -r bin