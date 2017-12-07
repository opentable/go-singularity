package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPriorityFreezeParent struct {
	present map[string]bool

	Timestamp int64 `json:"timestamp"`

	User string `json:"user,omitempty"`

	PriorityFreeze *SingularityPriorityFreeze `json:"priorityFreeze"`
}

func (self *SingularityPriorityFreezeParent) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPriorityFreezeParent) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreezeParent); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreezeParent cannot copy the values from %#v", other)
}

func (self *SingularityPriorityFreezeParent) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPriorityFreezeParent) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPriorityFreezeParent) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPriorityFreezeParent) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPriorityFreezeParent) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "priorityFreeze", "PriorityFreeze":
		v, ok := value.(*SingularityPriorityFreeze)
		if ok {
			self.PriorityFreeze = v
			self.present["priorityFreeze"] = true
			return nil
		} else {
			return fmt.Errorf("Field priorityFreeze/PriorityFreeze: value %v(%T) couldn't be cast to type *SingularityPriorityFreeze", value, value)
		}

	}
}

func (self *SingularityPriorityFreezeParent) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "priorityFreeze", "PriorityFreeze":
		if self.present != nil {
			if _, ok := self.present["priorityFreeze"]; ok {
				return self.PriorityFreeze, nil
			}
		}
		return nil, fmt.Errorf("Field PriorityFreeze no set on PriorityFreeze %+v", self)

	}
}

func (self *SingularityPriorityFreezeParent) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPriorityFreezeParent", name)

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "user", "User":
		self.present["user"] = false

	case "priorityFreeze", "PriorityFreeze":
		self.present["priorityFreeze"] = false

	}

	return nil
}

func (self *SingularityPriorityFreezeParent) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPriorityFreezeParentList []*SingularityPriorityFreezeParent

func (self *SingularityPriorityFreezeParentList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPriorityFreezeParentList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPriorityFreezeParentList cannot copy the values from %#v", other)
}

func (list *SingularityPriorityFreezeParentList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPriorityFreezeParentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPriorityFreezeParentList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
