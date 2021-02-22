package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN -> ConectarBD
var MongoCN = ConectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://pablo:1234@twittergolang.cocga.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConectarBD es la funcion que me permite conectar con la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection es el ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
