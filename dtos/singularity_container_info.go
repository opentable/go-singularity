package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityContainerInfoSingularityContainerType string

const (
	SingularityContainerInfoSingularityContainerTypeMESOS  SingularityContainerInfoSingularityContainerType = "MESOS"
	SingularityContainerInfoSingularityContainerTypeDOCKER SingularityContainerInfoSingularityContainerType = "DOCKER"
)

type SingularityContainerInfo struct {
	Type    *SingularityContainerInfoSingularityContainerType `json:"type,omitempty"`
	Volumes *SingularityVolumeList                            `json:"volumes,omitempty"`
	Docker  *SingularityDockerInfo                            `json:"docker,omitempty"`
}

func (self *SingularityContainerInfo) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityContainerInfo) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityContainerInfo); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityContainerInfo cannot copy the values from %#v", other)
}

func (self *SingularityContainerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityContainerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityContainerInfo) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "type", "Type":
		v, ok := value.(SingularityContainerInfoSingularityContainerType)
		if ok {
			self.Type = &v
			return nil
		}
		return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityContainerInfoSingularityContainerType", value, value)

	case "volumes", "Volumes":
		v, ok := value.(SingularityVolumeList)
		if ok {
			self.Volumes = &v
			return nil
		}
		return fmt.Errorf("Field volumes/Volumes: value %v(%T) couldn't be cast to type SingularityVolumeList", value, value)

	case "docker", "Docker":
		v, ok := value.(*SingularityDockerInfo)
		if ok {
			self.Docker = v
			return nil
		}
		return fmt.Errorf("Field docker/Docker: value %v(%T) couldn't be cast to type *SingularityDockerInfo", value, value)

	}
}

func (self *SingularityContainerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "type", "Type":
		return *self.Type, nil
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "volumes", "Volumes":
		return *self.Volumes, nil
		return nil, fmt.Errorf("Field Volumes no set on Volumes %+v", self)

	case "docker", "Docker":
		return self.Docker, nil
		return nil, fmt.Errorf("Field Docker no set on Docker %+v", self)

	}
}

func (self *SingularityContainerInfo) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "type", "Type":
		self.Type = nil

	case "volumes", "Volumes":
		self.Volumes = nil

	case "docker", "Docker":
		self.Docker = nil

	}

	return nil
}

func (self *SingularityContainerInfo) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityContainerInfoList []*SingularityContainerInfo

func (self *SingularityContainerInfoList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityContainerInfoList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityContainerInfoList cannot copy the values from %#v", other)
}

func (list *SingularityContainerInfoList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityContainerInfoList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityContainerInfoList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
