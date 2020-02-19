package model

import (
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
)

// User describes a user in the database
type User struct {
	USER     string         `json:"_id" xml:"_id" bson:"_id"`
	PASSWORD []byte         `json:"password" xml:"password" bson:"password"`
	KEY      rsa.PrivateKey `json:"key" xml:"key" bson:"key"`
}

// Token describes a user in the database
type Token struct {
	USER  string `json:"user" xml:"user" bson:"user"`
	ADMIN bool   `json:"admin" xml:"admin" bson:"admin"`
	jwt.StandardClaims
}
