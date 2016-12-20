package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type HTTP struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *HTTP `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$HealthCheck$HTTP> `json:"parserForType"`

	Path string `json:"path,omitempty"`

	PathBytes *ByteString `json:"pathBytes"`

	Port int32 `json:"port"`

	SerializedSize int32 `json:"serializedSize"`

	StatusesCount int32 `json:"statusesCount"`

	StatusesList []int32 `json:"statusesList"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *HTTP) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *HTTP) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HTTP); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HTTP cannot copy the values from %#v", other)
}

func (self *HTTP) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *HTTP) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HTTP) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HTTP) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *HTTP) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HTTP", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*HTTP)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *HTTP", value, value)
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

	case "path", "Path":
		v, ok := value.(string)
		if ok {
			self.Path = v
			self.present["path"] = true
			return nil
		} else {
			return fmt.Errorf("Field path/Path: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "pathBytes", "PathBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.PathBytes = v
			self.present["pathBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field pathBytes/PathBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "port", "Port":
		v, ok := value.(int32)
		if ok {
			self.Port = v
			self.present["port"] = true
			return nil
		} else {
			return fmt.Errorf("Field port/Port: value %v(%T) couldn't be cast to type int32", value, value)
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

	case "statusesCount", "StatusesCount":
		v, ok := value.(int32)
		if ok {
			self.StatusesCount = v
			self.present["statusesCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusesCount/StatusesCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "statusesList", "StatusesList":
		v, ok := value.([]int32)
		if ok {
			self.StatusesList = v
			self.present["statusesList"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusesList/StatusesList: value %v(%T) couldn't be cast to type []int32", value, value)
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

func (self *HTTP) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HTTP", name)

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

	case "path", "Path":
		if self.present != nil {
			if _, ok := self.present["path"]; ok {
				return self.Path, nil
			}
		}
		return nil, fmt.Errorf("Field Path no set on Path %+v", self)

	case "pathBytes", "PathBytes":
		if self.present != nil {
			if _, ok := self.present["pathBytes"]; ok {
				return self.PathBytes, nil
			}
		}
		return nil, fmt.Errorf("Field PathBytes no set on PathBytes %+v", self)

	case "port", "Port":
		if self.present != nil {
			if _, ok := self.present["port"]; ok {
				return self.Port, nil
			}
		}
		return nil, fmt.Errorf("Field Port no set on Port %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "statusesCount", "StatusesCount":
		if self.present != nil {
			if _, ok := self.present["statusesCount"]; ok {
				return self.StatusesCount, nil
			}
		}
		return nil, fmt.Errorf("Field StatusesCount no set on StatusesCount %+v", self)

	case "statusesList", "StatusesList":
		if self.present != nil {
			if _, ok := self.present["statusesList"]; ok {
				return self.StatusesList, nil
			}
		}
		return nil, fmt.Errorf("Field StatusesList no set on StatusesList %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *HTTP) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HTTP", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "path", "Path":
		self.present["path"] = false

	case "pathBytes", "PathBytes":
		self.present["pathBytes"] = false

	case "port", "Port":
		self.present["port"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "statusesCount", "StatusesCount":
		self.present["statusesCount"] = false

	case "statusesList", "StatusesList":
		self.present["statusesList"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *HTTP) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type HTTPList []*HTTP

func (self *HTTPList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HTTPList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HTTPList cannot copy the values from %#v", other)
}

func (list *HTTPList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *HTTPList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HTTPList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
