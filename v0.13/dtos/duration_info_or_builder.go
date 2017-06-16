package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type DurationInfoOrBuilder struct {
	present map[string]bool

	Nanoseconds int64 `json:"nanoseconds"`
}

func (self *DurationInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *DurationInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DurationInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DurationInfoOrBuilder cannot absorb the values from %v", other)
}

func (self *DurationInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *DurationInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *DurationInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *DurationInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *DurationInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DurationInfoOrBuilder", name)

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

func (self *DurationInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on DurationInfoOrBuilder", name)

	case "nanoseconds", "Nanoseconds":
		if self.present != nil {
			if _, ok := self.present["nanoseconds"]; ok {
				return self.Nanoseconds, nil
			}
		}
		return nil, fmt.Errorf("Field Nanoseconds no set on Nanoseconds %+v", self)

	}
}

func (self *DurationInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DurationInfoOrBuilder", name)

	case "nanoseconds", "Nanoseconds":
		self.present["nanoseconds"] = false

	}

	return nil
}

func (self *DurationInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DurationInfoOrBuilderList []*DurationInfoOrBuilder

func (self *DurationInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DurationInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DurationInfoOrBuilder cannot absorb the values from %v", other)
}

func (list *DurationInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DurationInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DurationInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
