package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestGroup struct {
	Id         *string                `json:"id,omitempty"`
	RequestIds *swaggering.StringList `json:"requestIds,omitempty"`
	Metadata   *map[string]string     `json:"metadata,omitempty"`
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

func (self *SingularityRequestGroup) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestGroup) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestGroup) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = &v
			return nil
		}
		return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)

	case "requestIds", "RequestIds":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.RequestIds = &v
			return nil
		}
		return fmt.Errorf("Field requestIds/RequestIds: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "metadata", "Metadata":
		v, ok := value.(map[string]string)
		if ok {
			self.Metadata = &v
			return nil
		}
		return fmt.Errorf("Field metadata/Metadata: value %v(%T) couldn't be cast to type map[string]string", value, value)

	}
}

func (self *SingularityRequestGroup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "id", "Id":
		return *self.Id, nil
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "requestIds", "RequestIds":
		return *self.RequestIds, nil
		return nil, fmt.Errorf("Field RequestIds no set on RequestIds %+v", self)

	case "metadata", "Metadata":
		return *self.Metadata, nil
		return nil, fmt.Errorf("Field Metadata no set on Metadata %+v", self)

	}
}

func (self *SingularityRequestGroup) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestGroup", name)

	case "id", "Id":
		self.Id = nil

	case "requestIds", "RequestIds":
		self.RequestIds = nil

	case "metadata", "Metadata":
		self.Metadata = nil

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
