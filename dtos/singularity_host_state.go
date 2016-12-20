package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityHostState struct {
	present map[string]bool

	DriverStatus string `json:"driverStatus,omitempty"`

	HostAddress string `json:"hostAddress,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Master bool `json:"master"`

	MesosMaster string `json:"mesosMaster,omitempty"`

	MillisSinceLastOffer int64 `json:"millisSinceLastOffer"`

	Uptime int64 `json:"uptime"`
}

func (self *SingularityHostState) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityHostState) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityHostState); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityHostState cannot copy the values from %#v", other)
}

func (self *SingularityHostState) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityHostState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityHostState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityHostState) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityHostState) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityHostState", name)

	case "driverStatus", "DriverStatus":
		v, ok := value.(string)
		if ok {
			self.DriverStatus = v
			self.present["driverStatus"] = true
			return nil
		} else {
			return fmt.Errorf("Field driverStatus/DriverStatus: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "hostAddress", "HostAddress":
		v, ok := value.(string)
		if ok {
			self.HostAddress = v
			self.present["hostAddress"] = true
			return nil
		} else {
			return fmt.Errorf("Field hostAddress/HostAddress: value %v(%T) couldn't be cast to type string", value, value)
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

	case "master", "Master":
		v, ok := value.(bool)
		if ok {
			self.Master = v
			self.present["master"] = true
			return nil
		} else {
			return fmt.Errorf("Field master/Master: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "mesosMaster", "MesosMaster":
		v, ok := value.(string)
		if ok {
			self.MesosMaster = v
			self.present["mesosMaster"] = true
			return nil
		} else {
			return fmt.Errorf("Field mesosMaster/MesosMaster: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		v, ok := value.(int64)
		if ok {
			self.MillisSinceLastOffer = v
			self.present["millisSinceLastOffer"] = true
			return nil
		} else {
			return fmt.Errorf("Field millisSinceLastOffer/MillisSinceLastOffer: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "uptime", "Uptime":
		v, ok := value.(int64)
		if ok {
			self.Uptime = v
			self.present["uptime"] = true
			return nil
		} else {
			return fmt.Errorf("Field uptime/Uptime: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityHostState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityHostState", name)

	case "driverStatus", "DriverStatus":
		if self.present != nil {
			if _, ok := self.present["driverStatus"]; ok {
				return self.DriverStatus, nil
			}
		}
		return nil, fmt.Errorf("Field DriverStatus no set on DriverStatus %+v", self)

	case "hostAddress", "HostAddress":
		if self.present != nil {
			if _, ok := self.present["hostAddress"]; ok {
				return self.HostAddress, nil
			}
		}
		return nil, fmt.Errorf("Field HostAddress no set on HostAddress %+v", self)

	case "hostname", "Hostname":
		if self.present != nil {
			if _, ok := self.present["hostname"]; ok {
				return self.Hostname, nil
			}
		}
		return nil, fmt.Errorf("Field Hostname no set on Hostname %+v", self)

	case "master", "Master":
		if self.present != nil {
			if _, ok := self.present["master"]; ok {
				return self.Master, nil
			}
		}
		return nil, fmt.Errorf("Field Master no set on Master %+v", self)

	case "mesosMaster", "MesosMaster":
		if self.present != nil {
			if _, ok := self.present["mesosMaster"]; ok {
				return self.MesosMaster, nil
			}
		}
		return nil, fmt.Errorf("Field MesosMaster no set on MesosMaster %+v", self)

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		if self.present != nil {
			if _, ok := self.present["millisSinceLastOffer"]; ok {
				return self.MillisSinceLastOffer, nil
			}
		}
		return nil, fmt.Errorf("Field MillisSinceLastOffer no set on MillisSinceLastOffer %+v", self)

	case "uptime", "Uptime":
		if self.present != nil {
			if _, ok := self.present["uptime"]; ok {
				return self.Uptime, nil
			}
		}
		return nil, fmt.Errorf("Field Uptime no set on Uptime %+v", self)

	}
}

func (self *SingularityHostState) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityHostState", name)

	case "driverStatus", "DriverStatus":
		self.present["driverStatus"] = false

	case "hostAddress", "HostAddress":
		self.present["hostAddress"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

	case "master", "Master":
		self.present["master"] = false

	case "mesosMaster", "MesosMaster":
		self.present["mesosMaster"] = false

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		self.present["millisSinceLastOffer"] = false

	case "uptime", "Uptime":
		self.present["uptime"] = false

	}

	return nil
}

func (self *SingularityHostState) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityHostStateList []*SingularityHostState

func (self *SingularityHostStateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityHostStateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityHostStateList cannot copy the values from %#v", other)
}

func (list *SingularityHostStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityHostStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityHostStateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
