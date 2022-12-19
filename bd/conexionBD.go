package bd

import (
	"context" //espacio de memoria donde se pueden compartir diferentes cosas
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()                                                                                                                          // url que va a conectar
var clientOptions = options.Client().ApplyURI("mongodb+srv://TobiasBanno:ContraMDB1234@redsocial.mvci4kr.mongodb.net/?retryWrites=true&w=majority") // setea URL de base de datos

func ConectarBD() *mongo.Client { // fun que me permite contectar la bd
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error()) // usamos fun .Error ya que lo convierte el obj en string
		return client
	}
	err = client.Ping(context.TODO(), nil) //ve si la conexion está disponible

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connect a BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
