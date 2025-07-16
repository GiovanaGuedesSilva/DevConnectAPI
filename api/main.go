package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load() // Load environment variables and configuration
	r := router.Generate()
	fmt.Println("API is running on port", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)) // Start the HTTP server
}
