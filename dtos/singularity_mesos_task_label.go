package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityMesosTaskLabel struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (self *SingularityMesosTaskLabel) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityMesosTaskLabel) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMesosTaskLabel); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMesosTaskLabel cannot copy the values from %#v", other)
}

func (self *SingularityMesosTaskLabel) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityMesosTaskLabel) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityMesosTaskLabel) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosTaskLabel", name)

	case "key", "Key":
		v, ok := value.(string)
		if ok {
			self.Key = &v
			return nil
		}
		return fmt.Errorf("Field key/Key: value %v(%T) couldn't be cast to type string", value, value)

	case "value", "Value":
		v, ok := value.(string)
		if ok {
			self.Value = &v
			return nil
		}
		return fmt.Errorf("Field value/Value: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityMesosTaskLabel) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityMesosTaskLabel", name)

	case "key", "Key":
		return *self.Key, nil
		return nil, fmt.Errorf("Field Key no set on Key %+v", self)

	case "value", "Value":
		return *self.Value, nil
		return nil, fmt.Errorf("Field Value no set on Value %+v", self)

	}
}

func (self *SingularityMesosTaskLabel) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMesosTaskLabel", name)

	case "key", "Key":
		self.Key = nil

	case "value", "Value":
		self.Value = nil

	}

	return nil
}

func (self *SingularityMesosTaskLabel) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityMesosTaskLabelList []*SingularityMesosTaskLabel

func (self *SingularityMesosTaskLabelList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMesosTaskLabelList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMesosTaskLabelList cannot copy the values from %#v", other)
}

func (list *SingularityMesosTaskLabelList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityMesosTaskLabelList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityMesosTaskLabelList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
