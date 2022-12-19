package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" //enrutador y despachador de solicitudes
	"github.com/rs/cors"     //permisos que le do a mi api para que sea accesible desde cualquier lugar
)

func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT") // devuelve el valor de la clave de la variable de entorno, si existe
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)        //le damos permisos a cualquiera
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //pongo al servidor a escuchar el puerto
}
