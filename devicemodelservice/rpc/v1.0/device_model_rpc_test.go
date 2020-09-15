package v1

import (
	"DeviceModelService/devicemodelservice/domain/model"
	"DeviceModelService/devicemodelservice/domain/service"
	"DeviceModelService/devicemodelservice/persistence/db"
	devicemodel2 "DeviceModelService/devicemodelservice/rpc/v1.0/protocol"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"reflect"
	"testing"
)

func devicemodelserver() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	mservice, err := buildDeviceModelService()
	if err != nil {
		//
	}
	rpc := NewDeviceModelRPC(mservice)
	devicemodel2.RegisterDeviceModelServiceServer(server, rpc)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func buildDeviceModelService() (service.DeviceModelService, error) {
	//repo := inmemory.NewDeviceModelRepository()
	repo := db.New()
	service := service.NewDeviceModelService(repo)
	return *service, nil
}

func TestDeviceModelRPC_FindAllModels(t *testing.T) {
	type fields struct {
		deviceModelService service.DeviceModelService
	}
	type args struct {
		ctx context.Context
		in  *devicemodel2.FindAllModelsParams
	}
	//m := make([]devicemodel2.ListDeviceModelResponseType, 3)
	p := make([]*devicemodel2.ListDeviceModelResponseType, 0)
	in := new(devicemodel2.FindAllModelsParams)
	tests := []struct {
		name    string
		want    []*devicemodel2.ListDeviceModelResponseType
		wantErr string
	}{
		{"",
			p,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(devicemodelserver()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//create a client connection
	client := devicemodel2.NewDeviceModelServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &DeviceModelRPC{
			//	deviceModelService: tt.fields.deviceModelService,
			//}
			got, err := client.FindAllModels(ctx, in)
			if err != nil {
				t.Errorf("FindAllModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(len(got.Models), tt.want) {
			//	t.Errorf("FindAllModels() got = %v, want %v", got, tt.want)
			//}
			if len(got.Models) == 0 {
				t.Errorf("FindAllModels() got = %v, want %s", len(got.Models), "> 0")
			}
		})
	}
}

func TestDeviceModelRPC_FindDeviceModelByName(t *testing.T) {
	type fields struct {
		deviceModelService service.DeviceModelService
	}
	type args struct {
		ctx  context.Context
		name *devicemodel2.ModelName
	}
	mname, ok := devicemodel2.DeviceType_value["Trio"]
	if ok == false {
		t.Errorf("Could not convert model = %s, to DeviceType valie", "Trio")
	}
	//client, ctx := Create_client()
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(devicemodelserver()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//create a client connection
	client := devicemodel2.NewDeviceModelServiceClient(conn)
	tests := []struct {
		name    string
		want    *devicemodel2.DeviceModelResponseType
		wantErr bool
	}{
		{
			"Trio",
			&devicemodel2.DeviceModelResponseType{
				ModelName:      "Trio",
				ModelNumber:    "8300",
				DeviceType:     devicemodel2.DeviceType(mname),
				DeviceTypeUUID: "f957617b-e773-4864-95f1-df82a6fe5e09",
				Vendor:         "Polycom",
				TypePatterns:   []string{"Trio"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.FindDeviceModelByName(ctx, &devicemodel2.ModelName{ModelName: "Trio"})
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDeviceModelByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FindDeviceModelByName() got = %v, want %v", got, tt.want)
			//}
			if got.ModelName != tt.want.ModelName || got.DeviceType != tt.want.DeviceType || got.DeviceTypeUUID != tt.want.DeviceTypeUUID {
				t.Errorf("FindDeviceModelByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeviceModelRPC(t *testing.T) {
	type args struct {
		dmService service.DeviceModelService
	}
	tests := []struct {
		name string
		args args
		want *DeviceModelRPC
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeviceModelRPC(tt.args.dmService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeviceModelRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDevModelResType(t *testing.T) {
	type args struct {
		model *model.DeviceModel
	}
	tests := []struct {
		name string
		args args
		want *devicemodel2.DeviceModelResponseType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDevModelResType(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDevModelResType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Create_client() (devicemodel2.DeviceModelServiceClient, context.Context) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(devicemodelserver()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//create a client connection
	client := devicemodel2.NewDeviceModelServiceClient(conn)
	defer conn.Close()
	return client, ctx
}
