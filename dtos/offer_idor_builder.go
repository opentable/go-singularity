package dtos

import (
	"fmt"
	"io"
)

type OfferIDOrBuilder struct {
	present    map[string]bool
	Value      string      `json:"value,omitempty"`
	ValueBytes *ByteString `json:"valueBytes"`
}

func (self *OfferIDOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *OfferIDOrBuilder) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *OfferIDOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *OfferIDOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

func (self *OfferIDOrBuilder) FieldsPresent() []string {
	return presenceFromMap(self.present)
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
	return loadMapIntoDTO(from, self)
}

type OfferIDOrBuilderList []*OfferIDOrBuilder

func (list *OfferIDOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
