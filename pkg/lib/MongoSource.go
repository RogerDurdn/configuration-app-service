package lib

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/lib/mongoDb"
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
)

type MongoData struct{}

func (md MongoData) FetchGroupById(id int) (*model.Group, error) {
	group := &model.Group{}
	if err := mongoDb.FindGroupById(id, group); err != nil {
		return nil, err
	}
	return group, nil
}

func (md MongoData) CreateOrUpdateGroup(group model.Group) error {
	return mongoDb.CreateGroup(group)
}

func (md MongoData) DeleteGroupById(id int) error {
	return mongoDb.DeleteGroupById(id)

}
