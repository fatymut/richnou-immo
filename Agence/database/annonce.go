package database

import (
	"context"
	"fmt"
	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)





func GetAllAnnonces(ctx context.Context, filter bson.M, limit int64, skip int64, sortField string, sortOrder int64) ([]models.Annonce, error) {

	

	annonces := []models.Annonce{}

	opts:= options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.D{{Key : sortField, Value : sortOrder}})

	cursor, err := Database.Collection("annonces").Find(ctx, filter,opts)
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

func GetAnnonceByID(ctx context.Context, id int) (models.Annonce, error) {
	var annonce models.Annonce

	err := Database.Collection("annonces").FindOne(ctx, bson.M{"id": id}).Decode(&annonce)
	if err != nil {
		return models.Annonce{}, fmt.Errorf("Annonce introuvable")
	}

	return annonce, nil
}

func ExportAnnonces(ctx context.Context) ([]models.Annonce, error) {
    // Récupérer toutes les annonces, triées par ID
	annonces := []models.Annonce{}

	opt := options.Find().SetSort(bson.D{{Key : "id", Value : 1}})

	cursor, err := Database.Collection("annonces").Find(ctx, bson.M{},opt)
		if err != nil {
    	return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &annonces); err != nil {
    	return nil, err
	}

return annonces, nil

}


func ImportAnnonces(ctx context.Context, annonces []models.Annonce) error {

    if len(annonces) == 0 {
        return fmt.Errorf("le slice d'annonces est vide")
    }

    _, err := Database.Collection("annonces").InsertMany(ctx, annonces)

    if err != nil {
        return fmt.Errorf("erreur lors de l'import : %w", err)
    }

    return nil
}
