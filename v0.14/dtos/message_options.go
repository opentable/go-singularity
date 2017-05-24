package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MessageOptions struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *MessageOptions `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	MessageSetWireFormat bool `json:"messageSetWireFormat"`

	NoStandardDescriptorAccessor bool `json:"noStandardDescriptorAccessor"`

	// ParserForType *com.google.protobuf.Parser<com.google.protobuf.DescriptorProtos$MessageOptions> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount"`

	// UninterpretedOptionList *List[UninterpretedOption] `json:"uninterpretedOptionList"`

	// UninterpretedOptionOrBuilderList *List[? extends com.google.protobuf.DescriptorProtos$UninterpretedOptionOrBuilder] `json:"uninterpretedOptionOrBuilderList"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *MessageOptions) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MessageOptions) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MessageOptions); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MessageOptions cannot copy the values from %#v", other)
}

func (self *MessageOptions) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *MessageOptions) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MessageOptions) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MessageOptions) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *MessageOptions) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MessageOptions", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*MessageOptions)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *MessageOptions", value, value)
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

	case "messageSetWireFormat", "MessageSetWireFormat":
		v, ok := value.(bool)
		if ok {
			self.MessageSetWireFormat = v
			self.present["messageSetWireFormat"] = true
			return nil
		} else {
			return fmt.Errorf("Field messageSetWireFormat/MessageSetWireFormat: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "noStandardDescriptorAccessor", "NoStandardDescriptorAccessor":
		v, ok := value.(bool)
		if ok {
			self.NoStandardDescriptorAccessor = v
			self.present["noStandardDescriptorAccessor"] = true
			return nil
		} else {
			return fmt.Errorf("Field noStandardDescriptorAccessor/NoStandardDescriptorAccessor: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		v, ok := value.(int32)
		if ok {
			self.UninterpretedOptionCount = v
			self.present["uninterpretedOptionCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field uninterpretedOptionCount/UninterpretedOptionCount: value %v(%T) couldn't be cast to type int32", value, value)
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

func (self *MessageOptions) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MessageOptions", name)

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

	case "messageSetWireFormat", "MessageSetWireFormat":
		if self.present != nil {
			if _, ok := self.present["messageSetWireFormat"]; ok {
				return self.MessageSetWireFormat, nil
			}
		}
		return nil, fmt.Errorf("Field MessageSetWireFormat no set on MessageSetWireFormat %+v", self)

	case "noStandardDescriptorAccessor", "NoStandardDescriptorAccessor":
		if self.present != nil {
			if _, ok := self.present["noStandardDescriptorAccessor"]; ok {
				return self.NoStandardDescriptorAccessor, nil
			}
		}
		return nil, fmt.Errorf("Field NoStandardDescriptorAccessor no set on NoStandardDescriptorAccessor %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		if self.present != nil {
			if _, ok := self.present["uninterpretedOptionCount"]; ok {
				return self.UninterpretedOptionCount, nil
			}
		}
		return nil, fmt.Errorf("Field UninterpretedOptionCount no set on UninterpretedOptionCount %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *MessageOptions) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MessageOptions", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "messageSetWireFormat", "MessageSetWireFormat":
		self.present["messageSetWireFormat"] = false

	case "noStandardDescriptorAccessor", "NoStandardDescriptorAccessor":
		self.present["noStandardDescriptorAccessor"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "uninterpretedOptionCount", "UninterpretedOptionCount":
		self.present["uninterpretedOptionCount"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *MessageOptions) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MessageOptionsList []*MessageOptions

func (self *MessageOptionsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MessageOptionsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MessageOptionsList cannot copy the values from %#v", other)
}

func (list *MessageOptionsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MessageOptionsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MessageOptionsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
