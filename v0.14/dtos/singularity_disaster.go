package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDisasterSingularityDisasterType string

const (
	SingularityDisasterSingularityDisasterTypeEXCESSIVE_TASK_LAG SingularityDisasterSingularityDisasterType = "EXCESSIVE_TASK_LAG"
	SingularityDisasterSingularityDisasterTypeLOST_SLAVES        SingularityDisasterSingularityDisasterType = "LOST_SLAVES"
	SingularityDisasterSingularityDisasterTypeLOST_TASKS         SingularityDisasterSingularityDisasterType = "LOST_TASKS"
	SingularityDisasterSingularityDisasterTypeUSER_INITIATED     SingularityDisasterSingularityDisasterType = "USER_INITIATED"
)

type SingularityDisaster struct {
	present map[string]bool

	Active bool `json:"active"`

	Type SingularityDisasterSingularityDisasterType `json:"type"`
}

func (self *SingularityDisaster) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDisaster) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisaster); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisaster cannot copy the values from %#v", other)
}

func (self *SingularityDisaster) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDisaster) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDisaster) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDisaster) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDisaster) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisaster", name)

	case "active", "Active":
		v, ok := value.(bool)
		if ok {
			self.Active = v
			self.present["active"] = true
			return nil
		} else {
			return fmt.Errorf("Field active/Active: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "type", "Type":
		v, ok := value.(SingularityDisasterSingularityDisasterType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityDisasterSingularityDisasterType", value, value)
		}

	}
}

func (self *SingularityDisaster) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDisaster", name)

	case "active", "Active":
		if self.present != nil {
			if _, ok := self.present["active"]; ok {
				return self.Active, nil
			}
		}
		return nil, fmt.Errorf("Field Active no set on Active %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	}
}

func (self *SingularityDisaster) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisaster", name)

	case "active", "Active":
		self.present["active"] = false

	case "type", "Type":
		self.present["type"] = false

	}

	return nil
}

func (self *SingularityDisaster) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDisasterList []*SingularityDisaster

func (self *SingularityDisasterList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisasterList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisasterList cannot copy the values from %#v", other)
}

func (list *SingularityDisasterList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDisasterList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDisasterList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
