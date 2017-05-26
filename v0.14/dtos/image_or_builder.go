package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type ImageOrBuilderType string

const (
	ImageOrBuilderTypeAPPC   ImageOrBuilderType = "APPC"
	ImageOrBuilderTypeDOCKER ImageOrBuilderType = "DOCKER"
)

type ImageOrBuilder struct {
	present map[string]bool

	Appc *Appc `json:"appc"`

	AppcOrBuilder *AppcOrBuilder `json:"appcOrBuilder"`

	Docker *Docker `json:"docker"`

	DockerOrBuilder *DockerOrBuilder `json:"dockerOrBuilder"`

	Type ImageOrBuilderType `json:"type"`
}

func (self *ImageOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *ImageOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ImageOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ImageOrBuilder cannot absorb the values from %v", other)
}

func (self *ImageOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *ImageOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *ImageOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *ImageOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *ImageOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ImageOrBuilder", name)

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

	case "type", "Type":
		v, ok := value.(ImageOrBuilderType)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type ImageOrBuilderType", value, value)
		}

	}
}

func (self *ImageOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on ImageOrBuilder", name)

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

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	}
}

func (self *ImageOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on ImageOrBuilder", name)

	case "appc", "Appc":
		self.present["appc"] = false

	case "appcOrBuilder", "AppcOrBuilder":
		self.present["appcOrBuilder"] = false

	case "docker", "Docker":
		self.present["docker"] = false

	case "dockerOrBuilder", "DockerOrBuilder":
		self.present["dockerOrBuilder"] = false

	case "type", "Type":
		self.present["type"] = false

	}

	return nil
}

func (self *ImageOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ImageOrBuilderList []*ImageOrBuilder

func (self *ImageOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ImageOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ImageOrBuilder cannot absorb the values from %v", other)
}

func (list *ImageOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ImageOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ImageOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
