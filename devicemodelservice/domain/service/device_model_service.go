package service

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"DeviceModelService/devicemodelservice/domain/repository"
	"fmt"
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
func (d *DeviceModelService) AddModel(m model.DeviceModel) (error, string) {
	modelExists, errModelExists := d.repo.FindByModelName(m.ModelName)
	// if model does not exist create the model and save
	if modelExists != nil {
		return fmt.Errorf("%s already exists", m.ModelName), m.DeviceTypeUUID
	}
	if errModelExists != nil {
		return errModelExists, ""
	}

	// now save the new model
	errSave, model := d.repo.Save(&m)

	if errSave != nil {
		return fmt.Errorf("Error saving %s ", m.ModelName), ""
	}
	return nil, model.DeviceTypeUUID
}

func (d *DeviceModelService) UpdateModel(m model.DeviceModel) error {
	model, err := d.GetModelByName(m.ModelName)
	if err != nil {
		return err
	} else if model == nil { // no pre-existing model
		fmt.Errorf("Error updating %s not pre existing model with this name exists! \n", m.ModelName)
		return fmt.Errorf("Error updating %s not pre existing model with this name exists! ", m.ModelName)
	}
	saveErr := d.repo.UpdateModel(m, *model)
	if saveErr != nil {
		return saveErr
	}
	return nil
}

func (d *DeviceModelService) GetModelByName(modelName string) (*model.DeviceModel, error) {
	if len(modelName) == 0 {
		fmt.Errorf("empty model name string received")
		return nil, fmt.Errorf("empty model name string received")
	}
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
