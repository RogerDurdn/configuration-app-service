package lib

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib/mongoDb"
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
)

type MongoData struct {

}

func (md MongoData) FetchGroupById(id int) model.Group  {
	mongoDb.Find("")
	return model.Group{Id: 123, Features: []model.Feature{
		{Key: "text", Value: "world"},
		{Key: "bool", Value: false},
	},
	}
}

func (md MongoData)FetchGroupFeatureByKey(id int, key string) model.Feature{
	return model.Feature{GroupId: id, Key: key}
}
func (md MongoData) CreateOrUpdateGroup(group model.Group) model.Group{
	return group
}

func (md MongoData) CreateOrUpdateFeature(feature model.Feature) model.Feature{
	return feature
}
func (md MongoData) DeleteGroupById(id int){

}
func (md MongoData) DeleteFeatureByKey(groupId int, key string){

}
