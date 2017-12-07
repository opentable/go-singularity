package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPriorityFreeze struct {
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	MinimumPriorityLevel float64 `json:"minimumPriorityLevel"`

	KillTasks bool `json:"killTasks"`

	Message string `json:"message,omitempty"`
}

func (self *SingularityPriorityFreeze) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPriorityFreeze) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreeze); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreeze cannot copy the values from %#v", other)
}

func (self *SingularityPriorityFreeze) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPriorityFreeze) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPriorityFreeze) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPriorityFreeze) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPriorityFreeze) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		v, ok := value.(float64)
		if ok {
			self.MinimumPriorityLevel = v
			self.present["minimumPriorityLevel"] = true
			return nil
		} else {
			return fmt.Errorf("Field minimumPriorityLevel/MinimumPriorityLevel: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = v
			self.present["killTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field killTasks/KillTasks: value %v(%T) couldn't be cast to type bool", value, value)
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

func (self *SingularityPriorityFreeze) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		if self.present != nil {
			if _, ok := self.present["minimumPriorityLevel"]; ok {
				return self.MinimumPriorityLevel, nil
			}
		}
		return nil, fmt.Errorf("Field MinimumPriorityLevel no set on MinimumPriorityLevel %+v", self)

	case "killTasks", "KillTasks":
		if self.present != nil {
			if _, ok := self.present["killTasks"]; ok {
				return self.KillTasks, nil
			}
		}
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityPriorityFreeze) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		self.present["minimumPriorityLevel"] = false

	case "killTasks", "KillTasks":
		self.present["killTasks"] = false

	case "message", "Message":
		self.present["message"] = false

	}

	return nil
}

func (self *SingularityPriorityFreeze) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPriorityFreezeList []*SingularityPriorityFreeze

func (self *SingularityPriorityFreezeList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreezeList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreezeList cannot copy the values from %#v", other)
}

func (list *SingularityPriorityFreezeList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPriorityFreezeList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPriorityFreezeList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
