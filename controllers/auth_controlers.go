package controllers

import (
	"encoding/json"
	"golang_jwt/helpers"
	"golang_jwt/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r Body.Close()

	if register.Password != register.password_confirm {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	
}
