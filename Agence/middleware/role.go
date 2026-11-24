
package middleware

import "net/http"

func RequireRole(requiredRole string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            role, ok := r.Context().Value(roleKey).(string)

            if !ok {
                http.Error(w, "Authentification requise", http.StatusUnauthorized)
                return
            }

            if role != requiredRole {
                http.Error(w, "Accès interdit", http.StatusForbidden)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
