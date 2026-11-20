package main

import (
	"context"
	"log"
	"net/http"

	"richnou-immo/agence/database"
	"richnou-immo/agence/routes"
	// "richnou-immo/agence/scripts"
)

func main() {
	database.NewDB()
	defer database.Db.Disconnect(context.Background())

	http.HandleFunc("GET /health", routes.HealthCheck)
	http.HandleFunc("GET /annonces", routes.GetAnnonces)
	http.HandleFunc("GET /annonces/{id}", routes.GetAnnonceByID)

	log.Println("Serveur démarré sur le port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
