package main

import (
	"portfolio/server"
	"log"
	"net/http"
)

func main() {
    // Initialize router
    router := server.InitServer()

    log.Println("Server starting on port 4040...")
    err := http.ListenAndServe(":4040", router)
    if err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}