package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Address struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Address `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Hostname string `json:"hostname,omitempty"`

	HostnameBytes *ByteString `json:"hostnameBytes"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	Ip string `json:"ip,omitempty"`

	IpBytes *ByteString `json:"ipBytes"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Address> `json:"parserForType"`

	Port int32 `json:"port"`

	SerializedSize int32 `json:"serializedSize"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Address) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Address) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Address); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Address cannot absorb the values from %v", other)
}

func (self *Address) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Address) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Address) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Address) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Address) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Address", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Address)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Address", value, value)
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

	case "hostname", "Hostname":
		v, ok := value.(string)
		if ok {
			self.Hostname = v
			self.present["hostname"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostname/Hostname: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "hostnameBytes", "HostnameBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.HostnameBytes = v
			self.present["hostnameBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostnameBytes/HostnameBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "ip", "Ip":
		v, ok := value.(string)
		if ok {
			self.Ip = v
			self.present["ip"] = true
			return nil
		} else {
			return fmt.Errorf("Field ip/Ip: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "ipBytes", "IpBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.IpBytes = v
			self.present["ipBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field ipBytes/IpBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

func (self *Address) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Address", name)

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

	case "hostname", "Hostname":
		if self.present != nil {
			if _, ok := self.present["hostname"]; ok {
				return self.Hostname, nil
			}
		}
		return nil, fmt.Errorf("Field Hostname no set on Hostname %+v", self)

	case "hostnameBytes", "HostnameBytes":
		if self.present != nil {
			if _, ok := self.present["hostnameBytes"]; ok {
				return self.HostnameBytes, nil
			}
		}
		return nil, fmt.Errorf("Field HostnameBytes no set on HostnameBytes %+v", self)

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

	case "ip", "Ip":
		if self.present != nil {
			if _, ok := self.present["ip"]; ok {
				return self.Ip, nil
			}
		}
		return nil, fmt.Errorf("Field Ip no set on Ip %+v", self)

	case "ipBytes", "IpBytes":
		if self.present != nil {
			if _, ok := self.present["ipBytes"]; ok {
				return self.IpBytes, nil
			}
		}
		return nil, fmt.Errorf("Field IpBytes no set on IpBytes %+v", self)

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

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *Address) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Address", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "hostnameBytes", "HostnameBytes":
		self.present["hostnameBytes"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "ip", "Ip":
		self.present["ip"] = false

	case "ipBytes", "IpBytes":
		self.present["ipBytes"] = false

	case "port", "Port":
		self.present["port"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Address) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type AddressList []*Address

func (self *AddressList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AddressList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Address cannot absorb the values from %v", other)
}

func (list *AddressList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *AddressList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *AddressList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
