package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AsteriaSoluciones/twittor/bd"
)

//LeoTweetsRelacion ...
func LeoTweetsRelacion(w http.ResponseWriter, r *http.Request) {
	pag := r.URL.Query().Get("pag")
	if len(pag) < 1 {
		http.Error(w, "Debe enviar el parámetro pag", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(pag)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pag como entero mayor a cero", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
