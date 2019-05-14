package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosResourcesObject struct {
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

func (self *MesosResourcesObject) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MesosResourcesObject) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosResourcesObject); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosResourcesObject cannot copy the values from %#v", other)
}

func (self *MesosResourcesObject) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosResourcesObject) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosResourcesObject) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		v, ok := value.(map[string]interface{})
		if ok {
			self.Properties = &v
			return nil
		}
		return fmt.Errorf("Field properties/Properties: value %v(%T) couldn't be cast to type map[string]interface{}", value, value)

	}
}

func (self *MesosResourcesObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		return *self.Properties, nil
		return nil, fmt.Errorf("Field Properties no set on Properties %+v", self)

	}
}

func (self *MesosResourcesObject) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		self.Properties = nil

	}

	return nil
}

func (self *MesosResourcesObject) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MesosResourcesObjectList []*MesosResourcesObject

func (self *MesosResourcesObjectList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosResourcesObjectList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosResourcesObjectList cannot copy the values from %#v", other)
}

func (list *MesosResourcesObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MesosResourcesObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosResourcesObjectList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
