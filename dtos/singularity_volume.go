package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityVolumeSingularityDockerVolumeMode string

const (
	SingularityVolumeSingularityDockerVolumeModeRO SingularityVolumeSingularityDockerVolumeMode = "RO"
	SingularityVolumeSingularityDockerVolumeModeRW SingularityVolumeSingularityDockerVolumeMode = "RW"
)

type SingularityVolume struct {
	HostPath      *string                                       `json:"hostPath,omitempty"`
	Mode          *SingularityVolumeSingularityDockerVolumeMode `json:"mode,omitempty"`
	ContainerPath *string                                       `json:"containerPath,omitempty"`
}

func (self *SingularityVolume) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityVolume) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityVolume); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityVolume cannot copy the values from %#v", other)
}

func (self *SingularityVolume) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityVolume) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityVolume) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityVolume", name)

	case "hostPath", "HostPath":
		v, ok := value.(string)
		if ok {
			self.HostPath = &v
			return nil
		}
		return fmt.Errorf("Field hostPath/HostPath: value %v(%T) couldn't be cast to type string", value, value)

	case "mode", "Mode":
		v, ok := value.(SingularityVolumeSingularityDockerVolumeMode)
		if ok {
			self.Mode = &v
			return nil
		}
		return fmt.Errorf("Field mode/Mode: value %v(%T) couldn't be cast to type SingularityVolumeSingularityDockerVolumeMode", value, value)

	case "containerPath", "ContainerPath":
		v, ok := value.(string)
		if ok {
			self.ContainerPath = &v
			return nil
		}
		return fmt.Errorf("Field containerPath/ContainerPath: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityVolume) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityVolume", name)

	case "hostPath", "HostPath":
		return *self.HostPath, nil
		return nil, fmt.Errorf("Field HostPath no set on HostPath %+v", self)

	case "mode", "Mode":
		return *self.Mode, nil
		return nil, fmt.Errorf("Field Mode no set on Mode %+v", self)

	case "containerPath", "ContainerPath":
		return *self.ContainerPath, nil
		return nil, fmt.Errorf("Field ContainerPath no set on ContainerPath %+v", self)

	}
}

func (self *SingularityVolume) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityVolume", name)

	case "hostPath", "HostPath":
		self.HostPath = nil

	case "mode", "Mode":
		self.Mode = nil

	case "containerPath", "ContainerPath":
		self.ContainerPath = nil

	}

	return nil
}

func (self *SingularityVolume) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityVolumeList []*SingularityVolume

func (self *SingularityVolumeList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityVolumeList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityVolumeList cannot copy the values from %#v", other)
}

func (list *SingularityVolumeList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityVolumeList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityVolumeList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
