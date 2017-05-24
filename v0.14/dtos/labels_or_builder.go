package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type LabelsOrBuilder struct {
	present map[string]bool

	LabelsCount int32 `json:"labelsCount"`

	// LabelsList *List[Label] `json:"labelsList"`

	// LabelsOrBuilderList *List[? extends org.apache.mesos.Protos$LabelOrBuilder] `json:"labelsOrBuilderList"`

}

func (self *LabelsOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *LabelsOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*LabelsOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A LabelsOrBuilder cannot copy the values from %#v", other)
}

func (self *LabelsOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *LabelsOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *LabelsOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *LabelsOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *LabelsOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on LabelsOrBuilder", name)

	case "labelsCount", "LabelsCount":
		v, ok := value.(int32)
		if ok {
			self.LabelsCount = v
			self.present["labelsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field labelsCount/LabelsCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *LabelsOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on LabelsOrBuilder", name)

	case "labelsCount", "LabelsCount":
		if self.present != nil {
			if _, ok := self.present["labelsCount"]; ok {
				return self.LabelsCount, nil
			}
		}
		return nil, fmt.Errorf("Field LabelsCount no set on LabelsCount %+v", self)

	}
}

func (self *LabelsOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on LabelsOrBuilder", name)

	case "labelsCount", "LabelsCount":
		self.present["labelsCount"] = false

	}

	return nil
}

func (self *LabelsOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type LabelsOrBuilderList []*LabelsOrBuilder

func (self *LabelsOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*LabelsOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A LabelsOrBuilderList cannot copy the values from %#v", other)
}

func (list *LabelsOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *LabelsOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *LabelsOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
