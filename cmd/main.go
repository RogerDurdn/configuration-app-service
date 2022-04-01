package main

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib"
	"github.com/RogerDurdn/ConfigurationService/rest"
	"log"
)

func main() {
	log.Println("Initializing configuration service...")
	rest.Server(lib.MongoData{})
}
