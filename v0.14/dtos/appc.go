package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Appc struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Appc `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Id string `json:"id,omitempty"`

	IdBytes *ByteString `json:"idBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	Labels *Labels `json:"labels"`

	LabelsOrBuilder *LabelsOrBuilder `json:"labelsOrBuilder"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Image$Appc> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Appc) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Appc) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Appc); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Appc cannot absorb the values from %v", other)
}

func (self *Appc) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Appc) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Appc) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Appc) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Appc) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Appc", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Appc)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Appc", value, value)
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

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "idBytes", "IdBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.IdBytes = v
			self.present["idBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field idBytes/IdBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "labels", "Labels":
		v, ok := value.(*Labels)
		if ok {
			self.Labels = v
			self.present["labels"] = true
			return nil
		} else {
			return fmt.Errorf("Field labels/Labels: value %v(%T) couldn't be cast to type *Labels", value, value)
		}

	case "labelsOrBuilder", "LabelsOrBuilder":
		v, ok := value.(*LabelsOrBuilder)
		if ok {
			self.LabelsOrBuilder = v
			self.present["labelsOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field labelsOrBuilder/LabelsOrBuilder: value %v(%T) couldn't be cast to type *LabelsOrBuilder", value, value)
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

func (self *Appc) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Appc", name)

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

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "idBytes", "IdBytes":
		if self.present != nil {
			if _, ok := self.present["idBytes"]; ok {
				return self.IdBytes, nil
			}
		}
		return nil, fmt.Errorf("Field IdBytes no set on IdBytes %+v", self)

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

	case "labels", "Labels":
		if self.present != nil {
			if _, ok := self.present["labels"]; ok {
				return self.Labels, nil
			}
		}
		return nil, fmt.Errorf("Field Labels no set on Labels %+v", self)

	case "labelsOrBuilder", "LabelsOrBuilder":
		if self.present != nil {
			if _, ok := self.present["labelsOrBuilder"]; ok {
				return self.LabelsOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field LabelsOrBuilder no set on LabelsOrBuilder %+v", self)

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

func (self *Appc) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Appc", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "id", "Id":
		self.present["id"] = false

	case "idBytes", "IdBytes":
		self.present["idBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "labels", "Labels":
		self.present["labels"] = false

	case "labelsOrBuilder", "LabelsOrBuilder":
		self.present["labelsOrBuilder"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Appc) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type AppcList []*Appc

func (self *AppcList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AppcList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Appc cannot absorb the values from %v", other)
}

func (list *AppcList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *AppcList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *AppcList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
