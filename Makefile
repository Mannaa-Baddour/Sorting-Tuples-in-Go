# Build all
build: build-sort build-server build-client

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

# Test server
test-server:
	@echo "-- Testing target: server"
	go test -v --cover github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/cmd/server

# Run client
run-client:
	@echo "-- Running client to request from server with specified parameters"
	go run cmd/client/client.go -server http://127.0.0.1 -port 30010 \
	-infile cmd/server/test.txt -outfile cmd/server/result.txt -sort-column 2

# Build client
build-client:
	@echo "-- Building target: client"
	go build -o bin/client cmd/client/client.go

# Test client
test-client:
	@echo "-- Testing target: client"
	go test -v --cover github.com/Mannaa-Baddour/Sorting-Tuples-in-Go/cmd/client

# Clean build
clean:
	rm -r bin