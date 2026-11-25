
# Richnou Immo - API Agence

API REST développée en Go pour consulter des annonces immobilières en Haute-Garonne.

## Technologies

- Go
- MongoDB
- Docker et Docker Compose
- JWT pour l'authentification

## Fonctionnalités

- Liste des annonces immobilières
- Consultation d'une annonce par ID
- Recherche par ville, prix et surface
- Pagination et tri
- Authentification admin et vendeur
- Import et export JSON réservés à l'administrateur
- Client CLI développé en Go

## Démarrage avec Docker

Depuis le dossier Agence :

```bash
docker compose up --build -d
```

L'API est accessible sur :

http://localhost:8080

Vérifier son fonctionnement :

```bash
curl http://localhost:8080/health
```

## Initialisation des données

Insérer les annonces :

```bash
MONGO_URI=mongodb://localhost:27018 go run ./scripts/seed_cli
```

Insérer les utilisateurs :

Utiliser le programme du dossier `scripts/seed_users` avec
la connexion MongoDB configurée pour l'environnement Docker.

## Endpoints REST

| Méthode | Endpoint | Description |
|---|---|---|
| GET | /health | Vérifier le serveur |
| GET | /annonces | Lister et rechercher les annonces |
| GET | /annonces/{id} | Consulter une annonce |
| POST | /login | Authentification |
| GET | /admin/export | Exporter les annonces (admin) |
| POST | /admin/import | Importer des annonces (admin) |

### Paramètres de recherche

- `ville`
- `prix_max`
- `surface_min`
- `page`
- `limit`
- `sort`

Exemple :

```bash
curl "http://localhost:8080/annonces?ville=Toulouse&limit=5&sort=-prix"
```

## Client CLI

Depuis le dossier Agence :

```bash
go run ./client
```

Le client permet de se connecter, de consulter les annonces
et d'effectuer des recherches avec filtres.

## Arrêter Docker

```bash
docker compose down
```

Les données MongoDB restent conservées dans le volume Docker.

## Intégration gRPC

La communication avec le service Réseau sera documentée
après son intégration.
