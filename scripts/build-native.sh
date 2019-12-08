#!/bin/sh
echo "Building hellosvc"
go build -o build/native/hellosvc cmd/hellosvc/main.go