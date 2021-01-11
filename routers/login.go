package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/jwt"
	"github.com/AsteriaSoluciones/twittor/models"
)

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Usuario y/o contraseña no válidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña no válidos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al generar token: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Grabar una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
