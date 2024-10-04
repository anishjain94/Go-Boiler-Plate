run: build
	./server 

build:
	go build -o server cmd/main.go