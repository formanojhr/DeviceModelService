package v1

import (
	"DeviceModelService/devicemodelservice/domain/service"
	"DeviceModelService/devicemodelservice/registry"
	"DeviceModelService/devicemodelservice/rpc/v1.0/protocol"
	"google.golang.org/grpc"
	"log"
)

// register with the grpc server the device model rpc service
func Apply(server *grpc.Server, ctn *registry.DeviceModelContainer) {
	protocol.RegisterDeviceModelServiceServer(server,
		NewDeviceModelRPC(ctn.Resolve("device-model-service").(service.DeviceModelService)))
	log.Println("Successfully registered  DeviceModelRPC!")
}
