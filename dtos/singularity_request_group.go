package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestGroup struct {
	present map[string]bool

	Metadata map[string]string `json:"metadata"`

	Id string `json:"id,omitempty"`

	RequestIds swaggering.StringList `json:"requestIds"`
}

func (self *SingularityRequestGroup) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequestGroup) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestGroup); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestGroup cannot copy the values from %#v", other)
}

func (self *SingularityRequestGroup) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequestGroup) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestGroup) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestGroup) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequestGroup) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "metadata", "Metadata":
		v, ok := value.(map[string]string)
		if ok {
			self.Metadata = v
			self.present["metadata"] = true
			return nil
		} else {
			return fmt.Errorf("Field metadata/Metadata: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "requestIds", "RequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.RequestIds = v
			self.present["requestIds"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestIds/RequestIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
		}

	}
}

func (self *SingularityRequestGroup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "metadata", "Metadata":
		if self.present != nil {
			if _, ok := self.present["metadata"]; ok {
				return self.Metadata, nil
			}
		}
		return nil, fmt.Errorf("Field Metadata no set on Metadata %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "requestIds", "RequestIds":
		if self.present != nil {
			if _, ok := self.present["requestIds"]; ok {
				return self.RequestIds, nil
			}
		}
		return nil, fmt.Errorf("Field RequestIds no set on RequestIds %+v", self)

	}
}

func (self *SingularityRequestGroup) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "metadata", "Metadata":
		self.present["metadata"] = false

	case "id", "Id":
		self.present["id"] = false

	case "requestIds", "RequestIds":
		self.present["requestIds"] = false

	}

	return nil
}

func (self *SingularityRequestGroup) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestGroupList []*SingularityRequestGroup

func (self *SingularityRequestGroupList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestGroupList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestGroupList cannot copy the values from %#v", other)
}

func (list *SingularityRequestGroupList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestGroupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestGroupList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
