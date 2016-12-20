package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type FileDescriptor struct {
	present map[string]bool

	Dependencies FileDescriptorList `json:"dependencies"`

	// EnumTypes *List[EnumDescriptor] `json:"enumTypes"`

	// Extensions *List[FieldDescriptor] `json:"extensions"`

	MessageTypes DescriptorList `json:"messageTypes"`

	Name string `json:"name,omitempty"`

	Options *FileOptions `json:"options"`

	Package string `json:"package,omitempty"`

	PublicDependencies FileDescriptorList `json:"publicDependencies"`

	// Services *List[ServiceDescriptor] `json:"services"`

}

func (self *FileDescriptor) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *FileDescriptor) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FileDescriptor); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FileDescriptor cannot copy the values from %#v", other)
}

func (self *FileDescriptor) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *FileDescriptor) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *FileDescriptor) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *FileDescriptor) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *FileDescriptor) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FileDescriptor", name)

	case "dependencies", "Dependencies":
		v, ok := value.(FileDescriptorList)
		if ok {
			self.Dependencies = v
			self.present["dependencies"] = true
			return nil
		} else {
			return fmt.Errorf("Field dependencies/Dependencies: value %v(%T) couldn't be cast to type FileDescriptorList", value, value)
		}

	case "messageTypes", "MessageTypes":
		v, ok := value.(DescriptorList)
		if ok {
			self.MessageTypes = v
			self.present["messageTypes"] = true
			return nil
		} else {
			return fmt.Errorf("Field messageTypes/MessageTypes: value %v(%T) couldn't be cast to type DescriptorList", value, value)
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
		v, ok := value.(*FileOptions)
		if ok {
			self.Options = v
			self.present["options"] = true
			return nil
		} else {
			return fmt.Errorf("Field options/Options: value %v(%T) couldn't be cast to type *FileOptions", value, value)
		}

	case "package", "Package":
		v, ok := value.(string)
		if ok {
			self.Package = v
			self.present["package"] = true
			return nil
		} else {
			return fmt.Errorf("Field package/Package: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "publicDependencies", "PublicDependencies":
		v, ok := value.(FileDescriptorList)
		if ok {
			self.PublicDependencies = v
			self.present["publicDependencies"] = true
			return nil
		} else {
			return fmt.Errorf("Field publicDependencies/PublicDependencies: value %v(%T) couldn't be cast to type FileDescriptorList", value, value)
		}

	}
}

func (self *FileDescriptor) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on FileDescriptor", name)

	case "dependencies", "Dependencies":
		if self.present != nil {
			if _, ok := self.present["dependencies"]; ok {
				return self.Dependencies, nil
			}
		}
		return nil, fmt.Errorf("Field Dependencies no set on Dependencies %+v", self)

	case "messageTypes", "MessageTypes":
		if self.present != nil {
			if _, ok := self.present["messageTypes"]; ok {
				return self.MessageTypes, nil
			}
		}
		return nil, fmt.Errorf("Field MessageTypes no set on MessageTypes %+v", self)

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

	case "package", "Package":
		if self.present != nil {
			if _, ok := self.present["package"]; ok {
				return self.Package, nil
			}
		}
		return nil, fmt.Errorf("Field Package no set on Package %+v", self)

	case "publicDependencies", "PublicDependencies":
		if self.present != nil {
			if _, ok := self.present["publicDependencies"]; ok {
				return self.PublicDependencies, nil
			}
		}
		return nil, fmt.Errorf("Field PublicDependencies no set on PublicDependencies %+v", self)

	}
}

func (self *FileDescriptor) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on FileDescriptor", name)

	case "dependencies", "Dependencies":
		self.present["dependencies"] = false

	case "messageTypes", "MessageTypes":
		self.present["messageTypes"] = false

	case "name", "Name":
		self.present["name"] = false

	case "options", "Options":
		self.present["options"] = false

	case "package", "Package":
		self.present["package"] = false

	case "publicDependencies", "PublicDependencies":
		self.present["publicDependencies"] = false

	}

	return nil
}

func (self *FileDescriptor) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type FileDescriptorList []*FileDescriptor

func (self *FileDescriptorList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*FileDescriptorList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A FileDescriptorList cannot copy the values from %#v", other)
}

func (list *FileDescriptorList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *FileDescriptorList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *FileDescriptorList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
