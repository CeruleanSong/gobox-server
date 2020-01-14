package model

import "time"

// FileData u
type FileData struct {
	ID       string    `json:"_id" xml:"_id" bson:"_id"`
	NAME     string    `json:"name" xml:"name" bson:"name"`
	TYPE     string    `json:"type" xml:"type" bson:"type"`
	BYTES    int64     `json:"bytes" xml:"bytes"`
	UPLOADED time.Time `json:"uploaded" xml:"uploaded"`
	EXPIRES  time.Time `json:"expires" xml:"expires"`
}

// FileResponce s
type FileResponce struct {
	NAME     string    `json:"name" xml:"name" bson:"name"`
	ID       string    `json:"_id" xml:"_id" bson:"_id"`
	URL      string    `json:"url" xml:"url"`
	TYPE     string    `json:"type" xml:"type" bson:"type"`
	BYTES    int64     `json:"bytes" xml:"bytes"`
	UPLOADED time.Time `json:"uploaded" xml:"uploaded"`
	EXPIRES  time.Time `json:"expires" xml:"expires"`
}
