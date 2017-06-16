package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type DockerInfoOrBuilderNetwork string

const (
	DockerInfoOrBuilderNetworkHOST   DockerInfoOrBuilderNetwork = "HOST"
	DockerInfoOrBuilderNetworkBRIDGE DockerInfoOrBuilderNetwork = "BRIDGE"
	DockerInfoOrBuilderNetworkNONE   DockerInfoOrBuilderNetwork = "NONE"
)

type DockerInfoOrBuilder struct {
	present map[string]bool

	ForcePullImage bool `json:"forcePullImage"`

	Image string `json:"image,omitempty"`

	ImageBytes *ByteString `json:"imageBytes"`

	Network DockerInfoOrBuilderNetwork `json:"network"`

	ParametersCount int32 `json:"parametersCount"`

	// ParametersList *List[Parameter] `json:"parametersList"`

	// ParametersOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"parametersOrBuilderList"`

	PortMappingsCount int32 `json:"portMappingsCount"`

	// PortMappingsList *List[PortMapping] `json:"portMappingsList"`

	// PortMappingsOrBuilderList *List[? extends org.apache.mesos.Protos$ContainerInfo$DockerInfo$PortMappingOrBuilder] `json:"portMappingsOrBuilderList"`

	Privileged bool `json:"privileged"`
}

func (self *DockerInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *DockerInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerInfoOrBuilder cannot copy the values from %#v", other)
}

func (self *DockerInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *DockerInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *DockerInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *DockerInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *DockerInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerInfoOrBuilder", name)

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

	case "network", "Network":
		v, ok := value.(DockerInfoOrBuilderNetwork)
		if ok {
			self.Network = v
			self.present["network"] = true
			return nil
		} else {
			return fmt.Errorf("Field network/Network: value %v(%T) couldn't be cast to type DockerInfoOrBuilderNetwork", value, value)
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

	}
}

func (self *DockerInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on DockerInfoOrBuilder", name)

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

	}
}

func (self *DockerInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerInfoOrBuilder", name)

	case "forcePullImage", "ForcePullImage":
		self.present["forcePullImage"] = false

	case "image", "Image":
		self.present["image"] = false

	case "imageBytes", "ImageBytes":
		self.present["imageBytes"] = false

	case "network", "Network":
		self.present["network"] = false

	case "parametersCount", "ParametersCount":
		self.present["parametersCount"] = false

	case "portMappingsCount", "PortMappingsCount":
		self.present["portMappingsCount"] = false

	case "privileged", "Privileged":
		self.present["privileged"] = false

	}

	return nil
}

func (self *DockerInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DockerInfoOrBuilderList []*DockerInfoOrBuilder

func (self *DockerInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerInfoOrBuilderList cannot copy the values from %#v", other)
}

func (list *DockerInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DockerInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DockerInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
