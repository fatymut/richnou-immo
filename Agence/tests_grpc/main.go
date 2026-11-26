
package main

import (
    "context"
    "log"
    "net"

    reseau "richnou-immo/agence/proto"

    "google.golang.org/grpc"
)

// Notre serveur de test
type serveur struct {
    reseau.UnimplementedReseauServiceServer
}

// Méthode appelée par le client Agence
func (s *serveur) GetEtablissements(
    ctx context.Context,
    demande *reseau.CodePostalRequest,
) (*reseau.EtablissementsResponse, error) {

    log.Println("Code postal reçu :", demande.CodePostal)

    return &reseau.EtablissementsResponse{
        Etablissements: []*reseau.Etablissement{
            {
                Id:         "1",
                Titre:      "École de test",
                Categorie:  "école",
                CodePostal: demande.CodePostal,
            },
            {
                Id:         "2",
                Titre:      "Gare de test",
                Categorie:  "transport",
                CodePostal: demande.CodePostal,
            },
        },
    }, nil
}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatal(err)
    }

    serveurGRPC := grpc.NewServer()

    reseau.RegisterReseauServiceServer(serveurGRPC, &serveur{})

    log.Println("Serveur gRPC de test démarré sur le port 50051")

    if err := serveurGRPC.Serve(listener); err != nil {
        log.Fatal(err)
    }
}
