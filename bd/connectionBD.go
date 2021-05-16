package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//exporta la funcion de coneccion
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://root_admin:EJmgSA8Z7nmGSIFu@cluster0.apgbl.mongodb.net/test")

//conectar bd
func ConectarBD() *mongo.Client {
	//context sirve para comunicar entre ejecucion
	//aprender sobre context
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
	log.Println("Conexion exitosa con la BD")
	return client
}

/* ping a la base de datos */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
