package bd

import (
	"context"
	"time"

	"github.com/TobiasBanno00/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* parada final con la bd para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redSocial")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
