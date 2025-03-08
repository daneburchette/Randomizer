#! /bin/bash

# Delete bin/ contents for clean build
echo "Clearing ./bin/ directory"
rm bin/*
# Build Linux and Windwos binaries
echo "Building Linux Binary"
go build -o bin/Randomizer src/main.go
export GOOS=windows
export GOARCH=amd64
echo "Building Windows Binary"
go build -o bin/Randomizer.exe src/main.go
