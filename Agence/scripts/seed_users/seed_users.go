package main
import (
	"context"
	"fmt"
	"log"
	"richnou-immo/agence/database"
	"richnou-immo/agence/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.NewDB()

	ctx := context.Background()
	collection := database.Database.Collection("users")

	users := []struct {
		ID       int
		Username string
		Password string
		Role     string
	}{
		{1, "admin", "Admin123!", "admin"},
		{2, "vendeur", "Vendeur123!", "vendeur"},
	}

	for _, u := range users {
		count, err := collection.CountDocuments(ctx, bson.M{
			"username": u.Username,
		})
		if err != nil {
			log.Fatal(err)
		}

		if count > 0 {
			fmt.Println("Utilisateur déjà existant :", u.Username)
			continue
		}

		// Hacher le mot de passe
		hash, err := bcrypt.GenerateFromPassword(
			[]byte(u.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Préparer l'utilisateur
		user := models.User{
			ID:           u.ID,
			Username:     u.Username,
			PasswordHash: string(hash),
			Role:         u.Role,
		}

		// Insérer dans MongoDB
		_, err = collection.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("user cree :", u.Username)
	}

	fmt.Println("Seed utilisateurs termine ")
}
