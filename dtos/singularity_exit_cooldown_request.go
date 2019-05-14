package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExitCooldownRequest struct {
	Message          *string `json:"message,omitempty"`
	ActionId         *string `json:"actionId,omitempty"`
	SkipHealthchecks *bool   `json:"skipHealthchecks,omitempty"`
}

func (self *SingularityExitCooldownRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityExitCooldownRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExitCooldownRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExitCooldownRequest cannot copy the values from %#v", other)
}

func (self *SingularityExitCooldownRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExitCooldownRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExitCooldownRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

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

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	}
}

func (self *SingularityExitCooldownRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityExitCooldownRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

	case "message", "Message":
		self.Message = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	}

	return nil
}

func (self *SingularityExitCooldownRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityExitCooldownRequestList []*SingularityExitCooldownRequest

func (self *SingularityExitCooldownRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExitCooldownRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExitCooldownRequestList cannot copy the values from %#v", other)
}

func (list *SingularityExitCooldownRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityExitCooldownRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExitCooldownRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
