package lib

import "github.com/RogerDurdn/ConfigurationService/pkg/model"

type MockData struct {
}

func (md MockData) FetchDcData() model.Data  {
	return model.Data{DcNumber: 123}
}