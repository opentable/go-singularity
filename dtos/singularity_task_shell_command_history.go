package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskShellCommandHistory struct {
	present map[string]bool

	ShellRequest *SingularityTaskShellCommandRequest `json:"shellRequest"`

	ShellUpdates SingularityTaskShellCommandUpdateList `json:"shellUpdates"`
}

func (self *SingularityTaskShellCommandHistory) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskShellCommandHistory) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandHistory); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandHistory cannot copy the values from %#v", other)
}

func (self *SingularityTaskShellCommandHistory) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskShellCommandHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandHistory) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskShellCommandHistory) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandHistory", name)

	case "shellRequest", "ShellRequest":
		v, ok := value.(*SingularityTaskShellCommandRequest)
		if ok {
			self.ShellRequest = v
			self.present["shellRequest"] = true
			return nil
		} else {
			return fmt.Errorf("Field shellRequest/ShellRequest: value %v(%T) couldn't be cast to type *SingularityTaskShellCommandRequest", value, value)
		}

	case "shellUpdates", "ShellUpdates":
		v, ok := value.(SingularityTaskShellCommandUpdateList)
		if ok {
			self.ShellUpdates = v
			self.present["shellUpdates"] = true
			return nil
		} else {
			return fmt.Errorf("Field shellUpdates/ShellUpdates: value %v(%T) couldn't be cast to type SingularityTaskShellCommandUpdateList", value, value)
		}

	}
}

func (self *SingularityTaskShellCommandHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandHistory", name)

	case "shellRequest", "ShellRequest":
		if self.present != nil {
			if _, ok := self.present["shellRequest"]; ok {
				return self.ShellRequest, nil
			}
		}
		return nil, fmt.Errorf("Field ShellRequest no set on ShellRequest %+v", self)

	case "shellUpdates", "ShellUpdates":
		if self.present != nil {
			if _, ok := self.present["shellUpdates"]; ok {
				return self.ShellUpdates, nil
			}
		}
		return nil, fmt.Errorf("Field ShellUpdates no set on ShellUpdates %+v", self)

	}
}

func (self *SingularityTaskShellCommandHistory) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandHistory", name)

	case "shellRequest", "ShellRequest":
		self.present["shellRequest"] = false

	case "shellUpdates", "ShellUpdates":
		self.present["shellUpdates"] = false

	}

	return nil
}

func (self *SingularityTaskShellCommandHistory) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskShellCommandHistoryList []*SingularityTaskShellCommandHistory

func (self *SingularityTaskShellCommandHistoryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandHistoryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandHistoryList cannot copy the values from %#v", other)
}

func (list *SingularityTaskShellCommandHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskShellCommandHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskShellCommandHistoryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
