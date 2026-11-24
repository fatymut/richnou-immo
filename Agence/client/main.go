package main

import (
	"fmt"
)

func main() {

	for {
		fmt.Println("\n     RICHNOU IMMO     \n")
		fmt.Println("1. Se connecter")
		fmt.Println("2. Afficher les annonces")
		fmt.Println("3. Rechercher des annonces")
		fmt.Println("4. Quitter")

		var choix int

		fmt.Print("\n Votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// Connexion
			Login()
			if tokenSession != "" {
				fmt.Println("Session active : oui")
			} else {
				fmt.Println("Session active : non")
			}

		case 2:
			// Afficher les annonces
			AfficherAnnonces("http://localhost:8080/annonces")
		case 3:
			RechercherAnnonces()

		case 4:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide")
		}
	}
}
