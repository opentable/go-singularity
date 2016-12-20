package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Descriptor struct {
	present map[string]bool

	ContainingType *Descriptor `json:"containingType"`

	// EnumTypes *List[EnumDescriptor] `json:"enumTypes"`

	// Extensions *List[FieldDescriptor] `json:"extensions"`

	// Fields *List[FieldDescriptor] `json:"fields"`

	File *FileDescriptor `json:"file"`

	FullName string `json:"fullName,omitempty"`

	Index int32 `json:"index"`

	Name string `json:"name,omitempty"`

	NestedTypes DescriptorList `json:"nestedTypes"`

	Options *MessageOptions `json:"options"`
}

func (self *Descriptor) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Descriptor) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Descriptor); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Descriptor cannot copy the values from %#v", other)
}

func (self *Descriptor) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Descriptor) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Descriptor) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Descriptor) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Descriptor) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Descriptor", name)

	case "containingType", "ContainingType":
		v, ok := value.(*Descriptor)
		if ok {
			self.ContainingType = v
			self.present["containingType"] = true
			return nil
		} else {
			return fmt.Errorf("Field containingType/ContainingType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "file", "File":
		v, ok := value.(*FileDescriptor)
		if ok {
			self.File = v
			self.present["file"] = true
			return nil
		} else {
			return fmt.Errorf("Field file/File: value %v(%T) couldn't be cast to type *FileDescriptor", value, value)
		}

	case "fullName", "FullName":
		v, ok := value.(string)
		if ok {
			self.FullName = v
			self.present["fullName"] = true
			return nil
		} else {
			return fmt.Errorf("Field fullName/FullName: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "index", "Index":
		v, ok := value.(int32)
		if ok {
			self.Index = v
			self.present["index"] = true
			return nil
		} else {
			return fmt.Errorf("Field index/Index: value %v(%T) couldn't be cast to type int32", value, value)
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

	case "nestedTypes", "NestedTypes":
		v, ok := value.(DescriptorList)
		if ok {
			self.NestedTypes = v
			self.present["nestedTypes"] = true
			return nil
		} else {
			return fmt.Errorf("Field nestedTypes/NestedTypes: value %v(%T) couldn't be cast to type DescriptorList", value, value)
		}

	case "options", "Options":
		v, ok := value.(*MessageOptions)
		if ok {
			self.Options = v
			self.present["options"] = true
			return nil
		} else {
			return fmt.Errorf("Field options/Options: value %v(%T) couldn't be cast to type *MessageOptions", value, value)
		}

	}
}

func (self *Descriptor) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Descriptor", name)

	case "containingType", "ContainingType":
		if self.present != nil {
			if _, ok := self.present["containingType"]; ok {
				return self.ContainingType, nil
			}
		}
		return nil, fmt.Errorf("Field ContainingType no set on ContainingType %+v", self)

	case "file", "File":
		if self.present != nil {
			if _, ok := self.present["file"]; ok {
				return self.File, nil
			}
		}
		return nil, fmt.Errorf("Field File no set on File %+v", self)

	case "fullName", "FullName":
		if self.present != nil {
			if _, ok := self.present["fullName"]; ok {
				return self.FullName, nil
			}
		}
		return nil, fmt.Errorf("Field FullName no set on FullName %+v", self)

	case "index", "Index":
		if self.present != nil {
			if _, ok := self.present["index"]; ok {
				return self.Index, nil
			}
		}
		return nil, fmt.Errorf("Field Index no set on Index %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "nestedTypes", "NestedTypes":
		if self.present != nil {
			if _, ok := self.present["nestedTypes"]; ok {
				return self.NestedTypes, nil
			}
		}
		return nil, fmt.Errorf("Field NestedTypes no set on NestedTypes %+v", self)

	case "options", "Options":
		if self.present != nil {
			if _, ok := self.present["options"]; ok {
				return self.Options, nil
			}
		}
		return nil, fmt.Errorf("Field Options no set on Options %+v", self)

	}
}

func (self *Descriptor) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Descriptor", name)

	case "containingType", "ContainingType":
		self.present["containingType"] = false

	case "file", "File":
		self.present["file"] = false

	case "fullName", "FullName":
		self.present["fullName"] = false

	case "index", "Index":
		self.present["index"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nestedTypes", "NestedTypes":
		self.present["nestedTypes"] = false

	case "options", "Options":
		self.present["options"] = false

	}

	return nil
}

func (self *Descriptor) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DescriptorList []*Descriptor

func (self *DescriptorList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DescriptorList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DescriptorList cannot copy the values from %#v", other)
}

func (list *DescriptorList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DescriptorList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DescriptorList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
