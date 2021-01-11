package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/models"
)

//ModificarPerfil ...
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Datos incorrectos: "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al modificar registro: "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No fue posible modificar registro", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
