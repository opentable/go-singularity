package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Labels struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Labels `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	LabelsCount int32 `json:"labelsCount"`

	// LabelsList *List[Label] `json:"labelsList"`

	// LabelsOrBuilderList *List[? extends org.apache.mesos.Protos$LabelOrBuilder] `json:"labelsOrBuilderList"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Labels> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Labels) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Labels) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Labels); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Labels cannot copy the values from %#v", other)
}

func (self *Labels) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Labels) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Labels) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Labels) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Labels) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Labels", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Labels)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Labels", value, value)
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

	case "labelsCount", "LabelsCount":
		v, ok := value.(int32)
		if ok {
			self.LabelsCount = v
			self.present["labelsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field labelsCount/LabelsCount: value %v(%T) couldn't be cast to type int32", value, value)
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

	}
}

func (self *Labels) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Labels", name)

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

	case "labelsCount", "LabelsCount":
		if self.present != nil {
			if _, ok := self.present["labelsCount"]; ok {
				return self.LabelsCount, nil
			}
		}
		return nil, fmt.Errorf("Field LabelsCount no set on LabelsCount %+v", self)

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

	}
}

func (self *Labels) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Labels", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "labelsCount", "LabelsCount":
		self.present["labelsCount"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Labels) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type LabelsList []*Labels

func (self *LabelsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*LabelsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A LabelsList cannot copy the values from %#v", other)
}

func (list *LabelsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *LabelsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *LabelsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
