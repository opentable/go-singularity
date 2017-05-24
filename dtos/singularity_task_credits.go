package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskCredits struct {
	present map[string]bool

	Enabled bool `json:"enabled"`

	Remaining int32 `json:"remaining"`
}

func (self *SingularityTaskCredits) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskCredits) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskCredits); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskCredits cannot copy the values from %#v", other)
}

func (self *SingularityTaskCredits) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskCredits) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskCredits) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskCredits) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskCredits) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCredits", name)

	case "enabled", "Enabled":
		v, ok := value.(bool)
		if ok {
			self.Enabled = v
			self.present["enabled"] = true
			return nil
		} else {
			return fmt.Errorf("Field enabled/Enabled: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "remaining", "Remaining":
		v, ok := value.(int32)
		if ok {
			self.Remaining = v
			self.present["remaining"] = true
			return nil
		} else {
			return fmt.Errorf("Field remaining/Remaining: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *SingularityTaskCredits) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskCredits", name)

	case "enabled", "Enabled":
		if self.present != nil {
			if _, ok := self.present["enabled"]; ok {
				return self.Enabled, nil
			}
		}
		return nil, fmt.Errorf("Field Enabled no set on Enabled %+v", self)

	case "remaining", "Remaining":
		if self.present != nil {
			if _, ok := self.present["remaining"]; ok {
				return self.Remaining, nil
			}
		}
		return nil, fmt.Errorf("Field Remaining no set on Remaining %+v", self)

	}
}

func (self *SingularityTaskCredits) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCredits", name)

	case "enabled", "Enabled":
		self.present["enabled"] = false

	case "remaining", "Remaining":
		self.present["remaining"] = false

	}

	return nil
}

func (self *SingularityTaskCredits) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskCreditsList []*SingularityTaskCredits

func (self *SingularityTaskCreditsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskCreditsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskCreditsList cannot copy the values from %#v", other)
}

func (list *SingularityTaskCreditsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskCreditsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskCreditsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
