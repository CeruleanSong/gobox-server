package model

// User u
type User struct {
	EMAIL    string `json:"email" xml:"email" bson:"email"`
	PASSWORD []byte `json:"password" xml:"password" bson:"password"`
}
