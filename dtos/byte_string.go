package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ByteString struct {
	present map[string]bool

	Empty bool `json:"empty"`

	ValidUtf8 bool `json:"validUtf8"`
}

func (self *ByteString) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ByteString) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ByteString); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ByteString cannot copy the values from %#v", other)
}

func (self *ByteString) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ByteString) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ByteString) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ByteString) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ByteString) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ByteString", name)

	case "empty", "Empty":
		v, ok := value.(bool)
		if ok {
			self.Empty = v
			self.present["empty"] = true
			return nil
		} else {
			return fmt.Errorf("Field empty/Empty: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "validUtf8", "ValidUtf8":
		v, ok := value.(bool)
		if ok {
			self.ValidUtf8 = v
			self.present["validUtf8"] = true
			return nil
		} else {
			return fmt.Errorf("Field validUtf8/ValidUtf8: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *ByteString) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ByteString", name)

	case "empty", "Empty":
		if self.present != nil {
			if _, ok := self.present["empty"]; ok {
				return self.Empty, nil
			}
		}
		return nil, fmt.Errorf("Field Empty no set on Empty %+v", self)

	case "validUtf8", "ValidUtf8":
		if self.present != nil {
			if _, ok := self.present["validUtf8"]; ok {
				return self.ValidUtf8, nil
			}
		}
		return nil, fmt.Errorf("Field ValidUtf8 no set on ValidUtf8 %+v", self)

	}
}

func (self *ByteString) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ByteString", name)

	case "empty", "Empty":
		self.present["empty"] = false

	case "validUtf8", "ValidUtf8":
		self.present["validUtf8"] = false

	}

	return nil
}

func (self *ByteString) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ByteStringList []*ByteString

func (self *ByteStringList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ByteStringList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ByteStringList cannot copy the values from %#v", other)
}

func (list *ByteStringList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ByteStringList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ByteStringList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
