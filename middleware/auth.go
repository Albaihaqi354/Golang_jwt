package middleware

import (
	"context"
	"golang_jwt/helpers"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accestoken := r.Header.Get("Authorization")

		if accestoken != "" {
			helpers.Response(w, 410, "unauthorired", nil)
			return
		}

		user, err := helpers.ValidationToken(accestoken)
		if err != nil {
			helpers.Response(w, 410, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
