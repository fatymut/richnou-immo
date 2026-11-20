package routes

import (
	"encoding/json"
	"net/http"
	"richnou-immo/agence/database"
	"strconv"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(map[string]string{
        "status": "ok",
    })
}



func GetAnnonces(w http.ResponseWriter, r *http.Request) {
    
    ctx := r.Context()
    annonces, err := database.GetAllAnnonces(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(annonces)
}


func GetAnnonceByID(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    idStr := r.PathValue("id")
    if idStr == "" {
        http.Error(w, "Id manquant", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Id invalide", http.StatusBadRequest)
        return
    }

    annonce, err := database.GetAnnonceByID(ctx, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(annonce)
}