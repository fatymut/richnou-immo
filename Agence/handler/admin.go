package handler

import (
	"encoding/json"
	"net/http"
	"richnou-immo/agence/database"
	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ExportAnnonces(w http.ResponseWriter, r *http.Request){

    file,err := database.ExportAnnonces(r.Context())

    if err != nil{
        http.Error(w,"Erreur dans la creation du fichier json",http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Disposition", `attachment; filename="annonces.json"`)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(file)

}

func ImportAnnonces(w http.ResponseWriter, r *http.Request) {

    var annonces = []models.Annonce{}

    err:= json.NewDecoder(r.Body).Decode(&annonces)

    if err != nil{
        http.Error(w,"erreur decodage annonces", http.StatusBadRequest)
        return
    }

    if len(annonces) ==0{
        http.Error(w,"annonce vide",http.StatusBadRequest)
        return
    }

    err = database.ImportAnnonces(r.Context(),annonces)

   if err != nil {
    if mongo.IsDuplicateKeyError(err) {
        http.Error(w, "Une annonce possède déjà cet ID", http.StatusConflict)
        return
    }

    http.Error(w, "Erreur lors de l'import", http.StatusInternalServerError)
    return
}

    w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{
        "annonces_importees": len(annonces),
        })


}