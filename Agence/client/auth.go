package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)
var tokenSession string

func RecupererToken(token string) {
    fmt.Println("Token reçu")
	tokenSession = token
}

func Login() {
    var username string
    var password string

    fmt.Print("Identifiant : ")
    fmt.Scanln(&username)

    fmt.Print("Mot de passe : ")
    fmt.Scanln(&password)

   identifiants := map[string]string{
    "username": username,
    "password": password,
	
}
	data, err := json.Marshal(identifiants)

	if err != nil {
    	fmt.Println("Erreur lors de la conversion en JSON :", err)
    	return
	}

	

	resp, err := http.Post(
    	"http://localhost:8080/login",
    	"application/json",
    	bytes.NewReader(data),
	)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
    fmt.Println("Échec de connexion. Statut HTTP :", resp.StatusCode)
    return
}
	var resultat struct {
    	Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&resultat)

	if err != nil {
    	fmt.Println("Erreur lors du décodage :", err)
    	return
	}

	if resultat.Token == "" {
    	fmt.Println("Aucun JWT reçu")
    	return
	}

	RecupererToken(resultat.Token)
	fmt.Println("Connexion réussie. JWT")
}