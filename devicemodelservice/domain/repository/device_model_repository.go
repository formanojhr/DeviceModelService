package repository

import (
	"DeviceModelService/devicemodelservice/domain/model"
)

//Repository interface
type DeviceModelRepository interface {
	FindAll() ([]*model.DeviceModel, error)
	FindByModelName(modelName string) (*model.DeviceModel, error)
	Save(model *model.DeviceModel) (error, model.DeviceModel) // TODO remove pointers pass be reference for struct; Makes it more confusin
	UpdateModel(modelUpdate model.DeviceModel, oldModel model.DeviceModel) error
	DeleteItem(modelName string) (error, string)
}
