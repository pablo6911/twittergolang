package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/pablo6911/twittergolang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BuscoPerfil busca un perfil en la BD
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, calcel := context.WithTimeout(context.Background(), time.Second*15)
	defer calcel()

	db := MongoCN.Database("twittergolang")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado")
	}
	return perfil, nil
}
