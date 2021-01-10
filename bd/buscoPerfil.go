package bd

import (
	"context"
	"time"

	"github.com/AsteriaSoluciones/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BuscoPerfil ...
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": objID}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	return perfil, err
}
