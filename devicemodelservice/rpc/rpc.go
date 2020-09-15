package rpc

import (
	"DeviceModelService/devicemodelservice/registry"
	"DeviceModelService/devicemodelservice/rpc/v1.0"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.DeviceModelContainer) {
	v1.Apply(server, ctn)
}
