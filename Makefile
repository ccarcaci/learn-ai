build:
	mkdir -p bin
	go build -o bin/leai main.go

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

run-iris-dataset:
	make build
	./bin/leai iris-recognition --iris-dataset-path datasets/iris.data

unit-test-perceptron:
	go test -v ./perceptron/...
