package registry

import (
	"DeviceModelService/devicemodelservice/api"
	service2 "DeviceModelService/devicemodelservice/domain/service"
	//"DeviceModelService/devicemodelservice/persistence/inmemory"
	"DeviceModelService/devicemodelservice/persistence/db"
	"github.com/sarulabs/di"
)

type DeviceModelContainer struct {
	ctn di.Container
}

//Constructor
func NewContainer() (*DeviceModelContainer, error) {
	builder, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "device-model-API",
			Build: buildDeviceModelAPI,
		},
		{
			Name:  "device-model-service",
			Build: buildDeviceModelService,
		},
	}...); err != nil {
		return nil, err
	}

	return &DeviceModelContainer{
		ctn: builder.Build(),
	}, nil
}

func (c *DeviceModelContainer) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *DeviceModelContainer) Clean() error {
	return c.ctn.Clean()
}

func buildDeviceModelAPI(ctn di.Container) (interface{}, error) {
	//repo := inmemory.NewDeviceModelRepository()
	repo := db.New()
	service := service2.NewDeviceModelService(repo)
	return api.NewDeviceModelAPI(service), nil
}

func buildDeviceModelService(ctn di.Container) (interface{}, error) {
	//repo := inmemory.NewDeviceModelRepository()
	repo := db.New()
	service := service2.NewDeviceModelService(repo)
	return *service, nil
}
