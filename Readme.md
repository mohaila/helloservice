# Tiny Go Sevice
A simple template for a Go service with a single /status route.<br>
The Linux build can be deployed with Docker using Scracth image.
Building static binary for Linux is simple whem using macOS and Windows.<br>
When on Linux, additional flags must be needed for building static binary.
Commands available with make:
- linux: build static binary for Linux
- native: build binary for actual OS (Windows, macOS, Linux, etc...)
- image: build Docker image using Linux static binary, depends on linux target
- deploy: deploy Docker image using Docker Compose, depends on image
- clean: clean build folder

The default target is native.
