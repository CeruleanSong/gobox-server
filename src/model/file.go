package model

// FileData u
type FileData struct {
	NAME string `json:"name" xml:"name" bson:"name"`
	ID   string `json:"_id" xml:"_id" bson:"_id"`
	TYPE string `json:"type" xml:"type" bson:"type"`
}

// FileResponce s
type FileResponce struct {
	NAME  string `json:"name" xml:"name" bson:"name"`
	ID    string `json:"_id" xml:"_id" bson:"_id"`
	URL   string `json:"url" xml:"url"`
	TYPE  string `json:"type" xml:"type" bson:"type"`
	BYTES int64  `json:"bytes" xml:"bytes"`
}
