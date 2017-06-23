package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Set struct {
	present map[string]bool

	Empty bool `json:"empty"`
}

func (self *Set) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Set) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Set); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Set cannot copy the values from %#v", other)
}

func (self *Set) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Set) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Set) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Set) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Set) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Set", name)

	case "empty", "Empty":
		v, ok := value.(bool)
		if ok {
			self.Empty = v
			self.present["empty"] = true
			return nil
		} else {
			return fmt.Errorf("Field empty/Empty: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *Set) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Set", name)

	case "empty", "Empty":
		if self.present != nil {
			if _, ok := self.present["empty"]; ok {
				return self.Empty, nil
			}
		}
		return nil, fmt.Errorf("Field Empty no set on Empty %+v", self)

	}
}

func (self *Set) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Set", name)

	case "empty", "Empty":
		self.present["empty"] = false

	}

	return nil
}

func (self *Set) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SetList []*Set

func (self *SetList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SetList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SetList cannot copy the values from %#v", other)
}

func (list *SetList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SetList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SetList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
