package model

// File u
type File struct {
	NAME  string `json:"name" xml:"name" bson:"name"`
	TOKEN string `json:"_id" xml:"_id" bson:"_id"`
	URL   string `json:"url" xml:"url"`
	FILE  []byte `json:"file" xml:"file" bson:"file"`
}
