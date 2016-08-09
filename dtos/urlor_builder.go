package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type URLOrBuilder struct {
	present map[string]bool

	Address *Address `json:"address"`

	AddressOrBuilder *AddressOrBuilder `json:"addressOrBuilder"`

	Fragment string `json:"fragment,omitempty"`

	FragmentBytes *ByteString `json:"fragmentBytes"`

	Path string `json:"path,omitempty"`

	PathBytes *ByteString `json:"pathBytes"`

	QueryCount int32 `json:"queryCount"`

	// QueryList *List[Parameter] `json:"queryList"`

	// QueryOrBuilderList *List[? extends org.apache.mesos.Protos$ParameterOrBuilder] `json:"queryOrBuilderList"`

	Scheme string `json:"scheme,omitempty"`

	SchemeBytes *ByteString `json:"schemeBytes"`
}

func (self *URLOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *URLOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*URLOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A URLOrBuilder cannot absorb the values from %v", other)
}

func (self *URLOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *URLOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *URLOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *URLOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *URLOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on URLOrBuilder", name)

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

	}
}

func (self *URLOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on URLOrBuilder", name)

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

	}
}

func (self *URLOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on URLOrBuilder", name)

	case "address", "Address":
		self.present["address"] = false

	case "addressOrBuilder", "AddressOrBuilder":
		self.present["addressOrBuilder"] = false

	case "fragment", "Fragment":
		self.present["fragment"] = false

	case "fragmentBytes", "FragmentBytes":
		self.present["fragmentBytes"] = false

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

	}

	return nil
}

func (self *URLOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type URLOrBuilderList []*URLOrBuilder

func (self *URLOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*URLOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A URLOrBuilder cannot absorb the values from %v", other)
}

func (list *URLOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *URLOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *URLOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
