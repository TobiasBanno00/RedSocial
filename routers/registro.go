package routers

import (
	"encoding/json"
	"net/http"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/models"
)

/* Func para crear en la bd el registro de usuario */

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //decodifica el body en t
	if err != nil {
		http.Error(w, "Error en datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 { // si el largo del email es 0
		http.Error(w, "Email de usuario requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password tiene que ser mayor que 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el Registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
