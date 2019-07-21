package main

import (
	"log"

	"github.com/rcdmk/shortest-flight-path/infra/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	return
}
