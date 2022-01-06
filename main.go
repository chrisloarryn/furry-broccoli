package main

import (
	"log"

	"github.com/ccontreras/furry-broccoli/bd"
	"github.com/ccontreras/furry-broccoli/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("There is no connection to DB")
		return
	}

	handlers.Handlers()
}
