package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario ...
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar          string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Banner          string             `bson:"banner,omitempty" json:"banner,omitempty"`
	Ubicacion       string             `bson:"ubicacion,omitempty" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioweb,omitempty" json:"sitioweb,omitempty"`
}
