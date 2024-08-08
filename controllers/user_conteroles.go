package controllers

import (
	"golang_jwt/helpers"
	"golang_jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	userResponse := &models.MyProfile{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
