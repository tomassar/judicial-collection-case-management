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

COPY Makefile /app/Makefile

ENTRYPOINT ["make"]
#ENTRYPOINT CompileDaemon -build="go build -o main -buildvcs=false" -command="./main"