package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/AsteriaSoluciones/twittor/bd"
	"github.com/AsteriaSoluciones/twittor/models"
)

//Email ...
var Email string

//IDUsuario ...
var IDUsuario string

//ProcesoToken ...
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Formato de token no válido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = ID
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("Token no válido")
	}
	return claims, false, "", err
}
