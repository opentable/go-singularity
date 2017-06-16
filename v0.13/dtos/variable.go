package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Variable struct {
	present map[string]bool

	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`
}

func (self *Variable) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Variable) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Variable); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Variable cannot copy the values from %#v", other)
}

func (self *Variable) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Variable) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Variable) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Variable) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Variable) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Variable", name)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = v
			self.present["name"] = true
			return nil
		} else {
			return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "value", "Value":
		v, ok := value.(string)
		if ok {
			self.Value = v
			self.present["value"] = true
			return nil
		} else {
			return fmt.Errorf("Field value/Value: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *Variable) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Variable", name)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "value", "Value":
		if self.present != nil {
			if _, ok := self.present["value"]; ok {
				return self.Value, nil
			}
		}
		return nil, fmt.Errorf("Field Value no set on Value %+v", self)

	}
}

func (self *Variable) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Variable", name)

	case "name", "Name":
		self.present["name"] = false

	case "value", "Value":
		self.present["value"] = false

	}

	return nil
}

func (self *Variable) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type VariableList []*Variable

func (self *VariableList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*VariableList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A VariableList cannot copy the values from %#v", other)
}

func (list *VariableList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *VariableList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *VariableList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
