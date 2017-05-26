package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type OfferIDOrBuilder struct {
	present map[string]bool

	Value string `json:"value,omitempty"`

	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *OfferIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *OfferIDOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*OfferIDOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A OfferIDOrBuilder cannot copy the values from %#v", other)
}

func (self *OfferIDOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *OfferIDOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *OfferIDOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *OfferIDOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *OfferIDOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on OfferIDOrBuilder", name)

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

func (self *OfferIDOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on OfferIDOrBuilder", name)

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

func (self *OfferIDOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on OfferIDOrBuilder", name)

	case "value", "Value":
		self.present["value"] = false

	case "valueBytes", "ValueBytes":
		self.present["valueBytes"] = false

	}

	return nil
}

func (self *OfferIDOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type OfferIDOrBuilderList []*OfferIDOrBuilder

func (self *OfferIDOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*OfferIDOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A OfferIDOrBuilderList cannot copy the values from %#v", other)
}

func (list *OfferIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *OfferIDOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *OfferIDOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
