package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AsteriaSoluciones/twittor/bd"
)

//LeoTweets ...
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	pag := r.URL.Query().Get("pag")
	if len(pag) < 1 {
		http.Error(w, "Debe enviar el parámetro pag", http.StatusBadRequest)
		return
	}

	paginaStr, err := strconv.Atoi(pag)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pag como numérico", http.StatusBadRequest)
		return
	}

	pagina := int64(paginaStr)
	respuesta, correcto := bd.LeoTweets(ID, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
