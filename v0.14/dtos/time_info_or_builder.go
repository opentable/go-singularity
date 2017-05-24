package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type TimeInfoOrBuilder struct {
	present map[string]bool

	Nanoseconds int64 `json:"nanoseconds"`
}

func (self *TimeInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *TimeInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*TimeInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A TimeInfoOrBuilder cannot absorb the values from %v", other)
}

func (self *TimeInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *TimeInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *TimeInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *TimeInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *TimeInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on TimeInfoOrBuilder", name)

	case "nanoseconds", "Nanoseconds":
		v, ok := value.(int64)
		if ok {
			self.Nanoseconds = v
			self.present["nanoseconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field nanoseconds/Nanoseconds: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *TimeInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on TimeInfoOrBuilder", name)

	case "nanoseconds", "Nanoseconds":
		if self.present != nil {
			if _, ok := self.present["nanoseconds"]; ok {
				return self.Nanoseconds, nil
			}
		}
		return nil, fmt.Errorf("Field Nanoseconds no set on Nanoseconds %+v", self)

	}
}

func (self *TimeInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on TimeInfoOrBuilder", name)

	case "nanoseconds", "Nanoseconds":
		self.present["nanoseconds"] = false

	}

	return nil
}

func (self *TimeInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type TimeInfoOrBuilderList []*TimeInfoOrBuilder

func (self *TimeInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*TimeInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A TimeInfoOrBuilder cannot absorb the values from %v", other)
}

func (list *TimeInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *TimeInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *TimeInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
