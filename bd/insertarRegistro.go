package bd

import (
	"context"
	"time"

	"github.com/pablo6911/twittergolang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarRegistro es la parada final con la BD para insertar los datos del user*/
func InsertarRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittergolang")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//pasando objectId a string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
