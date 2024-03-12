# Start from a Golang base image
FROM golang:1.22.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

# Install PostgreSQL client
RUN apk add --no-cache postgresql-client


CMD ["./main"]


