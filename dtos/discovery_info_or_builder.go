package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type DiscoveryInfoOrBuilderVisibility string

const (
	DiscoveryInfoOrBuilderVisibilityFRAMEWORK DiscoveryInfoOrBuilderVisibility = "FRAMEWORK"
	DiscoveryInfoOrBuilderVisibilityCLUSTER   DiscoveryInfoOrBuilderVisibility = "CLUSTER"
	DiscoveryInfoOrBuilderVisibilityEXTERNAL  DiscoveryInfoOrBuilderVisibility = "EXTERNAL"
)

type DiscoveryInfoOrBuilder struct {
	present map[string]bool

	Environment string `json:"environment,omitempty"`

	EnvironmentBytes *ByteString `json:"environmentBytes"`

	Labels *Labels `json:"labels"`

	LabelsOrBuilder *LabelsOrBuilder `json:"labelsOrBuilder"`

	Location string `json:"location,omitempty"`

	LocationBytes *ByteString `json:"locationBytes"`

	Name string `json:"name,omitempty"`

	NameBytes *ByteString `json:"nameBytes"`

	Ports *Ports `json:"ports"`

	PortsOrBuilder *PortsOrBuilder `json:"portsOrBuilder"`

	Version string `json:"version,omitempty"`

	VersionBytes *ByteString `json:"versionBytes"`

	Visibility DiscoveryInfoOrBuilderVisibility `json:"visibility"`
}

func (self *DiscoveryInfoOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *DiscoveryInfoOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DiscoveryInfoOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DiscoveryInfoOrBuilder cannot copy the values from %#v", other)
}

func (self *DiscoveryInfoOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *DiscoveryInfoOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *DiscoveryInfoOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *DiscoveryInfoOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *DiscoveryInfoOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DiscoveryInfoOrBuilder", name)

	case "environment", "Environment":
		v, ok := value.(string)
		if ok {
			self.Environment = v
			self.present["environment"] = true
			return nil
		} else {
			return fmt.Errorf("Field environment/Environment: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "environmentBytes", "EnvironmentBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.EnvironmentBytes = v
			self.present["environmentBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field environmentBytes/EnvironmentBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "location", "Location":
		v, ok := value.(string)
		if ok {
			self.Location = v
			self.present["location"] = true
			return nil
		} else {
			return fmt.Errorf("Field location/Location: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "locationBytes", "LocationBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.LocationBytes = v
			self.present["locationBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field locationBytes/LocationBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
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

	case "ports", "Ports":
		v, ok := value.(*Ports)
		if ok {
			self.Ports = v
			self.present["ports"] = true
			return nil
		} else {
			return fmt.Errorf("Field ports/Ports: value %v(%T) couldn't be cast to type *Ports", value, value)
		}

	case "portsOrBuilder", "PortsOrBuilder":
		v, ok := value.(*PortsOrBuilder)
		if ok {
			self.PortsOrBuilder = v
			self.present["portsOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field portsOrBuilder/PortsOrBuilder: value %v(%T) couldn't be cast to type *PortsOrBuilder", value, value)
		}

	case "version", "Version":
		v, ok := value.(string)
		if ok {
			self.Version = v
			self.present["version"] = true
			return nil
		} else {
			return fmt.Errorf("Field version/Version: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "versionBytes", "VersionBytes":
		v, ok := value.(*ByteString)
		if ok {
			self.VersionBytes = v
			self.present["versionBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field versionBytes/VersionBytes: value %v(%T) couldn't be cast to type *ByteString", value, value)
		}

	case "visibility", "Visibility":
		v, ok := value.(DiscoveryInfoOrBuilderVisibility)
		if ok {
			self.Visibility = v
			self.present["visibility"] = true
			return nil
		} else {
			return fmt.Errorf("Field visibility/Visibility: value %v(%T) couldn't be cast to type DiscoveryInfoOrBuilderVisibility", value, value)
		}

	}
}

func (self *DiscoveryInfoOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on DiscoveryInfoOrBuilder", name)

	case "environment", "Environment":
		if self.present != nil {
			if _, ok := self.present["environment"]; ok {
				return self.Environment, nil
			}
		}
		return nil, fmt.Errorf("Field Environment no set on Environment %+v", self)

	case "environmentBytes", "EnvironmentBytes":
		if self.present != nil {
			if _, ok := self.present["environmentBytes"]; ok {
				return self.EnvironmentBytes, nil
			}
		}
		return nil, fmt.Errorf("Field EnvironmentBytes no set on EnvironmentBytes %+v", self)

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

	case "location", "Location":
		if self.present != nil {
			if _, ok := self.present["location"]; ok {
				return self.Location, nil
			}
		}
		return nil, fmt.Errorf("Field Location no set on Location %+v", self)

	case "locationBytes", "LocationBytes":
		if self.present != nil {
			if _, ok := self.present["locationBytes"]; ok {
				return self.LocationBytes, nil
			}
		}
		return nil, fmt.Errorf("Field LocationBytes no set on LocationBytes %+v", self)

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

	case "ports", "Ports":
		if self.present != nil {
			if _, ok := self.present["ports"]; ok {
				return self.Ports, nil
			}
		}
		return nil, fmt.Errorf("Field Ports no set on Ports %+v", self)

	case "portsOrBuilder", "PortsOrBuilder":
		if self.present != nil {
			if _, ok := self.present["portsOrBuilder"]; ok {
				return self.PortsOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field PortsOrBuilder no set on PortsOrBuilder %+v", self)

	case "version", "Version":
		if self.present != nil {
			if _, ok := self.present["version"]; ok {
				return self.Version, nil
			}
		}
		return nil, fmt.Errorf("Field Version no set on Version %+v", self)

	case "versionBytes", "VersionBytes":
		if self.present != nil {
			if _, ok := self.present["versionBytes"]; ok {
				return self.VersionBytes, nil
			}
		}
		return nil, fmt.Errorf("Field VersionBytes no set on VersionBytes %+v", self)

	case "visibility", "Visibility":
		if self.present != nil {
			if _, ok := self.present["visibility"]; ok {
				return self.Visibility, nil
			}
		}
		return nil, fmt.Errorf("Field Visibility no set on Visibility %+v", self)

	}
}

func (self *DiscoveryInfoOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on DiscoveryInfoOrBuilder", name)

	case "environment", "Environment":
		self.present["environment"] = false

	case "environmentBytes", "EnvironmentBytes":
		self.present["environmentBytes"] = false

	case "labels", "Labels":
		self.present["labels"] = false

	case "labelsOrBuilder", "LabelsOrBuilder":
		self.present["labelsOrBuilder"] = false

	case "location", "Location":
		self.present["location"] = false

	case "locationBytes", "LocationBytes":
		self.present["locationBytes"] = false

	case "name", "Name":
		self.present["name"] = false

	case "nameBytes", "NameBytes":
		self.present["nameBytes"] = false

	case "ports", "Ports":
		self.present["ports"] = false

	case "portsOrBuilder", "PortsOrBuilder":
		self.present["portsOrBuilder"] = false

	case "version", "Version":
		self.present["version"] = false

	case "versionBytes", "VersionBytes":
		self.present["versionBytes"] = false

	case "visibility", "Visibility":
		self.present["visibility"] = false

	}

	return nil
}

func (self *DiscoveryInfoOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type DiscoveryInfoOrBuilderList []*DiscoveryInfoOrBuilder

func (self *DiscoveryInfoOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*DiscoveryInfoOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A DiscoveryInfoOrBuilderList cannot copy the values from %#v", other)
}

func (list *DiscoveryInfoOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *DiscoveryInfoOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *DiscoveryInfoOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
