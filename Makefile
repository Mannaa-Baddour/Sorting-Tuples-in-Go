# Build all
build: build-sort

# Build sort
build-sort:
	@echo "-- Building target: sort"
	go build -o bin/sort cmd/sort/sort.go

clean:
	rm -r bin