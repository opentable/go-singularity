package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type UnknownFieldSet struct {
	present map[string]bool

	DefaultInstanceForType *UnknownFieldSet `json:"defaultInstanceForType"`

	Initialized bool `json:"initialized"`

	// ParserForType *Parser `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	SerializedSizeAsMessageSet int32 `json:"serializedSizeAsMessageSet"`
}

func (self *UnknownFieldSet) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *UnknownFieldSet) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*UnknownFieldSet); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A UnknownFieldSet cannot copy the values from %#v", other)
}

func (self *UnknownFieldSet) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *UnknownFieldSet) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *UnknownFieldSet) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *UnknownFieldSet) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *UnknownFieldSet) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on UnknownFieldSet", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
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

	case "serializedSizeAsMessageSet", "SerializedSizeAsMessageSet":
		v, ok := value.(int32)
		if ok {
			self.SerializedSizeAsMessageSet = v
			self.present["serializedSizeAsMessageSet"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSizeAsMessageSet/SerializedSizeAsMessageSet: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *UnknownFieldSet) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on UnknownFieldSet", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

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

	case "serializedSizeAsMessageSet", "SerializedSizeAsMessageSet":
		if self.present != nil {
			if _, ok := self.present["serializedSizeAsMessageSet"]; ok {
				return self.SerializedSizeAsMessageSet, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSizeAsMessageSet no set on SerializedSizeAsMessageSet %+v", self)

	}
}

func (self *UnknownFieldSet) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on UnknownFieldSet", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "serializedSizeAsMessageSet", "SerializedSizeAsMessageSet":
		self.present["serializedSizeAsMessageSet"] = false

	}

	return nil
}

func (self *UnknownFieldSet) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type UnknownFieldSetList []*UnknownFieldSet

func (self *UnknownFieldSetList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*UnknownFieldSetList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A UnknownFieldSetList cannot copy the values from %#v", other)
}

func (list *UnknownFieldSetList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *UnknownFieldSetList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *UnknownFieldSetList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
