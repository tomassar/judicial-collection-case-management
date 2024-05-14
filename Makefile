# Makefile

# Define the build command
build:
	CompileDaemon -build="go build -o main -buildvcs=false" -command="./main"

# Define the default target to run both build and watch in parallel
all:
	build
