package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeleteRequestRequest struct {
	Message  *string `json:"message,omitempty"`
	ActionId *string `json:"actionId,omitempty"`
}

func (self *SingularityDeleteRequestRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeleteRequestRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeleteRequestRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeleteRequestRequest cannot copy the values from %#v", other)
}

func (self *SingularityDeleteRequestRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeleteRequestRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeleteRequestRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityDeleteRequestRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	}
}

func (self *SingularityDeleteRequestRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "message", "Message":
		self.Message = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	}

	return nil
}

func (self *SingularityDeleteRequestRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeleteRequestRequestList []*SingularityDeleteRequestRequest

func (self *SingularityDeleteRequestRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeleteRequestRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeleteRequestRequestList cannot copy the values from %#v", other)
}

func (list *SingularityDeleteRequestRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeleteRequestRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeleteRequestRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
