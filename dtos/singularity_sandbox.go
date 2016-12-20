package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularitySandbox struct {
	present map[string]bool

	CurrentDirectory string `json:"currentDirectory,omitempty"`

	Files SingularitySandboxFileList `json:"files"`

	FullPathToRoot string `json:"fullPathToRoot,omitempty"`

	SlaveHostname string `json:"slaveHostname,omitempty"`
}

func (self *SingularitySandbox) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularitySandbox) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySandbox); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySandbox cannot copy the values from %#v", other)
}

func (self *SingularitySandbox) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularitySandbox) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularitySandbox) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularitySandbox) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularitySandbox) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySandbox", name)

	case "currentDirectory", "CurrentDirectory":
		v, ok := value.(string)
		if ok {
			self.CurrentDirectory = v
			self.present["currentDirectory"] = true
			return nil
		} else {
			return fmt.Errorf("Field currentDirectory/CurrentDirectory: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "files", "Files":
		v, ok := value.(SingularitySandboxFileList)
		if ok {
			self.Files = v
			self.present["files"] = true
			return nil
		} else {
			return fmt.Errorf("Field files/Files: value %v(%T) couldn't be cast to type SingularitySandboxFileList", value, value)
		}

	case "fullPathToRoot", "FullPathToRoot":
		v, ok := value.(string)
		if ok {
			self.FullPathToRoot = v
			self.present["fullPathToRoot"] = true
			return nil
		} else {
			return fmt.Errorf("Field fullPathToRoot/FullPathToRoot: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "slaveHostname", "SlaveHostname":
		v, ok := value.(string)
		if ok {
			self.SlaveHostname = v
			self.present["slaveHostname"] = true
			return nil
		} else {
			return fmt.Errorf("Field slaveHostname/SlaveHostname: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularitySandbox) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularitySandbox", name)

	case "currentDirectory", "CurrentDirectory":
		if self.present != nil {
			if _, ok := self.present["currentDirectory"]; ok {
				return self.CurrentDirectory, nil
			}
		}
		return nil, fmt.Errorf("Field CurrentDirectory no set on CurrentDirectory %+v", self)

	case "files", "Files":
		if self.present != nil {
			if _, ok := self.present["files"]; ok {
				return self.Files, nil
			}
		}
		return nil, fmt.Errorf("Field Files no set on Files %+v", self)

	case "fullPathToRoot", "FullPathToRoot":
		if self.present != nil {
			if _, ok := self.present["fullPathToRoot"]; ok {
				return self.FullPathToRoot, nil
			}
		}
		return nil, fmt.Errorf("Field FullPathToRoot no set on FullPathToRoot %+v", self)

	case "slaveHostname", "SlaveHostname":
		if self.present != nil {
			if _, ok := self.present["slaveHostname"]; ok {
				return self.SlaveHostname, nil
			}
		}
		return nil, fmt.Errorf("Field SlaveHostname no set on SlaveHostname %+v", self)

	}
}

func (self *SingularitySandbox) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySandbox", name)

	case "currentDirectory", "CurrentDirectory":
		self.present["currentDirectory"] = false

	case "files", "Files":
		self.present["files"] = false

	case "fullPathToRoot", "FullPathToRoot":
		self.present["fullPathToRoot"] = false

	case "slaveHostname", "SlaveHostname":
		self.present["slaveHostname"] = false

	}

	return nil
}

func (self *SingularitySandbox) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularitySandboxList []*SingularitySandbox

func (self *SingularitySandboxList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySandboxList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySandboxList cannot copy the values from %#v", other)
}

func (list *SingularitySandboxList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularitySandboxList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySandboxList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
