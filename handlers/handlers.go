package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/TobiasBanno00/RedSocial/middlew"
	"github.com/TobiasBanno00/RedSocial/routers"
	"github.com/gorilla/mux" //enrutador y despachador de solicitudes
	"github.com/rs/cors"     //permisos que le doy a mi api para que sea accesible desde cualquier lugar
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT") // devuelve el valor de la clave de la variable de entorno, si existe
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)        //le damos permisos a cualquiera
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //pongo al servidor a escuchar el puerto
}
