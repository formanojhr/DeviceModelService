package inmemory

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"sync"
)

type DeviceModelRepository struct {
	mu           *sync.Mutex
	devicemodels map[string]model.DeviceModel
}

var models = map[string]model.DeviceModel{
	"f957617b-e773-4864-95f1-df82a6fe5e09": model.DeviceModel{ModelName: "Trio", ModelNumber: "8300", DeviceType: "IPPHONE", DeviceTypeUUID: "f957617b-e773-4864-95f1-df82a6fe5e09", Vendor: "Polycom", TypePatterns: []string{"8300"}},
	"89c4e433-34ac-4063-8c03-e32d08e1b748": model.DeviceModel{ModelName: "IP Phone", ModelNumber: "7800", DeviceType: "IPPHONE", DeviceTypeUUID: "89c4e433-34ac-4063-8c03-e32d08e1b748", Vendor: "CISCO", TypePatterns: []string{"7800"}},
}

//constructor
func NewDeviceModelRepository() *DeviceModelRepository {
	var d DeviceModelRepository
	//var models = map[string]model.DeviceModel{
	//	"f957617b-e773-4864-95f1-df82a6fe5e09": model.DeviceModel{ModelName: "Trio", ModelNumber: "8300", DeviceType: "IPPHONE", DeviceTypeUUID: "f957617b-e773-4864-95f1-df82a6fe5e09", Vendor: "Polycom", TypePatterns:[]string{"8300",}},
	//	"89c4e433-34ac-4063-8c03-e32d08e1b748": model.DeviceModel{ModelName: "IP Phone", ModelNumber: "7800", DeviceType: "IPPHONE", DeviceTypeUUID:"89c4e433-34ac-4063-8c03-e32d08e1b748", Vendor:"CISCO", TypePatterns:[]string{"7800",}},
	//}

	// initialize the models to standard basic ones
	d.devicemodels = map[string]model.DeviceModel{
		"f957617b-e773-4864-95f1-df82a6fe5e09": model.DeviceModel{ModelName: "Trio", ModelNumber: "8300", DeviceType: "IPPHONE", DeviceTypeUUID: "f957617b-e773-4864-95f1-df82a6fe5e09", Vendor: "Polycom", TypePatterns: []string{"8300"}},
		"89c4e433-34ac-4063-8c03-e32d08e1b748": model.DeviceModel{ModelName: "IP Phone", ModelNumber: "7800", DeviceType: "IPPHONE", DeviceTypeUUID: "89c4e433-34ac-4063-8c03-e32d08e1b748", Vendor: "CISCO", TypePatterns: []string{"7800"}},
	}
	d.mu = &sync.Mutex{}
	return &d
}

//get all the models
func (r *DeviceModelRepository) FindAll() ([]*model.DeviceModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	deviceModels := make([]*model.DeviceModel, len(r.devicemodels))
	i := 0
	for _, deviceModel := range r.devicemodels {
		deviceModels[i] = model.NewDeviceModel(deviceModel.ModelName, deviceModel.ModelNumber, deviceModel.DeviceType,
			deviceModel.DeviceTypeUUID, deviceModel.Vendor)
		i++
	}
	return deviceModels, nil
}

func (r *DeviceModelRepository) FindByModelName(name string) (*model.DeviceModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, deviceModel := range r.devicemodels {
		if deviceModel.ModelName == name {
			return &deviceModel, nil
		}
	}
	return nil, nil
}

func (r *DeviceModelRepository) Save(modelIn *model.DeviceModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.devicemodels[modelIn.DeviceTypeUUID] = model.DeviceModel{
		ModelName:   modelIn.ModelName,
		ModelNumber: modelIn.ModelNumber, DeviceType: modelIn.DeviceType,
		DeviceTypeUUID: modelIn.DeviceTypeUUID,
		Vendor:         modelIn.Vendor, TypePatterns: modelIn.TypePatterns,
	}
	return nil
}
