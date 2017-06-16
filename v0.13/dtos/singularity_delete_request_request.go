package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeleteRequestRequest struct {
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	Message string `json:"message,omitempty"`
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

func (self *SingularityDeleteRequestRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDeleteRequestRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeleteRequestRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeleteRequestRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDeleteRequestRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityDeleteRequestRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityDeleteRequestRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeleteRequestRequest", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "message", "Message":
		self.present["message"] = false

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
