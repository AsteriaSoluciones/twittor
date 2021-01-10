package main

import (
	"log"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a BD")
		return
	}
	handlers.Manejadores()
}
