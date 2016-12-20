package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ContainerInfoOrBuilderType string

const (
	ContainerInfoOrBuilderTypeDOCKER ContainerInfoOrBuilderType = "DOCKER"
	ContainerInfoOrBuilderTypeMESOS  ContainerInfoOrBuilderType = "MESOS"
)

type ContainerInfoOrBuilder struct {
	present map[string]bool

	Docker *DockerInfo `json:"docker"`

	DockerOrBuilder *DockerInfoOrBuilder `json:"dockerOrBuilder"`

	Hostname string `json:"hostname,omitempty"`

	HostnameBytes *ByteString `json:"hostnameBytes"`

	Type ContainerInfoOrBuilderType `json:"type"`

	VolumesCount int32 `json:"volumesCount"`

	// VolumesList *List[Volume] `json:"volumesList"`

	// VolumesOrBuilderList *List[? extends org.apache.mesos.Protos$VolumeOrBuilder] `json:"volumesOrBuilderList"`

}

func (self *ContainerInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ContainerInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ContainerInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ContainerInfoOrBuilder cannot copy the values from %#v", other)
}

func (self *ContainerInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ContainerInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ContainerInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ContainerInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ContainerInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ContainerInfoOrBuilder", name)

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

	case "type", "Type":
		v, ok := value.(ContainerInfoOrBuilderType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type ContainerInfoOrBuilderType", value, value)
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

func (self *ContainerInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ContainerInfoOrBuilder", name)

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

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "volumesCount", "VolumesCount":
		if self.present != nil {
			if _, ok := self.present["volumesCount"]; ok {
				return self.VolumesCount, nil
			}
		}
		return nil, fmt.Errorf("Field VolumesCount no set on VolumesCount %+v", self)

	}
}

func (self *ContainerInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ContainerInfoOrBuilder", name)

	case "docker", "Docker":
		self.present["docker"] = false

	case "dockerOrBuilder", "DockerOrBuilder":
		self.present["dockerOrBuilder"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "hostnameBytes", "HostnameBytes":
		self.present["hostnameBytes"] = false

	case "type", "Type":
		self.present["type"] = false

	case "volumesCount", "VolumesCount":
		self.present["volumesCount"] = false

	}

	return nil
}

func (self *ContainerInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ContainerInfoOrBuilderList []*ContainerInfoOrBuilder

func (self *ContainerInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ContainerInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ContainerInfoOrBuilderList cannot copy the values from %#v", other)
}

func (list *ContainerInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ContainerInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ContainerInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
