
package routes

import (
    "net/http"
    "richnou-immo/agence/handler"
)

func RegisterPublicRoutes(mux *http.ServeMux) {

    mux.HandleFunc("GET /health", handler.HealthCheck)

    mux.HandleFunc("GET /annonces", handler.GetAnnonces)

    mux.HandleFunc("GET /annonces/{id}", handler.GetAnnonceByID)

    mux.HandleFunc("POST /login", handler.Login)
}
