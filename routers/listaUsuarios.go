package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/estefania906/twittor/bd"
)

/*ListaUsuarios leo la lista de los usuarios*/
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTem, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTem)

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
