package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pablo6911/twittergolang/bd"
	"github.com/pablo6911/twittergolang/models"
)

//Registro es la func para crear en la BD el registro de user
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requqerido", 400)
		return
	}
	if len(t.Password) > 6 {
		http.Error(w, "Desbes especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	//Hacer validaxion contra los datos
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un Usuario con ese Email", 400)
		return
	}
	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Resgistrar el Usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
