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
	present map[string]bool

	Docker *SingularityDockerInfo `json:"docker"`

	Type SingularityContainerInfoSingularityContainerType `json:"type"`

	Volumes SingularityVolumeList `json:"volumes"`
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

func (self *SingularityContainerInfo) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityContainerInfo) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityContainerInfo) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityContainerInfo) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityContainerInfo) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "docker", "Docker":
		v, ok := value.(*SingularityDockerInfo)
		if ok {
			self.Docker = v
			self.present["docker"] = true
			return nil
		} else {
			return fmt.Errorf("Field docker/Docker: value %v(%T) couldn't be cast to type *SingularityDockerInfo", value, value)
		}

	case "type", "Type":
		v, ok := value.(SingularityContainerInfoSingularityContainerType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityContainerInfoSingularityContainerType", value, value)
		}

	case "volumes", "Volumes":
		v, ok := value.(SingularityVolumeList)
		if ok {
			self.Volumes = v
			self.present["volumes"] = true
			return nil
		} else {
			return fmt.Errorf("Field volumes/Volumes: value %v(%T) couldn't be cast to type SingularityVolumeList", value, value)
		}

	}
}

func (self *SingularityContainerInfo) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "docker", "Docker":
		if self.present != nil {
			if _, ok := self.present["docker"]; ok {
				return self.Docker, nil
			}
		}
		return nil, fmt.Errorf("Field Docker no set on Docker %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "volumes", "Volumes":
		if self.present != nil {
			if _, ok := self.present["volumes"]; ok {
				return self.Volumes, nil
			}
		}
		return nil, fmt.Errorf("Field Volumes no set on Volumes %+v", self)

	}
}

func (self *SingularityContainerInfo) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityContainerInfo", name)

	case "docker", "Docker":
		self.present["docker"] = false

	case "type", "Type":
		self.present["type"] = false

	case "volumes", "Volumes":
		self.present["volumes"] = false

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
