package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ContainerInfoType string

const (
	ContainerInfoTypeDOCKER ContainerInfoType = "DOCKER"
	ContainerInfoTypeMESOS  ContainerInfoType = "MESOS"
)

type ContainerInfo struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *ContainerInfo `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Docker *DockerInfo `json:"docker"`

	DockerOrBuilder *DockerInfoOrBuilder `json:"dockerOrBuilder"`

	Hostname string `json:"hostname,omitempty"`

	HostnameBytes *ByteString `json:"hostnameBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$ContainerInfo> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	Type ContainerInfoType `json:"type"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`

	VolumesCount int32 `json:"volumesCount"`

	// VolumesList *List[Volume] `json:"volumesList"`

	// VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder] `json:"volumesOrBuilderList"`

}

func (self *ContainerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ContainerInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ContainerInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ContainerInfo cannot copy the values from %#v", other)
}

func (self *ContainerInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ContainerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ContainerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ContainerInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ContainerInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ContainerInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*ContainerInfo)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *ContainerInfo", value, value)
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

	case "docker", "Docker":
		v, ok := value.(*DockerInfo)
		if ok {
			self.Docker = v
			self.present["docker"] = true
			return nil
		} else {
			return fmt.Errorf("Field docker/Docker: value %v(%T) couldn't be cast to type *DockerInfo", value, value)
		}

	case "dockerOrBuilder", "DockerOrBuilder":
		v, ok := value.(*DockerInfoOrBuilder)
		if ok {
			self.DockerOrBuilder = v
			self.present["dockerOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field dockerOrBuilder/DockerOrBuilder: value %v(%T) couldn't be cast to type *DockerInfoOrBuilder", value, value)
		}

	case "hostname", "Hostname":
		v, ok := value.(string)
		if ok {
			self.Hostname = v
			self.present["hostname"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostname/Hostname: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "hostnameBytes", "HostnameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.HostnameBytes = v
			self.present["hostnameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostnameBytes/HostnameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "type", "Type":
		v, ok := value.(ContainerInfoType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type ContainerInfoType", value, value)
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

	case "volumesCount", "VolumesCount":
		v, ok := value.(int32)
		if ok {
			self.VolumesCount = v
			self.present["volumesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field volumesCount/VolumesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *ContainerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ContainerInfo", name)

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

	case "docker", "Docker":
		if self.present != nil {
			if _, ok := self.present["docker"]; ok {
				return self.Docker, nil
			}
		}
		return nil, fmt.Errorf("Field Docker no set on Docker %+v", self)

	case "dockerOrBuilder", "DockerOrBuilder":
		if self.present != nil {
			if _, ok := self.present["dockerOrBuilder"]; ok {
				return self.DockerOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field DockerOrBuilder no set on DockerOrBuilder %+v", self)

	case "hostname", "Hostname":
		if self.present != nil {
			if _, ok := self.present["hostname"]; ok {
				return self.Hostname, nil
			}
		}
		return nil, fmt.Errorf("Field Hostname no set on Hostname %+v", self)

	case "hostnameBytes", "HostnameBytes":
		if self.present != nil {
			if _, ok := self.present["hostnameBytes"]; ok {
				return self.HostnameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field HostnameBytes no set on HostnameBytes %+v", self)

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

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	case "volumesCount", "VolumesCount":
		if self.present != nil {
			if _, ok := self.present["volumesCount"]; ok {
				return self.VolumesCount, nil
			}
		}
		return nil, fmt.Errorf("Field VolumesCount no set on VolumesCount %+v", self)

	}
}

func (self *ContainerInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ContainerInfo", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "docker", "Docker":
		self.present["docker"] = false

	case "dockerOrBuilder", "DockerOrBuilder":
		self.present["dockerOrBuilder"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "hostnameBytes", "HostnameBytes":
		self.present["hostnameBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "type", "Type":
		self.present["type"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	case "volumesCount", "VolumesCount":
		self.present["volumesCount"] = false

	}

	return nil
}

func (self *ContainerInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ContainerInfoList []*ContainerInfo

func (self *ContainerInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ContainerInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ContainerInfoList cannot copy the values from %#v", other)
}

func (list *ContainerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ContainerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ContainerInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
