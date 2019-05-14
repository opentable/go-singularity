package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityHostState struct {
	MesosMaster           *string  `json:"mesosMaster,omitempty"`
	DriverStatus          *string  `json:"driverStatus,omitempty"`
	AvailableCachedCpus   *float64 `json:"availableCachedCpus,omitempty"`
	MillisSinceLastOffer  *int64   `json:"millisSinceLastOffer,omitempty"`
	OfferCacheSize        *int32   `json:"offerCacheSize,omitempty"`
	AvailableCachedMemory *float64 `json:"availableCachedMemory,omitempty"`
	HostAddress           *string  `json:"hostAddress,omitempty"`
	Hostname              *string  `json:"hostname,omitempty"`
	MesosConnected        *bool    `json:"mesosConnected,omitempty"`
	Master                *bool    `json:"master,omitempty"`
	Uptime                *int64   `json:"uptime,omitempty"`
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

func (self *SingularityHostState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityHostState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityHostState) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityHostState", name)

	case "mesosMaster", "MesosMaster":
		v, ok := value.(string)
		if ok {
			self.MesosMaster = &v
			return nil
		}
		return fmt.Errorf("Field mesosMaster/MesosMaster: value %v(%T) couldn't be cast to type string", value, value)

	case "driverStatus", "DriverStatus":
		v, ok := value.(string)
		if ok {
			self.DriverStatus = &v
			return nil
		}
		return fmt.Errorf("Field driverStatus/DriverStatus: value %v(%T) couldn't be cast to type string", value, value)

	case "availableCachedCpus", "AvailableCachedCpus":
		v, ok := value.(float64)
		if ok {
			self.AvailableCachedCpus = &v
			return nil
		}
		return fmt.Errorf("Field availableCachedCpus/AvailableCachedCpus: value %v(%T) couldn't be cast to type float64", value, value)

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		v, ok := value.(int64)
		if ok {
			self.MillisSinceLastOffer = &v
			return nil
		}
		return fmt.Errorf("Field millisSinceLastOffer/MillisSinceLastOffer: value %v(%T) couldn't be cast to type int64", value, value)

	case "offerCacheSize", "OfferCacheSize":
		v, ok := value.(int32)
		if ok {
			self.OfferCacheSize = &v
			return nil
		}
		return fmt.Errorf("Field offerCacheSize/OfferCacheSize: value %v(%T) couldn't be cast to type int32", value, value)

	case "availableCachedMemory", "AvailableCachedMemory":
		v, ok := value.(float64)
		if ok {
			self.AvailableCachedMemory = &v
			return nil
		}
		return fmt.Errorf("Field availableCachedMemory/AvailableCachedMemory: value %v(%T) couldn't be cast to type float64", value, value)

	case "hostAddress", "HostAddress":
		v, ok := value.(string)
		if ok {
			self.HostAddress = &v
			return nil
		}
		return fmt.Errorf("Field hostAddress/HostAddress: value %v(%T) couldn't be cast to type string", value, value)

	case "hostname", "Hostname":
		v, ok := value.(string)
		if ok {
			self.Hostname = &v
			return nil
		}
		return fmt.Errorf("Field hostname/Hostname: value %v(%T) couldn't be cast to type string", value, value)

	case "mesosConnected", "MesosConnected":
		v, ok := value.(bool)
		if ok {
			self.MesosConnected = &v
			return nil
		}
		return fmt.Errorf("Field mesosConnected/MesosConnected: value %v(%T) couldn't be cast to type bool", value, value)

	case "master", "Master":
		v, ok := value.(bool)
		if ok {
			self.Master = &v
			return nil
		}
		return fmt.Errorf("Field master/Master: value %v(%T) couldn't be cast to type bool", value, value)

	case "uptime", "Uptime":
		v, ok := value.(int64)
		if ok {
			self.Uptime = &v
			return nil
		}
		return fmt.Errorf("Field uptime/Uptime: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityHostState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityHostState", name)

	case "mesosMaster", "MesosMaster":
		return *self.MesosMaster, nil
		return nil, fmt.Errorf("Field MesosMaster no set on MesosMaster %+v", self)

	case "driverStatus", "DriverStatus":
		return *self.DriverStatus, nil
		return nil, fmt.Errorf("Field DriverStatus no set on DriverStatus %+v", self)

	case "availableCachedCpus", "AvailableCachedCpus":
		return *self.AvailableCachedCpus, nil
		return nil, fmt.Errorf("Field AvailableCachedCpus no set on AvailableCachedCpus %+v", self)

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		return *self.MillisSinceLastOffer, nil
		return nil, fmt.Errorf("Field MillisSinceLastOffer no set on MillisSinceLastOffer %+v", self)

	case "offerCacheSize", "OfferCacheSize":
		return *self.OfferCacheSize, nil
		return nil, fmt.Errorf("Field OfferCacheSize no set on OfferCacheSize %+v", self)

	case "availableCachedMemory", "AvailableCachedMemory":
		return *self.AvailableCachedMemory, nil
		return nil, fmt.Errorf("Field AvailableCachedMemory no set on AvailableCachedMemory %+v", self)

	case "hostAddress", "HostAddress":
		return *self.HostAddress, nil
		return nil, fmt.Errorf("Field HostAddress no set on HostAddress %+v", self)

	case "hostname", "Hostname":
		return *self.Hostname, nil
		return nil, fmt.Errorf("Field Hostname no set on Hostname %+v", self)

	case "mesosConnected", "MesosConnected":
		return *self.MesosConnected, nil
		return nil, fmt.Errorf("Field MesosConnected no set on MesosConnected %+v", self)

	case "master", "Master":
		return *self.Master, nil
		return nil, fmt.Errorf("Field Master no set on Master %+v", self)

	case "uptime", "Uptime":
		return *self.Uptime, nil
		return nil, fmt.Errorf("Field Uptime no set on Uptime %+v", self)

	}
}

func (self *SingularityHostState) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityHostState", name)

	case "mesosMaster", "MesosMaster":
		self.MesosMaster = nil

	case "driverStatus", "DriverStatus":
		self.DriverStatus = nil

	case "availableCachedCpus", "AvailableCachedCpus":
		self.AvailableCachedCpus = nil

	case "millisSinceLastOffer", "MillisSinceLastOffer":
		self.MillisSinceLastOffer = nil

	case "offerCacheSize", "OfferCacheSize":
		self.OfferCacheSize = nil

	case "availableCachedMemory", "AvailableCachedMemory":
		self.AvailableCachedMemory = nil

	case "hostAddress", "HostAddress":
		self.HostAddress = nil

	case "hostname", "Hostname":
		self.Hostname = nil

	case "mesosConnected", "MesosConnected":
		self.MesosConnected = nil

	case "master", "Master":
		self.Master = nil

	case "uptime", "Uptime":
		self.Uptime = nil

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
