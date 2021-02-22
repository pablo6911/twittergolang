package bd

import (
	"context"
	"time"

	"github.com/pablo6911/twittergolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittergolang")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultdo models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultdo)
	ID := resultdo.ID.Hex()
	if err != nil {
		return resultdo, false, ID
	}
	return resultdo, true, ID
}
