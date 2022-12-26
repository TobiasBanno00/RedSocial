package routers

import (
	"encoding/json"
	"net/http"

	"github.com/TobiasBanno00/RedSocial/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrior un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
