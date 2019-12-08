package service

import (
	"log"
	"net/http"
)

// StartHTTPServer starts a new http server
func StartHTTPServer(port string) {
	router := NewRouter()
	http.Handle("/", router)

	log.Println("Starting HTTP service on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured when starting HTTP on port " + port)
		log.Println("Error: " + err.Error())
	}
}
