# Run server
run-server:
	@echo "-- Running server at localhost:30010"
	go run cmd/server/server.go

# Build server
build-server:
	@echo "-- Building target: server"
	go build -o bin/server cmd/server/server.go

# Clean build
clean:
	rm -r bin