package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularitySandboxFile struct {
	Size  *int64  `json:"size,omitempty"`
	Mode  *string `json:"mode,omitempty"`
	Name  *string `json:"name,omitempty"`
	Mtime *int64  `json:"mtime,omitempty"`
}

func (self *SingularitySandboxFile) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularitySandboxFile) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySandboxFile); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySandboxFile cannot copy the values from %#v", other)
}

func (self *SingularitySandboxFile) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularitySandboxFile) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularitySandboxFile) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySandboxFile", name)

	case "size", "Size":
		v, ok := value.(int64)
		if ok {
			self.Size = &v
			return nil
		}
		return fmt.Errorf("Field size/Size: value %v(%T) couldn't be cast to type int64", value, value)

	case "mode", "Mode":
		v, ok := value.(string)
		if ok {
			self.Mode = &v
			return nil
		}
		return fmt.Errorf("Field mode/Mode: value %v(%T) couldn't be cast to type string", value, value)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	case "mtime", "Mtime":
		v, ok := value.(int64)
		if ok {
			self.Mtime = &v
			return nil
		}
		return fmt.Errorf("Field mtime/Mtime: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularitySandboxFile) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularitySandboxFile", name)

	case "size", "Size":
		return *self.Size, nil
		return nil, fmt.Errorf("Field Size no set on Size %+v", self)

	case "mode", "Mode":
		return *self.Mode, nil
		return nil, fmt.Errorf("Field Mode no set on Mode %+v", self)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "mtime", "Mtime":
		return *self.Mtime, nil
		return nil, fmt.Errorf("Field Mtime no set on Mtime %+v", self)

	}
}

func (self *SingularitySandboxFile) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySandboxFile", name)

	case "size", "Size":
		self.Size = nil

	case "mode", "Mode":
		self.Mode = nil

	case "name", "Name":
		self.Name = nil

	case "mtime", "Mtime":
		self.Mtime = nil

	}

	return nil
}

func (self *SingularitySandboxFile) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularitySandboxFileList []*SingularitySandboxFile

func (self *SingularitySandboxFileList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySandboxFileList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySandboxFileList cannot copy the values from %#v", other)
}

func (list *SingularitySandboxFileList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularitySandboxFileList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySandboxFileList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
