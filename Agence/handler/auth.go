package handler

import (
	"encoding/json"
	"net/http"
	"richnou-immo/agence/database"
	"richnou-immo/agence/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
    // 1. Décoder le JSON envoyé par le client
    var loginReq models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 2. Rechercher l'utilisateur dans MongoDB
    user, err := database.GetUserByUsername(r.Context(), loginReq.Username)
    if err != nil {
        http.Error(w, "Utilisateur inexistant", http.StatusUnauthorized)
        return
    }

    // 3. Vérifier son mot de passe avec bcrypt
    if err := bcrypt.CompareHashAndPassword(
        []byte(user.PasswordHash),
        []byte(loginReq.Password),
    ); err != nil {
        http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
        return
    }

    // 4. Générer un JWT
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{		"username": user.Username,
		"role":     user.Role,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("laclesecretebtw"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // 5. Retourner le token en JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}