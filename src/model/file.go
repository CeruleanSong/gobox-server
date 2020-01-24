package model

// FileData describes a file's metadata in the database
type FileData struct {
	ID        string `json:"_id" xml:"_id" bson:"_id"`
	NAME      string `json:"name" xml:"name" bson:"name"`
	TYPE      string `json:"type" xml:"type" bson:"type"`
	DOWNLOADS int    `json:"downloads" xml:"downloads" bson:"downloads"`
	VIEWS     int    `json:"views" xml:"views" bson:"views"`
	BYTES     int64  `json:"bytes" xml:"bytes"`
	UPLOADED  string `json:"uploaded" xml:"uploaded"`
	EXPIRES   string `json:"expires" xml:"expires"`
}

// FileResponce reponce form the server when performing a file action
type FileResponce struct {
	ID        string `json:"_id" xml:"_id" bson:"_id"`
	NAME      string `json:"name" xml:"name" bson:"name"`
	URL       string `json:"url" xml:"url"`
	TYPE      string `json:"type" xml:"type" bson:"type"`
	DOWNLOADS int    `json:"downloads" xml:"downloads" bson:"downloads"`
	VIEWS     int    `json:"views" xml:"views" bson:"views"`
	BYTES     int64  `json:"bytes" xml:"bytes"`
	UPLOADED  string `json:"uploaded" xml:"uploaded"`
	EXPIRES   string `json:"expires" xml:"expires"`
}
