package main

import (
	"log"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("SIN CONEXIÓN A BD")
		return
	}
	handlers.Manejadores()
}
