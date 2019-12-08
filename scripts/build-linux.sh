#~/bin/sh
echo "Building hellosvc for Linux"
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/linux/hellosvc cmd/hellosvc/main.go