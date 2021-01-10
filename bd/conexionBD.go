package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexión a base de datos
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://benjamin:Tsukete1@cluster0.624ja.mongodb.net/twittor?retryWrites=true&w=majority")

func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa a la BD")
	return client
}

//ChequeoConexion verifica que la base de datos esté accesible
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
