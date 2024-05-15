FROM golang:latest

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \ 
    GOFLAGS="-buildvcs=false"

# Basic setup of the container
RUN mkdir /app
COPY . /app
WORKDIR /app

# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build /app/cmd/server/main.go" -command="/app/main"
