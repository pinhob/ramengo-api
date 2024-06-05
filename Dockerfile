FROM golang:1.22-alpine AS builder

COPY . .

RUN 	go build -o=/api cmd/api/main.go

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /api .

EXPOSE 8080

CMD ["./api"]