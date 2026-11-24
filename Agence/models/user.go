package models


type User struct {
    ID          int     `json:"id" bson:"id"`
    Username       string  `json:"username" bson:"username"`
    PasswordHash string  `json:"-" bson:"password_hash"`
    Role       string  `json:"role" bson:"role"`
    
}
