package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type FrameworkIDOrBuilder struct {
	present map[string]bool

	Value string `json:"value,omitempty"`

	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *FrameworkIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *FrameworkIDOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FrameworkIDOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FrameworkIDOrBuilder cannot copy the values from %#v", other)
}

func (self *FrameworkIDOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *FrameworkIDOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *FrameworkIDOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *FrameworkIDOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *FrameworkIDOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FrameworkIDOrBuilder", name)

	case "value", "Value":
		v, ok := value.(string)
		if ok {
			self.Value = v
			self.present["value"] = true
			return nil
		} else {
			return fmt.Errorf("Field value/Value: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "valueBytes", "ValueBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.ValueBytes = v
			self.present["valueBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field valueBytes/ValueBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	}
}

func (self *FrameworkIDOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on FrameworkIDOrBuilder", name)

	case "value", "Value":
		if self.present != nil {
			if _, ok := self.present["value"]; ok {
				return self.Value, nil
			}
		}
		return nil, fmt.Errorf("Field Value no set on Value %+v", self)

	case "valueBytes", "ValueBytes":
		if self.present != nil {
			if _, ok := self.present["valueBytes"]; ok {
				return self.ValueBytes, nil
			}
		}
		return nil, fmt.Errorf("Field ValueBytes no set on ValueBytes %+v", self)

	}
}

func (self *FrameworkIDOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FrameworkIDOrBuilder", name)

	case "value", "Value":
		self.present["value"] = false

	case "valueBytes", "ValueBytes":
		self.present["valueBytes"] = false

	}

	return nil
}

func (self *FrameworkIDOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type FrameworkIDOrBuilderList []*FrameworkIDOrBuilder

func (self *FrameworkIDOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FrameworkIDOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FrameworkIDOrBuilderList cannot copy the values from %#v", other)
}

func (list *FrameworkIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *FrameworkIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FrameworkIDOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
