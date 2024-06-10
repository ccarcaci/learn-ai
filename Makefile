build:
	mkdir -p bin
	go build -o bin/learn-ai cmd/main.go

clean:
	go clean
	rm -rf bin

deps:
	go mod tidy

discovery-test:
	go test -v ./discovery -timeout 10s

# format check
fc:
	go fmt ./...
	go vet -structtag=false ./...

# static check
sc:
	make clean
	make deps
	make fc
	make build
	make unit-test

unit-test:
	go test -v ./internal/...
