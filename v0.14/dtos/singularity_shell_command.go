package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityShellCommand struct {
	present map[string]bool

	LogfileName string `json:"logfileName,omitempty"`

	Name string `json:"name,omitempty"`

	Options swaggering.StringList `json:"options"`

	User string `json:"user,omitempty"`
}

func (self *SingularityShellCommand) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityShellCommand) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityShellCommand); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityShellCommand cannot copy the values from %#v", other)
}

func (self *SingularityShellCommand) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityShellCommand) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityShellCommand) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityShellCommand) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityShellCommand) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "logfileName", "LogfileName":
		v, ok := value.(string)
		if ok {
			self.LogfileName = v
			self.present["logfileName"] = true
			return nil
		} else {
			return fmt.Errorf("Field logfileName/LogfileName: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = v
			self.present["name"] = true
			return nil
		} else {
			return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "options", "Options":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Options = v
			self.present["options"] = true
			return nil
		} else {
			return fmt.Errorf("Field options/Options: value %v(%T) couldn't be cast to type StringList", value, value)
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

	}
}

func (self *SingularityShellCommand) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "logfileName", "LogfileName":
		if self.present != nil {
			if _, ok := self.present["logfileName"]; ok {
				return self.LogfileName, nil
			}
		}
		return nil, fmt.Errorf("Field LogfileName no set on LogfileName %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "options", "Options":
		if self.present != nil {
			if _, ok := self.present["options"]; ok {
				return self.Options, nil
			}
		}
		return nil, fmt.Errorf("Field Options no set on Options %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityShellCommand) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "logfileName", "LogfileName":
		self.present["logfileName"] = false

	case "name", "Name":
		self.present["name"] = false

	case "options", "Options":
		self.present["options"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityShellCommand) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityShellCommandList []*SingularityShellCommand

func (self *SingularityShellCommandList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityShellCommandList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityShellCommandList cannot copy the values from %#v", other)
}

func (list *SingularityShellCommandList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityShellCommandList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityShellCommandList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
