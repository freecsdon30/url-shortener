.PHONY: build 
build:
	mkdir -p ./build
	go build -o ./build ./cmd/shortener/main.go

.PHONY: clean 
clean:
	rm -rf ./build 

.PHONY: run 
run:
	./build/main

