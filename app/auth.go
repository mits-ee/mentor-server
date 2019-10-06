package app

import (
	"net/http"
	u "../utils"
	"strings"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response);
			return
		}

		tokenString := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		fmt.Sprint("User % authenticated successfully", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}