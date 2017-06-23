package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDockerInfo struct {
	present map[string]bool

	DockerParameters SingularityDockerParameterList `json:"dockerParameters"`

	ForcePullImage bool `json:"forcePullImage"`

	Image string `json:"image,omitempty"`

	// Network *SingularityDockerNetworkType `json:"network"`

	Parameters map[string]string `json:"parameters"`

	PortMappings SingularityDockerPortMappingList `json:"portMappings"`

	Privileged bool `json:"privileged"`
}

func (self *SingularityDockerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDockerInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerInfo cannot copy the values from %#v", other)
}

func (self *SingularityDockerInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDockerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDockerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDockerInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDockerInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "dockerParameters", "DockerParameters":
		v, ok := value.(SingularityDockerParameterList)
		if ok {
			self.DockerParameters = v
			self.present["dockerParameters"] = true
			return nil
		} else {
			return fmt.Errorf("Field dockerParameters/DockerParameters: value %v(%T) couldn't be cast to type SingularityDockerParameterList", value, value)
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

	case "parameters", "Parameters":
		v, ok := value.(map[string]string)
		if ok {
			self.Parameters = v
			self.present["parameters"] = true
			return nil
		} else {
			return fmt.Errorf("Field parameters/Parameters: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "portMappings", "PortMappings":
		v, ok := value.(SingularityDockerPortMappingList)
		if ok {
			self.PortMappings = v
			self.present["portMappings"] = true
			return nil
		} else {
			return fmt.Errorf("Field portMappings/PortMappings: value %v(%T) couldn't be cast to type SingularityDockerPortMappingList", value, value)
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

func (self *SingularityDockerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "dockerParameters", "DockerParameters":
		if self.present != nil {
			if _, ok := self.present["dockerParameters"]; ok {
				return self.DockerParameters, nil
			}
		}
		return nil, fmt.Errorf("Field DockerParameters no set on DockerParameters %+v", self)

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

	case "parameters", "Parameters":
		if self.present != nil {
			if _, ok := self.present["parameters"]; ok {
				return self.Parameters, nil
			}
		}
		return nil, fmt.Errorf("Field Parameters no set on Parameters %+v", self)

	case "portMappings", "PortMappings":
		if self.present != nil {
			if _, ok := self.present["portMappings"]; ok {
				return self.PortMappings, nil
			}
		}
		return nil, fmt.Errorf("Field PortMappings no set on PortMappings %+v", self)

	case "privileged", "Privileged":
		if self.present != nil {
			if _, ok := self.present["privileged"]; ok {
				return self.Privileged, nil
			}
		}
		return nil, fmt.Errorf("Field Privileged no set on Privileged %+v", self)

	}
}

func (self *SingularityDockerInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "dockerParameters", "DockerParameters":
		self.present["dockerParameters"] = false

	case "forcePullImage", "ForcePullImage":
		self.present["forcePullImage"] = false

	case "image", "Image":
		self.present["image"] = false

	case "parameters", "Parameters":
		self.present["parameters"] = false

	case "portMappings", "PortMappings":
		self.present["portMappings"] = false

	case "privileged", "Privileged":
		self.present["privileged"] = false

	}

	return nil
}

func (self *SingularityDockerInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDockerInfoList []*SingularityDockerInfo

func (self *SingularityDockerInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerInfoList cannot copy the values from %#v", other)
}

func (list *SingularityDockerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDockerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDockerInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
