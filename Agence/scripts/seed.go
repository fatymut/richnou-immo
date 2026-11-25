package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
func main() {
    uri := os.Getenv("MONGO_URI")

    if uri == "" {
        uri = "mongodb://localhost:27018"
    }

    client, err := mongo.Connect(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    defer client.Disconnect(context.Background())

    err = InsererAnnonces(client)
    if err != nil {
        fmt.Println("Erreur lors de l'insertion des annonces :", err)
        return
    }

    fmt.Println("Annonces insérées avec succès !")
    
}
func GenererAnnonces() []models.Annonce {
	villes := []string{"Toulouse", "Blagnac", "Colomiers", "Muret", "Tournefeuille"}
	codesPostaux := []string{"31000", "31700", "31770", "31600", "31170"}

	annonces := []models.Annonce{}

	for i := 0; i < 20; i++ {
		indexVille := i % len(villes)

		annonce := models.Annonce{
			ID:          i + 1,
			Titre:       fmt.Sprintf("Appartement T%d à %s", (i%4)+1, villes[indexVille]),
			Description: "Appartement à vendre en Haute-Garonne",
			Ville:       villes[indexVille],
			CodePostal:  codesPostaux[indexVille],
			Surface:     float64(35 + i*5),
			Prix:        float64(120000 + i*15000),
			Type:        "appartement",
		}

		annonces = append(annonces, annonce)
	}

	return annonces
}

func InsererAnnonces(client *mongo.Client) error {

	ctx := context.Background()
	annonces := GenererAnnonces()

	collection := client.Database("richnou_agence").Collection("annonces")

	documents := make([]interface{}, len(annonces))

	for i, annonce := range annonces {
		documents[i] = annonce
	}

	_, err := collection.InsertMany(ctx, documents)

	if err != nil {
		return err
	}

	return nil
}
