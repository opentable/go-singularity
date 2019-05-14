package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityShellCommand struct {
	Options     *swaggering.StringList `json:"options,omitempty"`
	User        *string                `json:"user,omitempty"`
	LogfileName *string                `json:"logfileName,omitempty"`
	Name        *string                `json:"name,omitempty"`
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

func (self *SingularityShellCommand) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityShellCommand) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityShellCommand) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "options", "Options":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.Options = &v
			return nil
		}
		return fmt.Errorf("Field options/Options: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "logfileName", "LogfileName":
		v, ok := value.(string)
		if ok {
			self.LogfileName = &v
			return nil
		}
		return fmt.Errorf("Field logfileName/LogfileName: value %v(%T) couldn't be cast to type string", value, value)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityShellCommand) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "options", "Options":
		return *self.Options, nil
		return nil, fmt.Errorf("Field Options no set on Options %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "logfileName", "LogfileName":
		return *self.LogfileName, nil
		return nil, fmt.Errorf("Field LogfileName no set on LogfileName %+v", self)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	}
}

func (self *SingularityShellCommand) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityShellCommand", name)

	case "options", "Options":
		self.Options = nil

	case "user", "User":
		self.User = nil

	case "logfileName", "LogfileName":
		self.LogfileName = nil

	case "name", "Name":
		self.Name = nil

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
