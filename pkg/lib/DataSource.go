package lib

import (
	"github.com/RogerDurdn/ConfigurationService/pkg/model"
)

type DataSource interface {
	FetchGroupById(id int) (*model.Group, error)
	CreateOrUpdateGroup(group model.Group) error
	DeleteGroupById(id int) error
}
