init:
	go mod init leai

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

read-mnist-image-discovery:
	go test -timeout 30s -v -run ^TestDiscoveryReadTestingImages github.com/ccarcaci/learn-ai/inputs/mnist

run-iris-dataset-batch-perceptron:
	make build
	./bin/leai iris-recognition \
		--iris-dataset-path datasets/iris.data \
		--training-type batch \
		--activation-type perceptron \
		--eta 0.01 \
		--epochs 100 \
		--training-ratio 0.8

run-iris-dataset-online-perceptron:
	make build
	./bin/leai iris-recognition \
		--iris-dataset-path datasets/iris.data \
		--training-type online \
		--activation-type perceptron \
		--eta 0.01 \
		--epochs 100 \
		--training-ratio 0.8

run-iris-dataset-batch-adaline:
	make build
	./bin/leai iris-recognition \
		--iris-dataset-path datasets/iris.data \
		--training-type batch \
		--activation-type adaline \
		--eta 0.01 \
		--epochs 100 \
		--training-ratio 0.8

run-iris-dataset-online-adaline:
	make build
	./bin/leai iris-recognition \
		--iris-dataset-path datasets/iris.data \
		--training-type online \
		--activation-type adaline \
		--eta 0.0001 \
		--epochs 100 \
		--training-ratio 0.8

unit-test:
	go test -v ./...
