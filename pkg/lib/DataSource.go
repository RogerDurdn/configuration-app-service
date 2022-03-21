package lib

import "github.com/RogerDurdn/ConfigurationService/pkg/model"

type DataSource interface {
	FetchGroupById(id int) model.Group
	FetchGroupFeatureByKey(dc int, key string) model.Feature
	CreateOrUpdateGroup(group model.Group) model.Group
	CreateOrUpdateFeature(group model.Feature) model.Feature
	DeleteGroupById(id int)
	DeleteFeatureByKey(groupId int, key string)
}

