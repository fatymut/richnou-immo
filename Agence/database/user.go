package database

import (
	"context"

	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
    var user models.User

    err := Database.Collection("users").
        FindOne(ctx, bson.M{"username": username}).
        Decode(&user)

    if err != nil {
        return nil, err
    }

    return &user, nil
}