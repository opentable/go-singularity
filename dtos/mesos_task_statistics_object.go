package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosTaskStatisticsObject struct {
	MemAnonBytes          *int64   `json:"memAnonBytes,omitempty"`
	MemFileBytes          *int64   `json:"memFileBytes,omitempty"`
	MemLimitBytes         *int64   `json:"memLimitBytes,omitempty"`
	Timestamp             *float64 `json:"timestamp,omitempty"`
	CpusUserTimeSecs      *float64 `json:"cpusUserTimeSecs,omitempty"`
	CpusNrPeriods         *int64   `json:"cpusNrPeriods,omitempty"`
	CpusNrThrottled       *int64   `json:"cpusNrThrottled,omitempty"`
	CpusSystemTimeSecs    *float64 `json:"cpusSystemTimeSecs,omitempty"`
	CpusThrottledTimeSecs *float64 `json:"cpusThrottledTimeSecs,omitempty"`
	MemMappedFileBytes    *int64   `json:"memMappedFileBytes,omitempty"`
	MemRssBytes           *int64   `json:"memRssBytes,omitempty"`
	CpusLimit             *int32   `json:"cpusLimit,omitempty"`
}

func (self *MesosTaskStatisticsObject) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *MesosTaskStatisticsObject) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosTaskStatisticsObject); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosTaskStatisticsObject cannot copy the values from %#v", other)
}

func (self *MesosTaskStatisticsObject) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosTaskStatisticsObject) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosTaskStatisticsObject) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "memAnonBytes", "MemAnonBytes":
		v, ok := value.(int64)
		if ok {
			self.MemAnonBytes = &v
			return nil
		}
		return fmt.Errorf("Field memAnonBytes/MemAnonBytes: value %v(%T) couldn't be cast to type int64", value, value)

	case "memFileBytes", "MemFileBytes":
		v, ok := value.(int64)
		if ok {
			self.MemFileBytes = &v
			return nil
		}
		return fmt.Errorf("Field memFileBytes/MemFileBytes: value %v(%T) couldn't be cast to type int64", value, value)

	case "memLimitBytes", "MemLimitBytes":
		v, ok := value.(int64)
		if ok {
			self.MemLimitBytes = &v
			return nil
		}
		return fmt.Errorf("Field memLimitBytes/MemLimitBytes: value %v(%T) couldn't be cast to type int64", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(float64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type float64", value, value)

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusUserTimeSecs = &v
			return nil
		}
		return fmt.Errorf("Field cpusUserTimeSecs/CpusUserTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)

	case "cpusNrPeriods", "CpusNrPeriods":
		v, ok := value.(int64)
		if ok {
			self.CpusNrPeriods = &v
			return nil
		}
		return fmt.Errorf("Field cpusNrPeriods/CpusNrPeriods: value %v(%T) couldn't be cast to type int64", value, value)

	case "cpusNrThrottled", "CpusNrThrottled":
		v, ok := value.(int64)
		if ok {
			self.CpusNrThrottled = &v
			return nil
		}
		return fmt.Errorf("Field cpusNrThrottled/CpusNrThrottled: value %v(%T) couldn't be cast to type int64", value, value)

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusSystemTimeSecs = &v
			return nil
		}
		return fmt.Errorf("Field cpusSystemTimeSecs/CpusSystemTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusThrottledTimeSecs = &v
			return nil
		}
		return fmt.Errorf("Field cpusThrottledTimeSecs/CpusThrottledTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)

	case "memMappedFileBytes", "MemMappedFileBytes":
		v, ok := value.(int64)
		if ok {
			self.MemMappedFileBytes = &v
			return nil
		}
		return fmt.Errorf("Field memMappedFileBytes/MemMappedFileBytes: value %v(%T) couldn't be cast to type int64", value, value)

	case "memRssBytes", "MemRssBytes":
		v, ok := value.(int64)
		if ok {
			self.MemRssBytes = &v
			return nil
		}
		return fmt.Errorf("Field memRssBytes/MemRssBytes: value %v(%T) couldn't be cast to type int64", value, value)

	case "cpusLimit", "CpusLimit":
		v, ok := value.(int32)
		if ok {
			self.CpusLimit = &v
			return nil
		}
		return fmt.Errorf("Field cpusLimit/CpusLimit: value %v(%T) couldn't be cast to type int32", value, value)

	}
}

