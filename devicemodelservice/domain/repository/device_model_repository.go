package repository

import (
	"DeviceModelService/devicemodelservice/domain/model"
)

//Repository interface
type DeviceModelRepository interface {
	FindAll() ([]*model.DeviceModel, error)
	FindByModelName(modeName string) (*model.DeviceModel, error)
	Save(*model.DeviceModel) error
}
