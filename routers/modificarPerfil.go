package routers

import (
	"encoding/json"
	"net/http"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegisto(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al intentar modificar registro "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro modificar registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
