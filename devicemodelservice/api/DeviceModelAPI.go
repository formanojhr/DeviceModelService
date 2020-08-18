package api

import (
	"DeviceModelService/devicemodelservice/domain/model"
	//"DeviceModelService/devicemodelservice/domain/repository"
	"DeviceModelService/devicemodelservice/domain/service"
	"fmt"
)

// an use case API for querying device models by name
type DeviceModelAPI interface {
	GetModelByName(name string) (model.DeviceModel, error)
	GetModels() []model.DeviceModel
}

type DeviceModelAPIInstance struct {
	//repo    DeviceModelRepository
	Service *service.DeviceModelService
}

// constructor
func NewDeviceModelAPI(service *service.DeviceModelService) *DeviceModelAPIInstance {
	return &DeviceModelAPIInstance{
		//repo:    repo,
		Service: service,
	}
}

// get model by name
func (s *DeviceModelAPIInstance) GetModelByName(name string) (*model.DeviceModel, error) {
	model, err := s.Service.GetModelByName(name)

	if err != nil {
		return nil, err
	}

	if model == nil {
		fmt.Println("%s model not found", name)
		return nil, fmt.Errorf("%s model not found", name)
	}
	return model, nil
}

func (s *DeviceModelAPIInstance) GetModels() ([]*model.DeviceModel, error) {
	models, err := s.Service.GetAllModels()

	if err != nil {
		return nil, err
	}

	if models == nil {
		fmt.Println("No models are found")
		return nil, fmt.Errorf("No models are found")
	}
	return models, nil
}
