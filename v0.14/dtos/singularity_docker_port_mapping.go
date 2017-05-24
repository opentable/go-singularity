package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDockerPortMappingSingularityPortMappingType string

const (
	SingularityDockerPortMappingSingularityPortMappingTypeLITERAL    SingularityDockerPortMappingSingularityPortMappingType = "LITERAL"
	SingularityDockerPortMappingSingularityPortMappingTypeFROM_OFFER SingularityDockerPortMappingSingularityPortMappingType = "FROM_OFFER"
)

type SingularityDockerPortMapping struct {
	present map[string]bool

	ContainerPort int32 `json:"containerPort"`

	ContainerPortType SingularityDockerPortMappingSingularityPortMappingType `json:"containerPortType"`

	HostPort int32 `json:"hostPort"`

	HostPortType SingularityDockerPortMappingSingularityPortMappingType `json:"hostPortType"`

	Protocol string `json:"protocol,omitempty"`
}

func (self *SingularityDockerPortMapping) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDockerPortMapping) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerPortMapping); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerPortMapping cannot copy the values from %#v", other)
}

func (self *SingularityDockerPortMapping) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDockerPortMapping) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDockerPortMapping) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDockerPortMapping) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDockerPortMapping) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerPortMapping", name)

	case "containerPort", "ContainerPort":
		v, ok := value.(int32)
		if ok {
			self.ContainerPort = v
			self.present["containerPort"] = true
			return nil
		} else {
			return fmt.Errorf("Field containerPort/ContainerPort: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "containerPortType", "ContainerPortType":
		v, ok := value.(SingularityDockerPortMappingSingularityPortMappingType)
		if ok {
			self.ContainerPortType = v
			self.present["containerPortType"] = true
			return nil
		} else {
			return fmt.Errorf("Field containerPortType/ContainerPortType: value %v(%T) couldn't be cast to type SingularityDockerPortMappingSingularityPortMappingType", value, value)
		}

	case "hostPort", "HostPort":
		v, ok := value.(int32)
		if ok {
			self.HostPort = v
			self.present["hostPort"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostPort/HostPort: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "hostPortType", "HostPortType":
		v, ok := value.(SingularityDockerPortMappingSingularityPortMappingType)
		if ok {
			self.HostPortType = v
			self.present["hostPortType"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostPortType/HostPortType: value %v(%T) couldn't be cast to type SingularityDockerPortMappingSingularityPortMappingType", value, value)
		}

	case "protocol", "Protocol":
		v, ok := value.(string)
		if ok {
			self.Protocol = v
			self.present["protocol"] = true
			return nil
		} else {
			return fmt.Errorf("Field protocol/Protocol: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityDockerPortMapping) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDockerPortMapping", name)

	case "containerPort", "ContainerPort":
		if self.present != nil {
			if _, ok := self.present["containerPort"]; ok {
				return self.ContainerPort, nil
			}
		}
		return nil, fmt.Errorf("Field ContainerPort no set on ContainerPort %+v", self)

	case "containerPortType", "ContainerPortType":
		if self.present != nil {
			if _, ok := self.present["containerPortType"]; ok {
				return self.ContainerPortType, nil
			}
		}
		return nil, fmt.Errorf("Field ContainerPortType no set on ContainerPortType %+v", self)

	case "hostPort", "HostPort":
		if self.present != nil {
			if _, ok := self.present["hostPort"]; ok {
				return self.HostPort, nil
			}
		}
		return nil, fmt.Errorf("Field HostPort no set on HostPort %+v", self)

	case "hostPortType", "HostPortType":
		if self.present != nil {
			if _, ok := self.present["hostPortType"]; ok {
				return self.HostPortType, nil
			}
		}
		return nil, fmt.Errorf("Field HostPortType no set on HostPortType %+v", self)

	case "protocol", "Protocol":
		if self.present != nil {
			if _, ok := self.present["protocol"]; ok {
				return self.Protocol, nil
			}
		}
		return nil, fmt.Errorf("Field Protocol no set on Protocol %+v", self)

	}
}

func (self *SingularityDockerPortMapping) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerPortMapping", name)

	case "containerPort", "ContainerPort":
		self.present["containerPort"] = false

	case "containerPortType", "ContainerPortType":
		self.present["containerPortType"] = false

	case "hostPort", "HostPort":
		self.present["hostPort"] = false

	case "hostPortType", "HostPortType":
		self.present["hostPortType"] = false

	case "protocol", "Protocol":
		self.present["protocol"] = false

	}

	return nil
}

func (self *SingularityDockerPortMapping) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDockerPortMappingList []*SingularityDockerPortMapping

func (self *SingularityDockerPortMappingList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerPortMappingList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerPortMappingList cannot copy the values from %#v", other)
}

func (list *SingularityDockerPortMappingList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDockerPortMappingList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDockerPortMappingList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
