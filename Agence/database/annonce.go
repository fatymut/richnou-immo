package database

import (
	"context"
	"fmt"
	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	// "go.mongodb.org/mongo-driver/v2/mongo/options"
)





func GetAllAnnonces(ctx context.Context) ([]models.Annonce, error) {

	

	annonces := []models.Annonce{}

	cursor, err := Database.Collection("annonces").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erreur recup annonces: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var annonce models.Annonce
		if err := cursor.Decode(&annonce); err != nil {
			return nil, fmt.Errorf("erreur de décodage de l'annonce: %v", err)
		}
		annonces = append(annonces, annonce)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return annonces,nil
}

func GetAnnonceByID(ctx context.Context, id int) (*models.Annonce, error) {
	var annonce models.Annonce

	err := Database.Collection("annonces").FindOne(ctx, bson.M{"id": id}).Decode(&annonce)
	if err != nil {
		return nil, fmt.Errorf("Annonce introuvable")
	}

	return &annonce, nil
}
