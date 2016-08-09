package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type CredentialOrBuilder struct {
	present map[string]bool

	Principal string `json:"principal,omitempty"`

	PrincipalBytes *ByteString `json:"principalBytes"`

	Secret string `json:"secret,omitempty"`

	SecretBytes *ByteString `json:"secretBytes"`
}

func (self *CredentialOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *CredentialOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*CredentialOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A CredentialOrBuilder cannot absorb the values from %v", other)
}

func (self *CredentialOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *CredentialOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *CredentialOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *CredentialOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *CredentialOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on CredentialOrBuilder", name)

	case "principal", "Principal":
		v, ok := value.(string)
		if ok {
			self.Principal = v
			self.present["principal"] = true
			return nil
		} else {
			return fmt.Errorf("Field principal/Principal: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "principalBytes", "PrincipalBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.PrincipalBytes = v
			self.present["principalBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field principalBytes/PrincipalBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "secret", "Secret":
		v, ok := value.(string)
		if ok {
			self.Secret = v
			self.present["secret"] = true
			return nil
		} else {
			return fmt.Errorf("Field secret/Secret: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "secretBytes", "SecretBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.SecretBytes = v
			self.present["secretBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field secretBytes/SecretBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	}
}

func (self *CredentialOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on CredentialOrBuilder", name)

	case "principal", "Principal":
		if self.present != nil {
			if _, ok := self.present["principal"]; ok {
				return self.Principal, nil
			}
		}
		return nil, fmt.Errorf("Field Principal no set on Principal %+v", self)

	case "principalBytes", "PrincipalBytes":
		if self.present != nil {
			if _, ok := self.present["principalBytes"]; ok {
				return self.PrincipalBytes, nil
			}
		}
		return nil, fmt.Errorf("Field PrincipalBytes no set on PrincipalBytes %+v", self)

	case "secret", "Secret":
		if self.present != nil {
			if _, ok := self.present["secret"]; ok {
				return self.Secret, nil
			}
		}
		return nil, fmt.Errorf("Field Secret no set on Secret %+v", self)

	case "secretBytes", "SecretBytes":
		if self.present != nil {
			if _, ok := self.present["secretBytes"]; ok {
				return self.SecretBytes, nil
			}
		}
		return nil, fmt.Errorf("Field SecretBytes no set on SecretBytes %+v", self)

	}
}

func (self *CredentialOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on CredentialOrBuilder", name)

	case "principal", "Principal":
		self.present["principal"] = false

	case "principalBytes", "PrincipalBytes":
		self.present["principalBytes"] = false

	case "secret", "Secret":
		self.present["secret"] = false

	case "secretBytes", "SecretBytes":
		self.present["secretBytes"] = false

	}

	return nil
}

func (self *CredentialOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type CredentialOrBuilderList []*CredentialOrBuilder

func (self *CredentialOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*CredentialOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A CredentialOrBuilder cannot absorb the values from %v", other)
}

func (list *CredentialOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *CredentialOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *CredentialOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
