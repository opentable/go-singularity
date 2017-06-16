package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type DockerOrBuilder struct {
	present map[string]bool

	Credential *Credential `json:"credential"`

	CredentialOrBuilder *CredentialOrBuilder `json:"credentialOrBuilder"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`
}

func (self *DockerOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *DockerOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerOrBuilder cannot absorb the values from %v", other)
}

func (self *DockerOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *DockerOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *DockerOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *DockerOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *DockerOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerOrBuilder", name)

	case "credential", "Credential":
		v, ok := value.(*Credential)
		if ok {
			self.Credential = v
			self.present["credential"] = true
			return nil
		} else {
			return fmt.Errorf("Field credential/Credential: value %v(%T) couldn't be cast to type *Credential", value, value)
		}

	case "credentialOrBuilder", "CredentialOrBuilder":
		v, ok := value.(*CredentialOrBuilder)
		if ok {
			self.CredentialOrBuilder = v
			self.present["credentialOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field credentialOrBuilder/CredentialOrBuilder: value %v(%T) couldn't be cast to type *CredentialOrBuilder", value, value)
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

	case "nameBytes", "NameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.NameBytes = v
			self.present["nameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field nameBytes/NameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	}
}

func (self *DockerOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on DockerOrBuilder", name)

	case "credential", "Credential":
		if self.present != nil {
			if _, ok := self.present["credential"]; ok {
				return self.Credential, nil
			}
		}
		return nil, fmt.Errorf("Field Credential no set on Credential %+v", self)

	case "credentialOrBuilder", "CredentialOrBuilder":
		if self.present != nil {
			if _, ok := self.present["credentialOrBuilder"]; ok {
				return self.CredentialOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field CredentialOrBuilder no set on CredentialOrBuilder %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "nameBytes", "NameBytes":
		if self.present != nil {
			if _, ok := self.present["nameBytes"]; ok {
				return self.NameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field NameBytes no set on NameBytes %+v", self)

	}
}

func (self *DockerOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DockerOrBuilder", name)

	case "credential", "Credential":
		self.present["credential"] = false

	case "credentialOrBuilder", "CredentialOrBuilder":
		self.present["credentialOrBuilder"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	}

	return nil
}

func (self *DockerOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DockerOrBuilderList []*DockerOrBuilder

func (self *DockerOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DockerOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DockerOrBuilder cannot absorb the values from %v", other)
}

func (list *DockerOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DockerOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DockerOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
