FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 81
ENTRYPOINT ["go", "run", "main.go"]
