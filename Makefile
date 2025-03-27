build:
	go build -o bin/go-blockchain-simple

run: build
	./bin/go-blockchain-simple

test:
	go test -v ./...