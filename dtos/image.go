package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ImageType string

const (
	ImageTypeAPPC   ImageType = "APPC"
	ImageTypeDOCKER ImageType = "DOCKER"
)

type Image struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	Appc *Appc `json:"appc"`

	AppcOrBuilder *AppcOrBuilder `json:"appcOrBuilder"`

	DefaultInstanceForType *Image `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Docker *Docker `json:"docker"`

	DockerOrBuilder *DockerOrBuilder `json:"dockerOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Image> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	Type ImageType `json:"type"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Image) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Image) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Image); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Image cannot absorb the values from %v", other)
}

func (self *Image) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Image) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Image) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Image) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Image) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Image", name)

	case "appc", "Appc":
		v, ok := value.(*Appc)
		if ok {
			self.Appc = v
			self.present["appc"] = true
			return nil
		} else {
			return fmt.Errorf("Field appc/Appc: value %v(%T) couldn't be cast to type *Appc", value, value)
		}

	case "appcOrBuilder", "AppcOrBuilder":
		v, ok := value.(*AppcOrBuilder)
		if ok {
			self.AppcOrBuilder = v
			self.present["appcOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field appcOrBuilder/AppcOrBuilder: value %v(%T) couldn't be cast to type *AppcOrBuilder", value, value)
		}

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Image)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Image", value, value)
		}

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
		}

	case "docker", "Docker":
		v, ok := value.(*Docker)
		if ok {
			self.Docker = v
			self.present["docker"] = true
			return nil
		} else {
			return fmt.Errorf("Field docker/Docker: value %v(%T) couldn't be cast to type *Docker", value, value)
		}

	case "dockerOrBuilder", "DockerOrBuilder":
		v, ok := value.(*DockerOrBuilder)
		if ok {
			self.DockerOrBuilder = v
			self.present["dockerOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field dockerOrBuilder/DockerOrBuilder: value %v(%T) couldn't be cast to type *DockerOrBuilder", value, value)
		}

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "type", "Type":
		v, ok := value.(ImageType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type ImageType", value, value)
		}

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	}
}

func (self *Image) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Image", name)

	case "appc", "Appc":
		if self.present != nil {
			if _, ok := self.present["appc"]; ok {
				return self.Appc, nil
			}
		}
		return nil, fmt.Errorf("Field Appc no set on Appc %+v", self)

	case "appcOrBuilder", "AppcOrBuilder":
		if self.present != nil {
			if _, ok := self.present["appcOrBuilder"]; ok {
				return self.AppcOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field AppcOrBuilder no set on AppcOrBuilder %+v", self)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "docker", "Docker":
		if self.present != nil {
			if _, ok := self.present["docker"]; ok {
				return self.Docker, nil
			}
		}
		return nil, fmt.Errorf("Field Docker no set on Docker %+v", self)

	case "dockerOrBuilder", "DockerOrBuilder":
		if self.present != nil {
			if _, ok := self.present["dockerOrBuilder"]; ok {
				return self.DockerOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field DockerOrBuilder no set on DockerOrBuilder %+v", self)

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *Image) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Image", name)

	case "appc", "Appc":
		self.present["appc"] = false

	case "appcOrBuilder", "AppcOrBuilder":
		self.present["appcOrBuilder"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "docker", "Docker":
		self.present["docker"] = false

	case "dockerOrBuilder", "DockerOrBuilder":
		self.present["dockerOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "type", "Type":
		self.present["type"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Image) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ImageList []*Image

func (self *ImageList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ImageList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Image cannot absorb the values from %v", other)
}

func (list *ImageList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ImageList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ImageList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
