package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	costo := 8 // esto elevado al cuadrado es la cantidad de pasadas que hace para encriptar, mayor costo mas seguridad
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
