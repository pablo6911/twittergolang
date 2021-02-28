package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pablo6911/twittergolang/bd"
	"github.com/pablo6911/twittergolang/models"
)

//ModificoPerfil modifica el perfil de usuario
func ModificoPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se halogrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
