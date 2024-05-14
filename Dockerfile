FROM golang:latest

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Basic setup of the container
RUN mkdir /app
COPY . /app
WORKDIR /app

# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Set WORKDIR to the directory containing main.go
WORKDIR /app/cmd/server

ENTRYPOINT CompileDaemon -build="go build -o /app/main -buildvcs=false" -command="/app/main"
