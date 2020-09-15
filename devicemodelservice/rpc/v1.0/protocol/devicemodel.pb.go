// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: devicemodel.proto

package protocol

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeviceType int32

const (
	DeviceType_IPPHONE DeviceType = 0
	DeviceType_VIDEO   DeviceType = 1
	DeviceType_HEADSET DeviceType = 2
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "IPPHONE",
		1: "VIDEO",
		2: "HEADSET",
	}
	DeviceType_value = map[string]int32{
		"IPPHONE": 0,
		"VIDEO":   1,
		"HEADSET": 2,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_devicemodel_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_devicemodel_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_devicemodel_proto_rawDescGZIP(), []int{0}
}

type ModelName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName string `protobuf:"bytes,1,opt,name=modelName,proto3" json:"modelName,omitempty"`
}

func (x *ModelName) Reset() {
	*x = ModelName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicemodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelName) ProtoMessage() {}

func (x *ModelName) ProtoReflect() protoreflect.Message {
	mi := &file_devicemodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelName.ProtoReflect.Descriptor instead.
func (*ModelName) Descriptor() ([]byte, []int) {
	return file_devicemodel_proto_rawDescGZIP(), []int{0}
}

func (x *ModelName) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

type FindAllModelsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllModelsParams) Reset() {
	*x = FindAllModelsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicemodel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllModelsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllModelsParams) ProtoMessage() {}

func (x *FindAllModelsParams) ProtoReflect() protoreflect.Message {
	mi := &file_devicemodel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllModelsParams.ProtoReflect.Descriptor instead.
func (*FindAllModelsParams) Descriptor() ([]byte, []int) {
	return file_devicemodel_proto_rawDescGZIP(), []int{1}
}

type DeviceModelResponseType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName      string     `protobuf:"bytes,1,opt,name=modelName,proto3" json:"modelName,omitempty"`                                // CP, Trio etc
	ModelNumber    string     `protobuf:"bytes,2,opt,name=modelNumber,proto3" json:"modelNumber,omitempty"`                            // 8300  etc
	DeviceType     DeviceType `protobuf:"varint,3,opt,name=deviceType,proto3,enum=devicemodel.DeviceType" json:"deviceType,omitempty"` // IPPHONE, VIDEO, HEADSET etc
	DeviceTypeUUID string     `protobuf:"bytes,4,opt,name=deviceTypeUUID,proto3" json:"deviceTypeUUID,omitempty"`                      // UUID for a device type a hexa
	Vendor         string     `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty"`                                      // CISCO, Polycom etc
	TypePatterns   []string   `protobuf:"bytes,6,rep,name=TypePatterns,proto3" json:"TypePatterns,omitempty"`                          // Trio CP what the device connection header string contains
}

func (x *DeviceModelResponseType) Reset() {
	*x = DeviceModelResponseType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicemodel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceModelResponseType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceModelResponseType) ProtoMessage() {}

func (x *DeviceModelResponseType) ProtoReflect() protoreflect.Message {
	mi := &file_devicemodel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceModelResponseType.ProtoReflect.Descriptor instead.
func (*DeviceModelResponseType) Descriptor() ([]byte, []int) {
	return file_devicemodel_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceModelResponseType) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *DeviceModelResponseType) GetModelNumber() string {
	if x != nil {
		return x.ModelNumber
	}
	return ""
}

func (x *DeviceModelResponseType) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_IPPHONE
}

func (x *DeviceModelResponseType) GetDeviceTypeUUID() string {
	if x != nil {
		return x.DeviceTypeUUID
	}
	return ""
}

func (x *DeviceModelResponseType) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *DeviceModelResponseType) GetTypePatterns() []string {
	if x != nil {
		return x.TypePatterns
	}
	return nil
}

