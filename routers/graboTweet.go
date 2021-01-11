package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/models"
)

//GraboTweet ...
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Datos incorrectos: "+err.Error(), http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al intentar insertar el registro: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se insertó el tweet: ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
