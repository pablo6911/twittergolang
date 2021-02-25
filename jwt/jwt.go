package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pablo6911/twittergolang/models"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("twittergolang")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenstr, err := token.SignedString(miClave)
	if err != nil {
		return tokenstr, err
	}
	return tokenstr, nil
}
