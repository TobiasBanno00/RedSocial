package routers

import (
	"errors"
	"strings"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

// extramos valores del token
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("IngEnSistemas")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
