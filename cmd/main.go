package main

import (
	"github.com/RogerDurdn/ConfigurationService/http"
	"log"
)

func main() {
	log.Println("Initializing configuration service...")
	http.Server()
}