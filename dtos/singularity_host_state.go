package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityHostState struct {
	present map[string]bool

	Master bool `json:"master"`

	DriverStatus string `json:"driverStatus,omitempty"`

	MillisSinceLastOffer int64 `json:"millisSinceLastOffer"`

	AvailableCachedCpus float64 `json:"availableCachedCpus"`

	MesosMaster string `json:"mesosMaster,omitempty"`

	MesosConnected bool `json:"mesosConnected"`

	Uptime int64 `json:"uptime"`

	OfferCacheSize int32 `json:"offerCacheSize"`

	AvailableCachedMemory float64 `json:"availableCachedMemory"`

	HostAddress string `json:"hostAddress,omitempty"`

	Hostname string `json:"hostname,omitempty"`
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

	case "master", "Master":
		v, ok := value.(bool)
		if ok {
			self.Master = v
			self.present["master"] = true
			return nil
		} else {
			return fmt.Errorf("Field master/Master: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "driverStatus", "DriverStatus":
		v, ok := value.(string)
		if ok {
			self.DriverStatus = v
			self.present["driverStatus"] = true
			return nil
		} else {
			return fmt.Errorf("Field driverStatus/DriverStatus: value %v(%T) couldn't be cast to type string", value, value)
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

	case "availableCachedCpus", "AvailableCachedCpus":
		v, ok := value.(float64)
		if ok {
			self.AvailableCachedCpus = v
			self.present["availableCachedCpus"] = true
			return nil
		} else {
			return fmt.Errorf("Field availableCachedCpus/AvailableCachedCpus: value %v(%T) couldn't be cast to type float64", value, value)
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

	case "mesosConnected", "MesosConnected":
		v, ok := value.(bool)
		if ok {
			self.MesosConnected = v
			self.present["mesosConnected"] = true
			return nil
		} else {
			return fmt.Errorf("Field mesosConnected/MesosConnected: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "offerCacheSize", "OfferCacheSize":
		v, ok := value.(int32)
		if ok {
			self.OfferCacheSize = v
			self.present["offerCacheSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field offerCacheSize/OfferCacheSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "availableCachedMemory", "AvailableCachedMemory":
		v, ok := value.(float64)
		if ok {
			self.AvailableCachedMemory = v
			self.present["availableCachedMemory"] = true
			return nil
		} else {
			return fmt.Errorf("Field availableCachedMemory/AvailableCachedMemory: value %v(%T) couldn't be cast to type float64", value, value)
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

	}
}

func (self *SingularityHostState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityHostState", name)

	case "master", "Master":
		if self.present != nil {
			if _, ok := self.present["master"]; ok {
				return self.Master, nil
			}
		}
		return nil, fmt.Errorf("Field Master no set on Master %+v", self)

	case "driverStatus", "DriverStatus":
		if self.present != nil {
			if _, ok := self.present["driverStatus"]; ok {
				return self.DriverStatus, nil
			}
		}
		return nil, fmt.Errorf("Field DriverStatus no set on DriverStatus %+v", self)

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		if self.present != nil {
			if _, ok := self.present["millisSinceLastOffer"]; ok {
				return self.MillisSinceLastOffer, nil
			}
		}
		return nil, fmt.Errorf("Field MillisSinceLastOffer no set on MillisSinceLastOffer %+v", self)

	case "availableCachedCpus", "AvailableCachedCpus":
		if self.present != nil {
			if _, ok := self.present["availableCachedCpus"]; ok {
				return self.AvailableCachedCpus, nil
			}
		}
		return nil, fmt.Errorf("Field AvailableCachedCpus no set on AvailableCachedCpus %+v", self)

	case "mesosMaster", "MesosMaster":
		if self.present != nil {
			if _, ok := self.present["mesosMaster"]; ok {
				return self.MesosMaster, nil
			}
		}
		return nil, fmt.Errorf("Field MesosMaster no set on MesosMaster %+v", self)

	case "mesosConnected", "MesosConnected":
		if self.present != nil {
			if _, ok := self.present["mesosConnected"]; ok {
				return self.MesosConnected, nil
			}
		}
		return nil, fmt.Errorf("Field MesosConnected no set on MesosConnected %+v", self)

	case "uptime", "Uptime":
		if self.present != nil {
			if _, ok := self.present["uptime"]; ok {
				return self.Uptime, nil
			}
		}
		return nil, fmt.Errorf("Field Uptime no set on Uptime %+v", self)

	case "offerCacheSize", "OfferCacheSize":
		if self.present != nil {
			if _, ok := self.present["offerCacheSize"]; ok {
				return self.OfferCacheSize, nil
			}
		}
		return nil, fmt.Errorf("Field OfferCacheSize no set on OfferCacheSize %+v", self)

	case "availableCachedMemory", "AvailableCachedMemory":
		if self.present != nil {
			if _, ok := self.present["availableCachedMemory"]; ok {
				return self.AvailableCachedMemory, nil
			}
		}
		return nil, fmt.Errorf("Field AvailableCachedMemory no set on AvailableCachedMemory %+v", self)

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

	}
}

func (self *SingularityHostState) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityHostState", name)

	case "master", "Master":
		self.present["master"] = false

	case "driverStatus", "DriverStatus":
		self.present["driverStatus"] = false

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		self.present["millisSinceLastOffer"] = false

	case "availableCachedCpus", "AvailableCachedCpus":
		self.present["availableCachedCpus"] = false

	case "mesosMaster", "MesosMaster":
		self.present["mesosMaster"] = false

	case "mesosConnected", "MesosConnected":
		self.present["mesosConnected"] = false

	case "uptime", "Uptime":
		self.present["uptime"] = false

	case "offerCacheSize", "OfferCacheSize":
		self.present["offerCacheSize"] = false

	case "availableCachedMemory", "AvailableCachedMemory":
		self.present["availableCachedMemory"] = false

	case "hostAddress", "HostAddress":
		self.present["hostAddress"] = false

	case "hostname", "Hostname":
		self.present["hostname"] = false

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
