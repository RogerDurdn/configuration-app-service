package lib

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
)

type MockData struct {
}

func (md MockData) FetchGroupById(id int) model.Group  {
	return model.Group{Id: 123, Features: []model.Feature{
		{Key: "text", Value: "world"},
		{Key: "bool", Value: false},
		},
	}
}

func (md MockData)FetchGroupFeatureByKey(id int, key string) model.Feature{
	return model.Feature{GroupId: id, Key: key}
}
