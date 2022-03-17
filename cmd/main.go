package main

import (
	"github.com/RogerDurdn/ConfigurationService/http"
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"log"
)

func main() {
	log.Println("Initializing configuration service...")
	http.Server(lib.MockData{})
}