package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type AddressOrBuilder struct {
	present map[string]bool

	Hostname string `json:"hostname,omitempty"`

	HostnameBytes *ByteString `json:"hostnameBytes"`

	Ip string `json:"ip,omitempty"`

	IpBytes *ByteString `json:"ipBytes"`

	Port int32 `json:"port"`
}

func (self *AddressOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *AddressOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AddressOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A AddressOrBuilder cannot absorb the values from %v", other)
}

func (self *AddressOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *AddressOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *AddressOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *AddressOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *AddressOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on AddressOrBuilder", name)

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

	}
}

func (self *AddressOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on AddressOrBuilder", name)

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

	}
}

func (self *AddressOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on AddressOrBuilder", name)

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "hostnameBytes", "HostnameBytes":
		self.present["hostnameBytes"] = false

	case "ip", "Ip":
		self.present["ip"] = false

	case "ipBytes", "IpBytes":
		self.present["ipBytes"] = false

	case "port", "Port":
		self.present["port"] = false

	}

	return nil
}

func (self *AddressOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type AddressOrBuilderList []*AddressOrBuilder

func (self *AddressOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*AddressOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A AddressOrBuilder cannot absorb the values from %v", other)
}

func (list *AddressOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *AddressOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *AddressOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
