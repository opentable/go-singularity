package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type URL struct {
	present map[string]bool

	Address *Address `json:"address"`

	AddressOrBuilder *AddressOrBuilder `json:"addressOrBuilder"`

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *URL `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Fragment string `json:"fragment,omitempty"`

	FragmentBytes *ByteString `json:"fragmentBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$URL> `json:"parserForType"`

	Path string `json:"path,omitempty"`

	PathBytes *ByteString `json:"pathBytes"`

	QueryCount int32 `json:"queryCount"`

	// QueryList *List[Parameter] `json:"queryList"`

	// QueryOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"queryOrBuilderList"`

	Scheme string `json:"scheme,omitempty"`

	SchemeBytes *ByteString `json:"schemeBytes"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *URL) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *URL) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*URL); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A URL cannot absorb the values from %v", other)
}

func (self *URL) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *URL) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *URL) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *URL) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *URL) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on URL", name)

	case "address", "Address":
		v, ok := value.(*Address)
		if ok {
			self.Address = v
			self.present["address"] = true
			return nil
		} else {
			return fmt.Errorf("Field address/Address: value %v(%T) couldn't be cast to type *Address", value, value)
		}

	case "addressOrBuilder", "AddressOrBuilder":
		v, ok := value.(*AddressOrBuilder)
		if ok {
			self.AddressOrBuilder = v
			self.present["addressOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field addressOrBuilder/AddressOrBuilder: value %v(%T) couldn't be cast to type *AddressOrBuilder", value, value)
		}

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*URL)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *URL", value, value)
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

	case "fragment", "Fragment":
		v, ok := value.(string)
		if ok {
			self.Fragment = v
			self.present["fragment"] = true
			return nil
		} else {
			return fmt.Errorf("Field fragment/Fragment: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "fragmentBytes", "FragmentBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.FragmentBytes = v
			self.present["fragmentBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field fragmentBytes/FragmentBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "queryCount", "QueryCount":
		v, ok := value.(int32)
		if ok {
			self.QueryCount = v
			self.present["queryCount"] = true
			return nil
		} else {
			return fmt.Errorf("Field queryCount/QueryCount: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "scheme", "Scheme":
		v, ok := value.(string)
		if ok {
			self.Scheme = v
			self.present["scheme"] = true
			return nil
		} else {
			return fmt.Errorf("Field scheme/Scheme: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "schemeBytes", "SchemeBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.SchemeBytes = v
			self.present["schemeBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field schemeBytes/SchemeBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

func (self *URL) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on URL", name)

	case "address", "Address":
		if self.present != nil {
			if _, ok := self.present["address"]; ok {
				return self.Address, nil
			}
		}
		return nil, fmt.Errorf("Field Address no set on Address %+v", self)

	case "addressOrBuilder", "AddressOrBuilder":
		if self.present != nil {
			if _, ok := self.present["addressOrBuilder"]; ok {
				return self.AddressOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field AddressOrBuilder no set on AddressOrBuilder %+v", self)

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

	case "fragment", "Fragment":
		if self.present != nil {
			if _, ok := self.present["fragment"]; ok {
				return self.Fragment, nil
			}
		}
		return nil, fmt.Errorf("Field Fragment no set on Fragment %+v", self)

	case "fragmentBytes", "FragmentBytes":
		if self.present != nil {
			if _, ok := self.present["fragmentBytes"]; ok {
				return self.FragmentBytes, nil
			}
		}
		return nil, fmt.Errorf("Field FragmentBytes no set on FragmentBytes %+v", self)

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

	case "queryCount", "QueryCount":
		if self.present != nil {
			if _, ok := self.present["queryCount"]; ok {
				return self.QueryCount, nil
			}
		}
		return nil, fmt.Errorf("Field QueryCount no set on QueryCount %+v", self)

	case "scheme", "Scheme":
		if self.present != nil {
			if _, ok := self.present["scheme"]; ok {
				return self.Scheme, nil
			}
		}
		return nil, fmt.Errorf("Field Scheme no set on Scheme %+v", self)

	case "schemeBytes", "SchemeBytes":
		if self.present != nil {
			if _, ok := self.present["schemeBytes"]; ok {
				return self.SchemeBytes, nil
			}
		}
		return nil, fmt.Errorf("Field SchemeBytes no set on SchemeBytes %+v", self)

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

func (self *URL) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on URL", name)

	case "address", "Address":
		self.present["address"] = false

	case "addressOrBuilder", "AddressOrBuilder":
		self.present["addressOrBuilder"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "fragment", "Fragment":
		self.present["fragment"] = false

	case "fragmentBytes", "FragmentBytes":
		self.present["fragmentBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "path", "Path":
		self.present["path"] = false

	case "pathBytes", "PathBytes":
		self.present["pathBytes"] = false

	case "queryCount", "QueryCount":
		self.present["queryCount"] = false

	case "scheme", "Scheme":
		self.present["scheme"] = false

	case "schemeBytes", "SchemeBytes":
		self.present["schemeBytes"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *URL) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type URLList []*URL

func (self *URLList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*URLList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A URL cannot absorb the values from %v", other)
}

func (list *URLList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *URLList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *URLList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
