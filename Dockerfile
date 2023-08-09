# Use the official Go image as the base image
FROM golang:1.20-alpine AS golang

# Set the working directory inside the container
WORKDIR /go-echo-server-template

COPY go.mod go.sum ./
RUN go mod download

# Copy the Go service source code into the container
COPY . .

# Build the Go service
RUN go build -o app ./cmd

FROM debian:11-slim

WORKDIR /usr/app

COPY --from=golang /go-echo-server-template/envs ./envs
COPY --from=golang /go-echo-server-template/app .

CMD ["./app"]
