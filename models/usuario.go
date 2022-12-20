package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Model ode usuario de la base de datos */
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`        //omitempty: si el campo viene vacio que no lo tenga en cuenta
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"` // bson datos de entrada a la base, json datos de salida al navegador
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb omitempty"`
}
