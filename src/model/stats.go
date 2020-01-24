package model

// CollectionStat describes a files server stats
type CollectionStat struct {
	NAME  string `json:"name" xml:"name" bson:"name"`
	ID    string `json:"_id" xml:"_id" bson:"_id"`
	TYPE  string `json:"type" xml:"type" bson:"type"`
	BYTES int64  `json:"bytes" xml:"bytes"`
}
