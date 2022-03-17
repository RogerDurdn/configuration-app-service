package lib

import "github.com/RogerDurdn/ConfigurationService/pkg/model"

type DataSource interface {
	FetchDcData() model.Data
}

