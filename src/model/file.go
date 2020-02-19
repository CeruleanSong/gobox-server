package model

// FileData describes a file's metadata in the database
type FileData struct {
	ID        string `json:"_id" xml:"_id" bson:"_id"`
	NAME      string `json:"name" xml:"name" bson:"name"`
	TYPE      string `json:"type" xml:"type" bson:"type"`
	OWNERID   string `json:"owner_id" xml:"nowner_idame" bson:"owner_id"`
	PROTECTED bool   `json:"protected" xml:"protected" bson:"protected"`
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

// FileStat describes a files server stats
type FileStat struct {
	NAME  string `json:"name" xml:"name" bson:"name"`
	ID    string `json:"_id" xml:"_id" bson:"_id"`
	TYPE  string `json:"type" xml:"type" bson:"type"`
	BYTES int64  `json:"bytes" xml:"bytes"`
}
