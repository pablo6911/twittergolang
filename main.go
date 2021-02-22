package main

import (
	"log"

	"github.com/pablo6911/twittergolang/bd"
	"github.com/pablo6911/twittergolang/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
