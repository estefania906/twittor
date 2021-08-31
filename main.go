package main

import (
	"log"

	"github.com/estefania906/twittor/bd"
	"github.com/estefania906/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
