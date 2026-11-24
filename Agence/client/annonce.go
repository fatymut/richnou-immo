package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"richnou-immo/agence/models"
)

func AfficherAnnonces(address string) {
	

    resp, err := http.Get(address)
    if err != nil {
        fmt.Println("Erreur de récupération des annonces :", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Erreur HTTP :", resp.StatusCode)
        return
    }

    var annonces []models.Annonce

    err = json.NewDecoder(resp.Body).Decode(&annonces)
    if err != nil {
        fmt.Println("Erreur de décodage JSON :", err)
        return
    }

	if len(annonces) == 0 {
    fmt.Println("Aucune annonce trouvée.")
    return
}

    for _, annonce := range annonces {
        fmt.Printf(
            " Id : %d | Titre : %s | Description : %s | Ville : %s | Prix : %f €\n",
            annonce.ID,
            annonce.Titre,
            annonce.Description,
            annonce.Ville,
            annonce.Prix,
        )
    }
}

func RechercherAnnonces() {
    var ville string
	var prixMax string
	var surfaceMin string
	var page string
	var limit string
	var sort string

    fmt.Print("Ville recherchée : ")
    fmt.Scanln(&ville)

	fmt.Print("Prix max : ")
    fmt.Scanln(&prixMax)

	fmt.Print("Surface minimum ")
    fmt.Scanln(&surfaceMin)

	fmt.Print("page : ")
    fmt.Scanln(&page)

	fmt.Print("limit : ")
    fmt.Scanln(&limit)

	fmt.Print("sort : ")
    fmt.Scanln(&sort)

    filtres := url.Values{}

    if ville != "" {
        filtres.Set("ville", ville)
    }
	if prixMax != "" {
        filtres.Set("prix_max", prixMax)
    }
	if surfaceMin != "" {
        filtres.Set("surface_min", surfaceMin)
    }
	if page != "" {
        filtres.Set("page", page)
    }
	if limit != "" {
        filtres.Set("limit", limit)
    }
	if sort != "" {
        filtres.Set("sort", sort)
    }



    adresse := "http://localhost:8080/annonces"

    if len(filtres) > 0 {
        adresse += "?" + filtres.Encode()
    }

    fmt.Println("URL générée :", adresse)

	AfficherAnnonces(adresse)
}
