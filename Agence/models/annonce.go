package models

type Annonce struct {
    ID          int     `json:"id" bson:"id"`
    Titre       string  `json:"titre" bson:"titre"`
    Description string  `json:"description" bson:"description"`
    Ville       string  `json:"ville" bson:"ville"`
    CodePostal  string  `json:"code_postal" bson:"code_postal"`
    Surface     float64 `json:"surface" bson:"surface"`
    Prix        float64 `json:"prix" bson:"prix"`
    Type        string  `json:"type" bson:"type"`
}
