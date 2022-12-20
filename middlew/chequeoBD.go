package middlew // los middlewares reciben algo y devuelven lo mismo

import (
	"net/http"

	"github.com/TobiasBanno00/RedSocial/bd"
)

/* Middlew que me permite saber el estado de la bd */

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdida con base de datos", 500)
			return
		}
		next.ServeHTTP(w, r) // le paso lo que recibí al proximo eslabón
	}
}
