package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TobiasBanno00/RedSocial/bd"
	"github.com/TobiasBanno00/RedSocial/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar isertar registro "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se insertó tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
