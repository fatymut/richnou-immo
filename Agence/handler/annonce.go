package handler

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
	"richnou-immo/agence/database"
	"richnou-immo/agence/grpc"
	"richnou-immo/agence/models"
	reseau "richnou-immo/agence/proto"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(map[string]string{
        "status": "ok",
    })
}

type AnnonceDetailResponse struct {
    Annonce        models.Annonce         `json:"annonce"`
    Etablissements []*reseau.Etablissement `json:"etablissements"`
}

func GetAnnonces(w http.ResponseWriter, r *http.Request) {
    var filter = bson.M{}
    var limit int64 = 5
    var page int64= 1
    var err error

    var sortField = "id"
    var sortOrder int64 = 1 // 1 pour ascendant, -1 pour descendant


    if r.URL.Query().Get("sort") != "" {
        sort := r.URL.Query().Get("sort")

        if sort != "prix" && sort != "surface" && sort != "ville" && sort != "-prix" && sort != "-surface" && sort != "-ville"{
            http.Error(w, "champs de triee autorise : prix, surface, ville", http.StatusBadRequest)
			return
        }
        
        if sort[0] == '-'{
            sortOrder = -1
            sortField = sort[1:]
        }else{
            sortField = sort
        }
    }

    if r.URL.Query().Get("page") != "" {
        page, err = strconv.ParseInt(r.URL.Query().Get("page"),10,64)
        if err != nil{
            http.Error(w, "page doit être un nombre", http.StatusBadRequest)
			return
        }

        if page<1{
           http.Error(w, "page doit être > 0", http.StatusBadRequest)
			return 
        }
    }

     if r.URL.Query().Get("limit") != "" {
        limit, err = strconv.ParseInt(r.URL.Query().Get("limit"),10,64)
        if err != nil{
            http.Error(w, "limit doit être un nombre", http.StatusBadRequest)
			return
        }

        if limit<1 ||limit>10 {
            http.Error(w, "limit doit être compris entre 1 et 10", http.StatusBadRequest)
			return
        }

    }

    var skip int64 = (page -1) * limit

    if r.URL.Query().Get("prix_max") != "" {
        prixMax, err := strconv.ParseFloat(r.URL.Query().Get("prix_max"),64)

        if err != nil || math.IsNaN(prixMax) || math.IsInf(prixMax, 0){
			http.Error(w, "prix_max doit être un nombre", http.StatusBadRequest)
			return
		}
         if prixMax < 1{
            http.Error(w, "prix_max doit être superieur a 0", http.StatusBadRequest)
            return

        }

        filter["prix"] = bson.M{"$lte": prixMax}
        
    }

        if r.URL.Query().Get("surface_min") != "" {
        surfaceMin, err := strconv.ParseFloat(r.URL.Query().Get("surface_min"),64)

        if err != nil || math.IsNaN(surfaceMin) || math.IsInf(surfaceMin, 0){
			http.Error(w, "surface_min doit être un nombre", http.StatusBadRequest)
			return
		}

        if surfaceMin < 1{
            http.Error(w, "surface_min doit être superieur a 0", http.StatusBadRequest)
            return

        }

    

        filter["surface"] = bson.M{"$gte": surfaceMin}
        
    }

    if r.URL.Query().Get("ville") != "" {
        ville := r.URL.Query().Get("ville")
        filter ["ville"] = bson.M{"$regex": ville, "$options": "i"}
    }
    
    ctx := r.Context()
    annonces, err := database.GetAllAnnonces(ctx,filter, limit, skip, sortField, sortOrder)
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
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

   adresseReseau := os.Getenv("GRPC_RESEAU_ADDR")

if adresseReseau == "" {
    adresseReseau = "localhost:50051"
}

client, connexion, err := grpcclient.ConnecterReseau(adresseReseau)
    if err != nil {
        http.Error(w, "Erreur client gRPC", http.StatusBadGateway)
        return
    }
    defer connexion.Close()

    reponse, err := grpcclient.GetEtablissements(client, annonce.CodePostal)
    if err != nil {
        http.Error(w, "Service Réseau indisponible", http.StatusBadGateway)
        return
    }

    resultat := AnnonceDetailResponse{
        Annonce:        annonce,
        Etablissements: reponse.Etablissements,
    }

    // 4. Renvoyer le résultat en JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resultat)
}


