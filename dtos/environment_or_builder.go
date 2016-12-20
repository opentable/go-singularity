package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type EnvironmentOrBuilder struct {
	present map[string]bool

	VariablesCount int32 `json:"variablesCount"`

	VariablesList VariableList `json:"variablesList"`

	// VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder] `json:"variablesOrBuilderList"`

}

func (self *EnvironmentOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *EnvironmentOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*EnvironmentOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A EnvironmentOrBuilder cannot copy the values from %#v", other)
}

func (self *EnvironmentOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *EnvironmentOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *EnvironmentOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *EnvironmentOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *EnvironmentOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on EnvironmentOrBuilder", name)

	case "variablesCount", "VariablesCount":
		v, ok := value.(int32)
		if ok {
			self.VariablesCount = v
			self.present["variablesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field variablesCount/VariablesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "variablesList", "VariablesList":
		v, ok := value.(VariableList)
		if ok {
			self.VariablesList = v
			self.present["variablesList"] = true
			return nil
		} else {
			return fmt.Errorf("Field variablesList/VariablesList: value %v(%T) couldn't be cast to type VariableList", value, value)
		}

	}
}

func (self *EnvironmentOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on EnvironmentOrBuilder", name)

	case "variablesCount", "VariablesCount":
		if self.present != nil {
			if _, ok := self.present["variablesCount"]; ok {
				return self.VariablesCount, nil
			}
		}
		return nil, fmt.Errorf("Field VariablesCount no set on VariablesCount %+v", self)

	case "variablesList", "VariablesList":
		if self.present != nil {
			if _, ok := self.present["variablesList"]; ok {
				return self.VariablesList, nil
			}
		}
		return nil, fmt.Errorf("Field VariablesList no set on VariablesList %+v", self)

	}
}

func (self *EnvironmentOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on EnvironmentOrBuilder", name)

	case "variablesCount", "VariablesCount":
		self.present["variablesCount"] = false

	case "variablesList", "VariablesList":
		self.present["variablesList"] = false

	}

	return nil
}

func (self *EnvironmentOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type EnvironmentOrBuilderList []*EnvironmentOrBuilder

func (self *EnvironmentOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*EnvironmentOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A EnvironmentOrBuilderList cannot copy the values from %#v", other)
}

func (list *EnvironmentOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *EnvironmentOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EnvironmentOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
