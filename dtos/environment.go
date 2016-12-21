package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Environment struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Environment `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Environment> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`

	Variables VariableList `json:"variables"`

	VariablesCount int32 `json:"variablesCount"`

	VariablesList VariableList `json:"variablesList"`

	// VariablesOrBuilderList *List[? extends org.apache.mesos.Protos$Environment$VariableOrBuilder] `json:"variablesOrBuilderList"`

}

func (self *Environment) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Environment) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Environment); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Environment cannot copy the values from %#v", other)
}

func (self *Environment) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Environment) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Environment) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Environment) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Environment) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Environment)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Environment", value, value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	case "variables", "Variables":
		v, ok := value.(VariableList)
		if ok {
			self.Variables = v
			self.present["variables"] = true
			return nil
		} else {
			return fmt.Errorf("Field variables/Variables: value %v(%T) couldn't be cast to type VariableList", value, value)
		}

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

func (self *Environment) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	case "variables", "Variables":
		if self.present != nil {
			if _, ok := self.present["variables"]; ok {
				return self.Variables, nil
			}
		}
		return nil, fmt.Errorf("Field Variables no set on Variables %+v", self)

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

func (self *Environment) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Environment", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	case "variables", "Variables":
		self.present["variables"] = false

	case "variablesCount", "VariablesCount":
		self.present["variablesCount"] = false

	case "variablesList", "VariablesList":
		self.present["variablesList"] = false

	}

	return nil
}

func (self *Environment) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type EnvironmentList []*Environment

func (self *EnvironmentList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*EnvironmentList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A EnvironmentList cannot copy the values from %#v", other)
}

func (list *EnvironmentList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *EnvironmentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *EnvironmentList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
