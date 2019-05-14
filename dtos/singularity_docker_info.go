package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDockerInfoSingularityDockerNetworkType string

const (
	SingularityDockerInfoSingularityDockerNetworkTypeHOST   SingularityDockerInfoSingularityDockerNetworkType = "HOST"
	SingularityDockerInfoSingularityDockerNetworkTypeBRIDGE SingularityDockerInfoSingularityDockerNetworkType = "BRIDGE"
	SingularityDockerInfoSingularityDockerNetworkTypeNONE   SingularityDockerInfoSingularityDockerNetworkType = "NONE"
)

type SingularityDockerInfo struct {
	Image            *string                                            `json:"image,omitempty"`
	Privileged       *bool                                              `json:"privileged,omitempty"`
	Network          *SingularityDockerInfoSingularityDockerNetworkType `json:"network,omitempty"`
	PortMappings     *SingularityDockerPortMappingList                  `json:"portMappings,omitempty"`
	ForcePullImage   *bool                                              `json:"forcePullImage,omitempty"`
	Parameters       *map[string]string                                 `json:"parameters,omitempty"`
	DockerParameters *SingularityDockerParameterList                    `json:"dockerParameters,omitempty"`
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

func (self *SingularityDockerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDockerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDockerInfo) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "image", "Image":
		v, ok := value.(string)
		if ok {
			self.Image = &v
			return nil
		}
		return fmt.Errorf("Field image/Image: value %v(%T) couldn't be cast to type string", value, value)

	case "privileged", "Privileged":
		v, ok := value.(bool)
		if ok {
			self.Privileged = &v
			return nil
		}
		return fmt.Errorf("Field privileged/Privileged: value %v(%T) couldn't be cast to type bool", value, value)

	case "network", "Network":
		v, ok := value.(SingularityDockerInfoSingularityDockerNetworkType)
		if ok {
			self.Network = &v
			return nil
		}
		return fmt.Errorf("Field network/Network: value %v(%T) couldn't be cast to type SingularityDockerInfoSingularityDockerNetworkType", value, value)

	case "portMappings", "PortMappings":
		v, ok := value.(SingularityDockerPortMappingList)
		if ok {
			self.PortMappings = &v
			return nil
		}
		return fmt.Errorf("Field portMappings/PortMappings: value %v(%T) couldn't be cast to type SingularityDockerPortMappingList", value, value)

	case "forcePullImage", "ForcePullImage":
		v, ok := value.(bool)
		if ok {
			self.ForcePullImage = &v
			return nil
		}
		return fmt.Errorf("Field forcePullImage/ForcePullImage: value %v(%T) couldn't be cast to type bool", value, value)

	case "parameters", "Parameters":
		v, ok := value.(map[string]string)
		if ok {
			self.Parameters = &v
			return nil
		}
		return fmt.Errorf("Field parameters/Parameters: value %v(%T) couldn't be cast to type map[string]string", value, value)

	case "dockerParameters", "DockerParameters":
		v, ok := value.(SingularityDockerParameterList)
		if ok {
			self.DockerParameters = &v
			return nil
		}
		return fmt.Errorf("Field dockerParameters/DockerParameters: value %v(%T) couldn't be cast to type SingularityDockerParameterList", value, value)

	}
}

func (self *SingularityDockerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "image", "Image":
		return *self.Image, nil
		return nil, fmt.Errorf("Field Image no set on Image %+v", self)

	case "privileged", "Privileged":
		return *self.Privileged, nil
		return nil, fmt.Errorf("Field Privileged no set on Privileged %+v", self)

	case "network", "Network":
		return *self.Network, nil
		return nil, fmt.Errorf("Field Network no set on Network %+v", self)

	case "portMappings", "PortMappings":
		return *self.PortMappings, nil
		return nil, fmt.Errorf("Field PortMappings no set on PortMappings %+v", self)

	case "forcePullImage", "ForcePullImage":
		return *self.ForcePullImage, nil
		return nil, fmt.Errorf("Field ForcePullImage no set on ForcePullImage %+v", self)

	case "parameters", "Parameters":
		return *self.Parameters, nil
		return nil, fmt.Errorf("Field Parameters no set on Parameters %+v", self)

	case "dockerParameters", "DockerParameters":
		return *self.DockerParameters, nil
		return nil, fmt.Errorf("Field DockerParameters no set on DockerParameters %+v", self)

	}
}

func (self *SingularityDockerInfo) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerInfo", name)

	case "image", "Image":
		self.Image = nil

	case "privileged", "Privileged":
		self.Privileged = nil

	case "network", "Network":
		self.Network = nil

	case "portMappings", "PortMappings":
		self.PortMappings = nil

	case "forcePullImage", "ForcePullImage":
		self.ForcePullImage = nil

	case "parameters", "Parameters":
		self.Parameters = nil

	case "dockerParameters", "DockerParameters":
		self.DockerParameters = nil

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
