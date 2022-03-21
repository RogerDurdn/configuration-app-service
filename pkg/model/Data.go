package model

type Group struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Features []Feature `json:"features"`
}

type Feature struct {
	GroupId int         `json:"groupId" binding:"required"`
	Key     string      `json:"key" binding:"required"`
	Value   interface{} `json:"value" binding:"required"`
}
