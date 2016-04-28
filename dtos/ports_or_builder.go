package dtos

import (
	"fmt"
	"io"
)

type PortsOrBuilder struct {
	present    map[string]bool
	PortsCount int32 `json:"portsCount"`
	//	PortsList *List[Port] `json:"portsList"`
	//	PortsOrBuilderList *List[? extends org.apache.mesos.Protos$PortOrBuilder] `json:"portsOrBuilderList"`

}

func (self *PortsOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *PortsOrBuilder) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *PortsOrBuilder) FormatText() string {
	return FormatText(self)
}

func (self *PortsOrBuilder) FormatJSON() string {
	return FormatJSON(self)
}

func (self *PortsOrBuilder) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *PortsOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on PortsOrBuilder", name)

	case "portsCount", "PortsCount":
		v, ok := value.(int32)
		if ok {
			self.PortsCount = v
			self.present["portsCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field portsCount/PortsCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *PortsOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on PortsOrBuilder", name)

	case "portsCount", "PortsCount":
		if self.present != nil {
			if _, ok := self.present["portsCount"]; ok {
				return self.PortsCount, nil
			}
		}
		return nil, fmt.Errorf("Field PortsCount no set on PortsCount %+v", self)

	}
}

func (self *PortsOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on PortsOrBuilder", name)

	case "portsCount", "PortsCount":
		self.present["portsCount"] = false

	}

	return nil
}

func (self *PortsOrBuilder) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type PortsOrBuilderList []*PortsOrBuilder

func (list *PortsOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *PortsOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *PortsOrBuilderList) FormatJSON() string {
	return FormatJSON(list)
}