func (self *MesosTaskStatisticsObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "memAnonBytes", "MemAnonBytes":
		return *self.MemAnonBytes, nil
		return nil, fmt.Errorf("Field MemAnonBytes no set on MemAnonBytes %+v", self)

	case "memFileBytes", "MemFileBytes":
		return *self.MemFileBytes, nil
		return nil, fmt.Errorf("Field MemFileBytes no set on MemFileBytes %+v", self)

	case "memLimitBytes", "MemLimitBytes":
		return *self.MemLimitBytes, nil
		return nil, fmt.Errorf("Field MemLimitBytes no set on MemLimitBytes %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		return *self.CpusUserTimeSecs, nil
		return nil, fmt.Errorf("Field CpusUserTimeSecs no set on CpusUserTimeSecs %+v", self)

	case "cpusNrPeriods", "CpusNrPeriods":
		return *self.CpusNrPeriods, nil
		return nil, fmt.Errorf("Field CpusNrPeriods no set on CpusNrPeriods %+v", self)

	case "cpusNrThrottled", "CpusNrThrottled":
		return *self.CpusNrThrottled, nil
		return nil, fmt.Errorf("Field CpusNrThrottled no set on CpusNrThrottled %+v", self)

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		return *self.CpusSystemTimeSecs, nil
		return nil, fmt.Errorf("Field CpusSystemTimeSecs no set on CpusSystemTimeSecs %+v", self)

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		return *self.CpusThrottledTimeSecs, nil
		return nil, fmt.Errorf("Field CpusThrottledTimeSecs no set on CpusThrottledTimeSecs %+v", self)

	case "memMappedFileBytes", "MemMappedFileBytes":
		return *self.MemMappedFileBytes, nil
		return nil, fmt.Errorf("Field MemMappedFileBytes no set on MemMappedFileBytes %+v", self)

	case "memRssBytes", "MemRssBytes":
		return *self.MemRssBytes, nil
		return nil, fmt.Errorf("Field MemRssBytes no set on MemRssBytes %+v", self)

	case "cpusLimit", "CpusLimit":
		return *self.CpusLimit, nil
		return nil, fmt.Errorf("Field CpusLimit no set on CpusLimit %+v", self)

	}
}

func (self *MesosTaskStatisticsObject) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "memAnonBytes", "MemAnonBytes":
		self.MemAnonBytes = nil

	case "memFileBytes", "MemFileBytes":
		self.MemFileBytes = nil

	case "memLimitBytes", "MemLimitBytes":
		self.MemLimitBytes = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		self.CpusUserTimeSecs = nil

	case "cpusNrPeriods", "CpusNrPeriods":
		self.CpusNrPeriods = nil

	case "cpusNrThrottled", "CpusNrThrottled":
		self.CpusNrThrottled = nil

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		self.CpusSystemTimeSecs = nil

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		self.CpusThrottledTimeSecs = nil

	case "memMappedFileBytes", "MemMappedFileBytes":
		self.MemMappedFileBytes = nil

	case "memRssBytes", "MemRssBytes":
		self.MemRssBytes = nil

	case "cpusLimit", "CpusLimit":
		self.CpusLimit = nil

	}

	return nil
}

func (self *MesosTaskStatisticsObject) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type MesosTaskStatisticsObjectList []*MesosTaskStatisticsObject

func (self *MesosTaskStatisticsObjectList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*MesosTaskStatisticsObjectList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A MesosTaskStatisticsObjectList cannot copy the values from %#v", other)
}

func (list *MesosTaskStatisticsObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *MesosTaskStatisticsObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosTaskStatisticsObjectList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
