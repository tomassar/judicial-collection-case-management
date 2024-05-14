# Makefile

# Define the build command
build:
	CompileDaemon -build="go build -o main ./cmd/server/main.go -buildvcs=false" -command="./main"

# Define the default target to run both build and watch in parallel
all:
	build
