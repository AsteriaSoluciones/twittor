package bd

import (
	"context"
	"time"

	"github.com/AsteriaSoluciones/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro de usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//Recuperar ID del usuario creado y convertirlo a string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
