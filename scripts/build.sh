#! /bin/bash

go build -o bin/Randomizer src/main.go
export GOOS=windows
export GOARCH=amd64
go build -o bin/Randomizer.exe src/main.go