type ListDeviceModelResponseType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*DeviceModelResponseType `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *ListDeviceModelResponseType) Reset() {
	*x = ListDeviceModelResponseType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicemodel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceModelResponseType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceModelResponseType) ProtoMessage() {}

func (x *ListDeviceModelResponseType) ProtoReflect() protoreflect.Message {
	mi := &file_devicemodel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceModelResponseType.ProtoReflect.Descriptor instead.
func (*ListDeviceModelResponseType) Descriptor() ([]byte, []int) {
	return file_devicemodel_proto_rawDescGZIP(), []int{3}
}

func (x *ListDeviceModelResponseType) GetModels() []*DeviceModelResponseType {
	if x != nil {
		return x.Models
	}
	return nil
}

var File_devicemodel_proto protoreflect.FileDescriptor

var file_devicemodel_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x29, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x22, 0x5b, 0x0a, 0x1b, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2a, 0x31, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x50, 0x50, 0x48, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x48, 0x45, 0x41, 0x44, 0x53, 0x45, 0x54, 0x10, 0x02, 0x32, 0xcc, 0x01, 0x0a, 0x12,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0d, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x28,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_devicemodel_proto_rawDescOnce sync.Once
	file_devicemodel_proto_rawDescData = file_devicemodel_proto_rawDesc
)

func file_devicemodel_proto_rawDescGZIP() []byte {
	file_devicemodel_proto_rawDescOnce.Do(func() {
		file_devicemodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_devicemodel_proto_rawDescData)
	})
	return file_devicemodel_proto_rawDescData
}

var file_devicemodel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_devicemodel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_devicemodel_proto_goTypes = []interface{}{
	(DeviceType)(0),                     // 0: devicemodel.DeviceType
	(*ModelName)(nil),                   // 1: devicemodel.ModelName
	(*FindAllModelsParams)(nil),         // 2: devicemodel.FindAllModelsParams
	(*DeviceModelResponseType)(nil),     // 3: devicemodel.DeviceModelResponseType
	(*ListDeviceModelResponseType)(nil), // 4: devicemodel.ListDeviceModelResponseType
}
var file_devicemodel_proto_depIdxs = []int32{
	0, // 0: devicemodel.DeviceModelResponseType.deviceType:type_name -> devicemodel.DeviceType
	3, // 1: devicemodel.ListDeviceModelResponseType.models:type_name -> devicemodel.DeviceModelResponseType
	1, // 2: devicemodel.DeviceModelService.FindDeviceModelByName:input_type -> devicemodel.ModelName
	2, // 3: devicemodel.DeviceModelService.FindAllModels:input_type -> devicemodel.FindAllModelsParams
	3, // 4: devicemodel.DeviceModelService.FindDeviceModelByName:output_type -> devicemodel.DeviceModelResponseType
	4, // 5: devicemodel.DeviceModelService.FindAllModels:output_type -> devicemodel.ListDeviceModelResponseType
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_devicemodel_proto_init() }
func file_devicemodel_proto_init() {
	if File_devicemodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devicemodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelName); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_devicemodel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllModelsParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_devicemodel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceModelResponseType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_devicemodel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceModelResponseType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_devicemodel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devicemodel_proto_goTypes,
		DependencyIndexes: file_devicemodel_proto_depIdxs,
		EnumInfos:         file_devicemodel_proto_enumTypes,
		MessageInfos:      file_devicemodel_proto_msgTypes,
	}.Build()
	File_devicemodel_proto = out.File
	file_devicemodel_proto_rawDesc = nil
	file_devicemodel_proto_goTypes = nil
	file_devicemodel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceModelServiceClient is the client API for DeviceModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceModelServiceClient interface {
	//A simple rpc ; Obtains the device model metadata by name
	FindDeviceModelByName(ctx context.Context, in *ModelName, opts ...grpc.CallOption) (*DeviceModelResponseType, error)
	FindAllModels(ctx context.Context, in *FindAllModelsParams, opts ...grpc.CallOption) (*ListDeviceModelResponseType, error)
}

type deviceModelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceModelServiceClient(cc grpc.ClientConnInterface) DeviceModelServiceClient {
	return &deviceModelServiceClient{cc}
}

func (c *deviceModelServiceClient) FindDeviceModelByName(ctx context.Context, in *ModelName, opts ...grpc.CallOption) (*DeviceModelResponseType, error) {
	out := new(DeviceModelResponseType)
	err := c.cc.Invoke(ctx, "/devicemodel.DeviceModelService/FindDeviceModelByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceModelServiceClient) FindAllModels(ctx context.Context, in *FindAllModelsParams, opts ...grpc.CallOption) (*ListDeviceModelResponseType, error) {
	out := new(ListDeviceModelResponseType)
	err := c.cc.Invoke(ctx, "/devicemodel.DeviceModelService/FindAllModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceModelServiceServer is the server API for DeviceModelService service.
type DeviceModelServiceServer interface {
	//A simple rpc ; Obtains the device model metadata by name
	FindDeviceModelByName(context.Context, *ModelName) (*DeviceModelResponseType, error)
	FindAllModels(context.Context, *FindAllModelsParams) (*ListDeviceModelResponseType, error)
}

// UnimplementedDeviceModelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceModelServiceServer struct {
}

func (*UnimplementedDeviceModelServiceServer) FindDeviceModelByName(context.Context, *ModelName) (*DeviceModelResponseType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDeviceModelByName not implemented")
}
func (*UnimplementedDeviceModelServiceServer) FindAllModels(context.Context, *FindAllModelsParams) (*ListDeviceModelResponseType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllModels not implemented")
}

func RegisterDeviceModelServiceServer(s *grpc.Server, srv DeviceModelServiceServer) {
	s.RegisterService(&_DeviceModelService_serviceDesc, srv)
}

func _DeviceModelService_FindDeviceModelByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceModelServiceServer).FindDeviceModelByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devicemodel.DeviceModelService/FindDeviceModelByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceModelServiceServer).FindDeviceModelByName(ctx, req.(*ModelName))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceModelService_FindAllModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllModelsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceModelServiceServer).FindAllModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devicemodel.DeviceModelService/FindAllModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceModelServiceServer).FindAllModels(ctx, req.(*FindAllModelsParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceModelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "devicemodel.DeviceModelService",
	HandlerType: (*DeviceModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindDeviceModelByName",
			Handler:    _DeviceModelService_FindDeviceModelByName_Handler,
		},
		{
			MethodName: "FindAllModels",
			Handler:    _DeviceModelService_FindAllModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devicemodel.proto",
}
