package main

import (
	"log"

	"github.com/Aitugan/Techt/internal/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Panicln("Failed to start the server:", err)
	}
	server.Start()
}
