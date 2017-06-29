package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosResourcesObject struct {
	present map[string]bool

	Properties map[string]interface{} `json:"properties"`
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

func (self *MesosResourcesObject) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *MesosResourcesObject) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosResourcesObject) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosResourcesObject) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *MesosResourcesObject) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		v, ok := value.(map[string]interface{})
		if ok {
			self.Properties = v
			self.present["properties"] = true
			return nil
		} else {
			return fmt.Errorf("Field properties/Properties: value %v(%T) couldn't be cast to type map[string]interface{}", value, value)
		}

	}
}

func (self *MesosResourcesObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		if self.present != nil {
			if _, ok := self.present["properties"]; ok {
				return self.Properties, nil
			}
		}
		return nil, fmt.Errorf("Field Properties no set on Properties %+v", self)

	}
}

func (self *MesosResourcesObject) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosResourcesObject", name)

	case "properties", "Properties":
		self.present["properties"] = false

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
