package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type DockerInfoNetwork string

const (
	DockerInfoNetworkHOST   DockerInfoNetwork = "HOST"
	DockerInfoNetworkBRIDGE DockerInfoNetwork = "BRIDGE"
	DockerInfoNetworkNONE   DockerInfoNetwork = "NONE"
)

type DockerInfo struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *DockerInfo `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	ForcePullImage bool `json:"forcePullImage"`

	Image string `json:"image,omitempty"`

	ImageBytes *ByteString `json:"imageBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	Network DockerInfoNetwork `json:"network"`

	ParametersCount int32 `json:"parametersCount"`

	// ParametersList *List[Parameter] `json:"parametersList"`

	// ParametersOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"parametersOrBuilderList"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$ContainerInfo$DockerInfo> `json:"parserForType"`

	PortMappingsCount int32 `json:"portMappingsCount"`

	// PortMappingsList *List[PortMapping] `json:"portMappingsList"`

	// PortMappingsOrBuilderList *List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder] `json:"portMappingsOrBuilderList"`

	Privileged bool `json:"privileged"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *DockerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *DockerInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerInfo cannot copy the values from %#v", other)
}

func (self *DockerInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *DockerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *DockerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *DockerInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *DockerInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*DockerInfo)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *DockerInfo", value, value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "forcePullImage", "ForcePullImage":
		v, ok := value.(bool)
		if ok {
			self.ForcePullImage = v
			self.present["forcePullImage"] = true
			return nil
		} else {
			return fmt.Errorf("Field forcePullImage/ForcePullImage: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "image", "Image":
		v, ok := value.(string)
		if ok {
			self.Image = v
			self.present["image"] = true
			return nil
		} else {
			return fmt.Errorf("Field image/Image: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "imageBytes", "ImageBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.ImageBytes = v
			self.present["imageBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field imageBytes/ImageBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "network", "Network":
		v, ok := value.(DockerInfoNetwork)
		if ok {
			self.Network = v
			self.present["network"] = true
			return nil
		} else {
			return fmt.Errorf("Field network/Network: value %v(%T) couldn't be cast to type DockerInfoNetwork", value, value)
		}

	case "parametersCount", "ParametersCount":
		v, ok := value.(int32)
		if ok {
			self.ParametersCount = v
			self.present["parametersCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field parametersCount/ParametersCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "portMappingsCount", "PortMappingsCount":
		v, ok := value.(int32)
		if ok {
			self.PortMappingsCount = v
			self.present["portMappingsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field portMappingsCount/PortMappingsCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "privileged", "Privileged":
		v, ok := value.(bool)
		if ok {
			self.Privileged = v
			self.present["privileged"] = true
			return nil
		} else {
			return fmt.Errorf("Field privileged/Privileged: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	}
}

func (self *DockerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on DockerInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "forcePullImage", "ForcePullImage":
		if self.present != nil {
			if _, ok := self.present["forcePullImage"]; ok {
				return self.ForcePullImage, nil
			}
		}
		return nil, fmt.Errorf("Field ForcePullImage no set on ForcePullImage %+v", self)

	case "image", "Image":
		if self.present != nil {
			if _, ok := self.present["image"]; ok {
				return self.Image, nil
			}
		}
		return nil, fmt.Errorf("Field Image no set on Image %+v", self)

	case "imageBytes", "ImageBytes":
		if self.present != nil {
			if _, ok := self.present["imageBytes"]; ok {
				return self.ImageBytes, nil
			}
		}
		return nil, fmt.Errorf("Field ImageBytes no set on ImageBytes %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "network", "Network":
		if self.present != nil {
			if _, ok := self.present["network"]; ok {
				return self.Network, nil
			}
		}
		return nil, fmt.Errorf("Field Network no set on Network %+v", self)

	case "parametersCount", "ParametersCount":
		if self.present != nil {
			if _, ok := self.present["parametersCount"]; ok {
				return self.ParametersCount, nil
			}
		}
		return nil, fmt.Errorf("Field ParametersCount no set on ParametersCount %+v", self)

	case "portMappingsCount", "PortMappingsCount":
		if self.present != nil {
			if _, ok := self.present["portMappingsCount"]; ok {
				return self.PortMappingsCount, nil
			}
		}
		return nil, fmt.Errorf("Field PortMappingsCount no set on PortMappingsCount %+v", self)

	case "privileged", "Privileged":
		if self.present != nil {
			if _, ok := self.present["privileged"]; ok {
				return self.Privileged, nil
			}
		}
		return nil, fmt.Errorf("Field Privileged no set on Privileged %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *DockerInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "forcePullImage", "ForcePullImage":
		self.present["forcePullImage"] = false

	case "image", "Image":
		self.present["image"] = false

	case "imageBytes", "ImageBytes":
		self.present["imageBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "network", "Network":
		self.present["network"] = false

	case "parametersCount", "ParametersCount":
		self.present["parametersCount"] = false

	case "portMappingsCount", "PortMappingsCount":
		self.present["portMappingsCount"] = false

	case "privileged", "Privileged":
		self.present["privileged"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *DockerInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DockerInfoList []*DockerInfo

func (self *DockerInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerInfoList cannot copy the values from %#v", other)
}

func (list *DockerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DockerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DockerInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
