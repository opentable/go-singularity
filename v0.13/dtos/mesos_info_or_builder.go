package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosInfoOrBuilder struct {
	present map[string]bool

	Image *Image `json:"image"`

	ImageOrBuilder *ImageOrBuilder `json:"imageOrBuilder"`
}

func (self *MesosInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MesosInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosInfoOrBuilder cannot absorb the values from %v", other)
}

func (self *MesosInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *MesosInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *MesosInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosInfoOrBuilder", name)

	case "image", "Image":
		v, ok := value.(*Image)
		if ok {
			self.Image = v
			self.present["image"] = true
			return nil
		} else {
			return fmt.Errorf("Field image/Image: value %v(%T) couldn't be cast to type *Image", value, value)
		}

	case "imageOrBuilder", "ImageOrBuilder":
		v, ok := value.(*ImageOrBuilder)
		if ok {
			self.ImageOrBuilder = v
			self.present["imageOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field imageOrBuilder/ImageOrBuilder: value %v(%T) couldn't be cast to type *ImageOrBuilder", value, value)
		}

	}
}

func (self *MesosInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosInfoOrBuilder", name)

	case "image", "Image":
		if self.present != nil {
			if _, ok := self.present["image"]; ok {
				return self.Image, nil
			}
		}
		return nil, fmt.Errorf("Field Image no set on Image %+v", self)

	case "imageOrBuilder", "ImageOrBuilder":
		if self.present != nil {
			if _, ok := self.present["imageOrBuilder"]; ok {
				return self.ImageOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field ImageOrBuilder no set on ImageOrBuilder %+v", self)

	}
}

func (self *MesosInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosInfoOrBuilder", name)

	case "image", "Image":
		self.present["image"] = false

	case "imageOrBuilder", "ImageOrBuilder":
		self.present["imageOrBuilder"] = false

	}

	return nil
}

func (self *MesosInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MesosInfoOrBuilderList []*MesosInfoOrBuilder

func (self *MesosInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosInfoOrBuilder cannot absorb the values from %v", other)
}

func (list *MesosInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MesosInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
