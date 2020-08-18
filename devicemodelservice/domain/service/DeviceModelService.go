package service

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"DeviceModelService/devicemodelservice/domain/repository"
)

type DeviceModelService struct {
	repo repository.DeviceModelRepository
}

//Constructor for DeviceModelService
func NewDeviceModelService(repo repository.DeviceModelRepository) *DeviceModelService {
	//return &DeviceModelService{repo: repo}
	deviceModelService := new(DeviceModelService)
	deviceModelService.repo = repo
	return deviceModelService
}

// Add a new model
func (d *DeviceModelService) AddModel(m model.DeviceModel) error {
	//modelExists, errModelExists := d.repo.FindByModelName(m.ModelName())
	//// if model does not exist create the model and save
	//if modelExists != nil {
	//	return fmt.Errorf("%s already exists", m.ModelName())
	//}
	//if(errModelExists != nil){
	//	return errModelExists
	//}
	//
	//// now save the new model
	//errSave := d.repo.Save(&m)
	//
	//if(errSave != nil){
	//	return fmt.Errorf("Error saving %s ", m.ModelName())
	//}
	return nil
}

func (d *DeviceModelService) GetModelByName(modelName string) (*model.DeviceModel, error) {
	model, err := d.repo.FindByModelName(modelName)

	if err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DeviceModelService) GetAllModels() ([]*model.DeviceModel, error) {
	models, err := d.repo.FindAll()

	if err != nil {
		return nil, err
	}
	return models, nil
}
