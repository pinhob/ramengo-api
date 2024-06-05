.DEFAULT_GOAL := run

.PHONY: build run tests image container start

build: 
	go build -o=bin/api cmd/api/main.go

run: build
	./bin/api

tests: 
	go test -v ./...

image:
	docker build --tag ramengo-api-img .

container:
	docker run -p 8080:8080 --name ramengo-api ramengo-api-img

start: 
	docker start ramengo-api