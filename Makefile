.DEFAULT_GOAL := run

.PHONY: build run tests image container start

build: 
	go build -o=bin/api cmd/api/main.go

run: build
	./bin/api

tests: 
	go test -v ./...

image:
	sudo docker build --tag ramengo-api-img .

container:
	sudo docker run -p 8080:8080 --name ramengo-api ramengo-api-img

start: 
	sudo docker start ramengo-api