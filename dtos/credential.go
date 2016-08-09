package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Credential struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Credential `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Credential> `json:"parserForType"`

	Principal string `json:"principal,omitempty"`

	PrincipalBytes *ByteString `json:"principalBytes"`

	Secret string `json:"secret,omitempty"`

	SecretBytes *ByteString `json:"secretBytes"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Credential) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Credential) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Credential); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Credential cannot absorb the values from %v", other)
}

func (self *Credential) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Credential) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Credential) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Credential) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Credential) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Credential", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Credential)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Credential", value, value)
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

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
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

func (self *Credential) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Credential", name)

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

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *Credential) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Credential", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "principal", "Principal":
		self.present["principal"] = false

	case "principalBytes", "PrincipalBytes":
		self.present["principalBytes"] = false

	case "secret", "Secret":
		self.present["secret"] = false

	case "secretBytes", "SecretBytes":
		self.present["secretBytes"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Credential) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type CredentialList []*Credential

func (self *CredentialList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*CredentialList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Credential cannot absorb the values from %v", other)
}

func (list *CredentialList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *CredentialList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *CredentialList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
