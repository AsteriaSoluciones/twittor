package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/AsteriaSoluciones/twittor/middlewares"
	"github.com/AsteriaSoluciones/twittor/routers"
)

//Manejadores ...
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")

	//Buscar el puerto en variables de entorno, si no se encuentra establecerlo
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//Establecer configuración CORS
	handler := cors.AllowAll().Handler(router)

	//Levantar servidor
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
