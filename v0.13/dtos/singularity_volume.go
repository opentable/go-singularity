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
	present map[string]bool

	ContainerPath string `json:"containerPath,omitempty"`

	HostPath string `json:"hostPath,omitempty"`

	Mode SingularityVolumeSingularityDockerVolumeMode `json:"mode"`
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

func (self *SingularityVolume) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityVolume) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityVolume) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityVolume) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityVolume) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityVolume", name)

	case "containerPath", "ContainerPath":
		v, ok := value.(string)
		if ok {
			self.ContainerPath = v
			self.present["containerPath"] = true
			return nil
		} else {
			return fmt.Errorf("Field containerPath/ContainerPath: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "hostPath", "HostPath":
		v, ok := value.(string)
		if ok {
			self.HostPath = v
			self.present["hostPath"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostPath/HostPath: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "mode", "Mode":
		v, ok := value.(SingularityVolumeSingularityDockerVolumeMode)
		if ok {
			self.Mode = v
			self.present["mode"] = true
			return nil
		} else {
			return fmt.Errorf("Field mode/Mode: value %v(%T) couldn't be cast to type SingularityVolumeSingularityDockerVolumeMode", value, value)
		}

	}
}

func (self *SingularityVolume) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityVolume", name)

	case "containerPath", "ContainerPath":
		if self.present != nil {
			if _, ok := self.present["containerPath"]; ok {
				return self.ContainerPath, nil
			}
		}
		return nil, fmt.Errorf("Field ContainerPath no set on ContainerPath %+v", self)

	case "hostPath", "HostPath":
		if self.present != nil {
			if _, ok := self.present["hostPath"]; ok {
				return self.HostPath, nil
			}
		}
		return nil, fmt.Errorf("Field HostPath no set on HostPath %+v", self)

	case "mode", "Mode":
		if self.present != nil {
			if _, ok := self.present["mode"]; ok {
				return self.Mode, nil
			}
		}
		return nil, fmt.Errorf("Field Mode no set on Mode %+v", self)

	}
}

func (self *SingularityVolume) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityVolume", name)

	case "containerPath", "ContainerPath":
		self.present["containerPath"] = false

	case "hostPath", "HostPath":
		self.present["hostPath"] = false

	case "mode", "Mode":
		self.present["mode"] = false

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
