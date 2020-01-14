package model

// ErrorResponce s
type ErrorResponce struct {
	MESSAGE string `json:"message" xml:"message" bson:"message"`
	STATUS  string `json:"status" xml:"status" bson:"status"`
}
