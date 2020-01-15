package model

import "time"

// FileData u
type FileData struct {
	ID        string    `json:"_id" xml:"_id" bson:"_id"`
	NAME      string    `json:"name" xml:"name" bson:"name"`
	TYPE      string    `json:"type" xml:"type" bson:"type"`
	DOWNLOADS int       `json:"downloads" xml:"downloads" bson:"downloads"`
	VIEWS     int       `json:"views" xml:"views" bson:"views"`
	BYTES     int64     `json:"bytes" xml:"bytes"`
	UPLOADED  time.Time `json:"uploaded" xml:"uploaded"`
	EXPIRES   time.Time `json:"expires" xml:"expires"`
}

// FileResponce s
type FileResponce struct {
	ID        string    `json:"_id" xml:"_id" bson:"_id"`
	NAME      string    `json:"name" xml:"name" bson:"name"`
	URL       string    `json:"url" xml:"url"`
	TYPE      string    `json:"type" xml:"type" bson:"type"`
	DOWNLOADS string    `json:"downloads" xml:"downloads" bson:"downloads"`
	VIEWS     string    `json:"views" xml:"views" bson:"views"`
	BYTES     int64     `json:"bytes" xml:"bytes"`
	UPLOADED  time.Time `json:"uploaded" xml:"uploaded"`
	EXPIRES   time.Time `json:"expires" xml:"expires"`
}
