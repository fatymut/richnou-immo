package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	reseau "richnou-immo/agence/proto"
)

func ConnecterReseau(adresse string) (reseau.ReseauServiceClient, *grpc.ClientConn, error) {
    connexion, err := grpc.NewClient(
        adresse,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )

    if err != nil {
        return nil, nil, err
    }

    client := reseau.NewReseauServiceClient(connexion)

    return client, connexion, nil
}

func GetEtablissements(client reseau.ReseauServiceClient, codePostal string) (*reseau.EtablissementsResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    demande := &reseau.CodePostalRequest{
        CodePostal: codePostal,
    }

    reponse, err := client.GetEtablissements(ctx, demande)

    if err != nil {
        return nil, err
    }

    return reponse, nil
}