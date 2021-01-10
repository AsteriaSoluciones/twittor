package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/models"
)

//Registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	//Tomar los datos del cuerpo de la petición y convertirlas en un objeto Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese mail", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar usuario: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logró registrar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
