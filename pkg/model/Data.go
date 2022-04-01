package model

type Group struct {
	Id       int       `json:"id" binding:"required" bson:"_id, omitempty"`
	Name     string    `json:"name" binding:"required" bson:"name, omitempty"`
	Features []Feature `json:"features" bson:"features, omitempty"`
}

type Feature struct {
	Key   string      `json:"key" binding:"required" bson:"key, omitempty"`
	Value interface{} `json:"value" binding:"required" bson:"value, omitempty"`
}
