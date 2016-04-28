package dtos

import (
	"fmt"
	"io"
)

type EnvironmentOrBuilder struct {
	present        map[string]bool
	VariablesCount int32 `json:"variablesCount"`
	//	VariablesList *List[Variable] `json:"variablesList"`
	//	VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder] `json:"variablesOrBuilderList"`

}

func (self *EnvironmentOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *EnvironmentOrBuilder) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *EnvironmentOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *EnvironmentOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

func (self *EnvironmentOrBuilder) FieldsPresent() []string {
	return presenceFromMap(self.present)
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

	}

	return nil
}

func (self *EnvironmentOrBuilder) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type EnvironmentOrBuilderList []*EnvironmentOrBuilder

func (list *EnvironmentOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
