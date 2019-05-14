package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPriorityFreeze struct {
	Message              *string  `json:"message,omitempty"`
	ActionId             *string  `json:"actionId,omitempty"`
	MinimumPriorityLevel *float64 `json:"minimumPriorityLevel,omitempty"`
	KillTasks            *bool    `json:"killTasks,omitempty"`
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

func (self *SingularityPriorityFreeze) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPriorityFreeze) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPriorityFreeze) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

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

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		v, ok := value.(float64)
		if ok {
			self.MinimumPriorityLevel = &v
			return nil
		}
		return fmt.Errorf("Field minimumPriorityLevel/MinimumPriorityLevel: value %v(%T) couldn't be cast to type float64", value, value)

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = &v
			return nil
		}
		return fmt.Errorf("Field killTasks/KillTasks: value %v(%T) couldn't be cast to type bool", value, value)

	}
}

func (self *SingularityPriorityFreeze) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		return *self.MinimumPriorityLevel, nil
		return nil, fmt.Errorf("Field MinimumPriorityLevel no set on MinimumPriorityLevel %+v", self)

	case "killTasks", "KillTasks":
		return *self.KillTasks, nil
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	}
}

func (self *SingularityPriorityFreeze) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreeze", name)

	case "message", "Message":
		self.Message = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "minimumPriorityLevel", "MinimumPriorityLevel":
		self.MinimumPriorityLevel = nil

	case "killTasks", "KillTasks":
		self.KillTasks = nil

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
