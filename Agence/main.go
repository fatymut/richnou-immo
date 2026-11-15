
package main

import (
    "context"
    "log"
    "net/http"

    "richnou-immo/agence/database"
    "richnou-immo/agence/routes"
)

func main() {
    client, err := database.Connect()

    if err != nil {
        log.Fatal("Erreur de connexion MongoDB : ", err)
    }
    defer client.Disconnect(context.Background())

    log.Println("Connexion à MongoDB réussie !")
    http.HandleFunc("GET /health", routes.HealthCheck)

    log.Println("Serveur démarré sur le port 8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Println("Arrêt du serveur :", err)
    }
}
