
package main

import (
    "log"
    "net/http"

    "richnou-immo/agence/database"
    "richnou-immo/agence/routes"
)

func main() {

    database.NewDB()
    mux := http.NewServeMux()

    routes.RegisterPublicRoutes(mux)
    routes.RegisterAdminRoutes(mux)

    log.Println("Serveur démarré sur le port 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
