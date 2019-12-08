package main

import (
	"log"
	"os"

	"github.com/mohaila/helloservice/internal/service"
)

const (
	svcName = "accountsvc"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "50080"
	}

	log.Println("Starting " + svcName)
	service.StartHTTPServer(port)
}
