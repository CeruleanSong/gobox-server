package model

import "crypto/rsa"

// User describes a user in the database
type User struct {
	EMAIL    string         `json:"_id" xml:"_id" bson:"_id"`
	PASSWORD []byte         `json:"password" xml:"password" bson:"password"`
	KEY      rsa.PrivateKey `json:"key" xml:"key" bson:"key"`
}
