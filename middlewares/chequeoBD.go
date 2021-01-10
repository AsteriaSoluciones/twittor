package middlewares

import (
	"net/http"

	"github.com/AsteriaSoluciones/twittor/bd"
)

//ChequeoBD verifica que la BD está accesible
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "No hay conexión con Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
