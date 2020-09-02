package model

// device model describes the different types to devices that this service supports
import (
	//"github.com/google/uuid"
	"encoding/json"
)

type DeviceModel struct {
	ModelName      string   `json:"modelName"`      // CP, Trio etc
	ModelNumber    string   `json:"ModelNumber"`    // 8300  etc
	DeviceType     string   `json:"DeviceType"`     // IPPHONE, VIDEO, HEADSET etc
	DeviceTypeUUID string   `json:"DeviceTypeUUID"` // UUID for a device
	Vendor         string   `json:"Vendor"`         // CISCO, Polycom etc
	TypePatterns   []string `json:"TypePatterns"`   // Trio CP what the device connection header string contains
}

var models = map[string]DeviceModel{
	"f957617b-e773-4864-95f1-df82a6fe5e09": DeviceModel{ModelName: "Trio", ModelNumber: "8300", DeviceType: "IPPHONE", DeviceTypeUUID: "f957617b-e773-4864-95f1-df82a6fe5e09", Vendor: "Polycom", TypePatterns: []string{"8300"}},
	"89c4e433-34ac-4063-8c03-e32d08e1b748": DeviceModel{ModelName: "IP Phone", ModelNumber: "7800", DeviceType: "IPPHONE", DeviceTypeUUID: "89c4e433-34ac-4063-8c03-e32d08e1b748", Vendor: "CISCO", TypePatterns: []string{"7800"}},
}

//constructor
func NewDeviceModel(modelName string, modelNumber string, deviceType string, deviceTypeUUID string, vendor string, typePatterns []string) *DeviceModel {
	return &DeviceModel{
		ModelName:      modelName,
		ModelNumber:    modelNumber,
		DeviceType:     deviceType,
		DeviceTypeUUID: deviceTypeUUID,
		Vendor:         vendor,
		TypePatterns:   typePatterns,
	}
}

// ToJSON to be used for marshalling of device model type
func (d DeviceModel) ToJSON() []byte {
	ToJSON, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return ToJSON
}
