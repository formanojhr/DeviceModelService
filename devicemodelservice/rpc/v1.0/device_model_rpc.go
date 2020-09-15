package v1

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"DeviceModelService/devicemodelservice/domain/service"
	devicemodel2 "DeviceModelService/devicemodelservice/rpc/v1.0/protocol"
	"context"
)

type DeviceModelRPC struct {
	deviceModelService service.DeviceModelService
}

// constructor for rpc instance
func NewDeviceModelRPC(dmService service.DeviceModelService) *DeviceModelRPC {
	return &DeviceModelRPC{
		deviceModelService: dmService,
	}
}

func (s *DeviceModelRPC) FindDeviceModelByName(ctx context.Context, name *devicemodel2.ModelName) (*devicemodel2.DeviceModelResponseType, error) {
	model, err := s.deviceModelService.GetModelByName(name.ModelName)
	if err != nil {
		return nil, err
	}

	return toDevModelResType(model), nil
}

func (s *DeviceModelRPC) FindAllModels(ctx context.Context, in *devicemodel2.FindAllModelsParams) (*devicemodel2.ListDeviceModelResponseType, error) {
	models, err := s.deviceModelService.GetAllModels()

	if err != nil {
		return nil, err
	}
	i := 0
	// iterate the model domain type to rpc response type
	res := make([]*devicemodel2.DeviceModelResponseType, len(models))
	for _, deviceModel := range models {
		res[i] = toDevModelResType(deviceModel)
		i++
	}

	resList := devicemodel2.ListDeviceModelResponseType{Models: res}

	return &resList, nil
}

//convert from domain model type to rpc protocol serialization type
func toDevModelResType(model *model.DeviceModel) *devicemodel2.DeviceModelResponseType {
	mname, ok := devicemodel2.DeviceType_value[model.DeviceType]
	//model.DeviceType(mname)
	if !ok {
		panic("invalid enum value")
	}
	res := devicemodel2.DeviceModelResponseType{
		ModelName:      model.ModelName,
		ModelNumber:    model.ModelNumber,
		DeviceType:     devicemodel2.DeviceType(mname),
		DeviceTypeUUID: model.DeviceTypeUUID,
		Vendor:         model.Vendor,
		TypePatterns:   model.TypePatterns,
	}
	return &res
}
