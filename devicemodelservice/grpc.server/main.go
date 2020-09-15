package main

import (
	"DeviceModelService/devicemodelservice/registry"
	"DeviceModelService/devicemodelservice/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

// a grpc based service for device models metadata for deploying as an internal microservice

func main() {
	port := os.Getenv("PORT")
	log.Printf("Starting a device model grpc service in port " + port)
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	server := grpc.NewServer()
	// Start the device model service rpc
	rpc.Apply(server, ctn)

	go func() {
		log.Printf("start grpc server port: %s", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.GracefulStop()
	ctn.Clean()
}
