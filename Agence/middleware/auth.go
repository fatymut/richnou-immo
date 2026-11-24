
package middleware

import (
    "net/http"
    "strings"
	"github.com/golang-jwt/jwt/v5"
    "context"
)
type contextKey string

const roleKey contextKey = "role"

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // 1. Récupérer le header Authorization
        authHeader := r.Header.Get("Authorization")

        // 2. Vérifier la présence du header
        if authHeader == "" {
            http.Error(w, "Token manquant", http.StatusUnauthorized)
            return
        }

        // 3. Vérifier le format Bearer
        parts := strings.SplitN(authHeader, " ", 2)

        if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
            http.Error(w, "Format du token invalide", http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]

// 4. Vérifier le JWT
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("laclesecretebtw"), nil
}, jwt.WithValidMethods([]string{"HS256"}), jwt.WithExpirationRequired())

if err != nil || !token.Valid {
    http.Error(w, "Token invalide ou expiré", http.StatusUnauthorized)
    return
}

// 5. Récupérer les informations du JWT
claims, ok := token.Claims.(jwt.MapClaims)

if !ok {
    http.Error(w, "Claims invalides", http.StatusUnauthorized)
    return
}

// Récupérer le rôle de l'utilisateur
role, ok := claims["role"].(string)

if !ok || (role != "admin" && role != "vendeur") {
    http.Error(w, "Rôle invalide", http.StatusUnauthorized)
    return
}

 ctx := context.WithValue(r.Context(), roleKey, role)
next.ServeHTTP(w, r.WithContext(ctx))

})
}
