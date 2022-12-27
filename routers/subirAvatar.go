package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar") // procesamos como formulario
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_CREATE|os.O_WRONLY, 0666) // creando archivo
	if err != nil {
		http.Error(w, "Error al subir img "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file) //lo que copio y en donde
	if err != nil {
		http.Error(w, "Error copiando img  "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegisto(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en BD  "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
