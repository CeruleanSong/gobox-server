package model

// ErrorResponce s
type ErrorResponce struct {
	MESSAGE string `json:"message" xml:"message" bson:"message"`
	STATUS  uint   `json:"status" xml:"status" bson:"status"`
}
