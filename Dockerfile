FROM golang:1.22-alpine

WORKDIR /app
COPY . /app

RUN cd ./cmd/api && go build -o /bin/api

EXPOSE 8080

CMD ["/bin/api"]

