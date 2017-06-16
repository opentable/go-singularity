package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExitCooldownRequest struct {
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	Message string `json:"message,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`
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

func (self *SingularityExitCooldownRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityExitCooldownRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExitCooldownRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExitCooldownRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityExitCooldownRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

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

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = v
			self.present["skipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *SingularityExitCooldownRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

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

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityExitCooldownRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExitCooldownRequest", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "message", "Message":
		self.present["message"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

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
