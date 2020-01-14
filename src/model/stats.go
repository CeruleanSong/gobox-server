package model

// CollectionStat s
type CollectionStat struct {
	NAME   string `json:"name" xml:"name" bson:"name"`
	LENGTH string `json:"_id" xml:"_id" bson:"_id"`
	SIZE   string `json:"url" xml:"url"`
	TYPE   string `json:"type" xml:"type" bson:"type"`
	BYTES  int64  `json:"bytes" xml:"bytes"`
}
