package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type MesosTaskStatisticsObject struct {
	present map[string]bool

	CpusLimit int32 `json:"cpusLimit"`

	CpusNrPeriods int64 `json:"cpusNrPeriods"`

	CpusNrThrottled int64 `json:"cpusNrThrottled"`

	CpusSystemTimeSecs float64 `json:"cpusSystemTimeSecs"`

	CpusThrottledTimeSecs float64 `json:"cpusThrottledTimeSecs"`

	CpusUserTimeSecs float64 `json:"cpusUserTimeSecs"`

	MemAnonBytes int64 `json:"memAnonBytes"`

	MemFileBytes int64 `json:"memFileBytes"`

	MemLimitBytes int64 `json:"memLimitBytes"`

	MemMappedFileBytes int64 `json:"memMappedFileBytes"`

	MemRssBytes int64 `json:"memRssBytes"`

	Timestamp float64 `json:"timestamp"`
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

func (self *MesosTaskStatisticsObject) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *MesosTaskStatisticsObject) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *MesosTaskStatisticsObject) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *MesosTaskStatisticsObject) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *MesosTaskStatisticsObject) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "cpusLimit", "CpusLimit":
		v, ok := value.(int32)
		if ok {
			self.CpusLimit = v
			self.present["cpusLimit"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusLimit/CpusLimit: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "cpusNrPeriods", "CpusNrPeriods":
		v, ok := value.(int64)
		if ok {
			self.CpusNrPeriods = v
			self.present["cpusNrPeriods"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusNrPeriods/CpusNrPeriods: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "cpusNrThrottled", "CpusNrThrottled":
		v, ok := value.(int64)
		if ok {
			self.CpusNrThrottled = v
			self.present["cpusNrThrottled"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusNrThrottled/CpusNrThrottled: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusSystemTimeSecs = v
			self.present["cpusSystemTimeSecs"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusSystemTimeSecs/CpusSystemTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusThrottledTimeSecs = v
			self.present["cpusThrottledTimeSecs"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusThrottledTimeSecs/CpusThrottledTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		v, ok := value.(float64)
		if ok {
			self.CpusUserTimeSecs = v
			self.present["cpusUserTimeSecs"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpusUserTimeSecs/CpusUserTimeSecs: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "memAnonBytes", "MemAnonBytes":
		v, ok := value.(int64)
		if ok {
			self.MemAnonBytes = v
			self.present["memAnonBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field memAnonBytes/MemAnonBytes: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "memFileBytes", "MemFileBytes":
		v, ok := value.(int64)
		if ok {
			self.MemFileBytes = v
			self.present["memFileBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field memFileBytes/MemFileBytes: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "memLimitBytes", "MemLimitBytes":
		v, ok := value.(int64)
		if ok {
			self.MemLimitBytes = v
			self.present["memLimitBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field memLimitBytes/MemLimitBytes: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "memMappedFileBytes", "MemMappedFileBytes":
		v, ok := value.(int64)
		if ok {
			self.MemMappedFileBytes = v
			self.present["memMappedFileBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field memMappedFileBytes/MemMappedFileBytes: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "memRssBytes", "MemRssBytes":
		v, ok := value.(int64)
		if ok {
			self.MemRssBytes = v
			self.present["memRssBytes"] = true
			return nil
		} else {
			return fmt.Errorf("Field memRssBytes/MemRssBytes: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(float64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type float64", value, value)
		}

	}
}

func (self *MesosTaskStatisticsObject) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "cpusLimit", "CpusLimit":
		if self.present != nil {
			if _, ok := self.present["cpusLimit"]; ok {
				return self.CpusLimit, nil
			}
		}
		return nil, fmt.Errorf("Field CpusLimit no set on CpusLimit %+v", self)

	case "cpusNrPeriods", "CpusNrPeriods":
		if self.present != nil {
			if _, ok := self.present["cpusNrPeriods"]; ok {
				return self.CpusNrPeriods, nil
			}
		}
		return nil, fmt.Errorf("Field CpusNrPeriods no set on CpusNrPeriods %+v", self)

	case "cpusNrThrottled", "CpusNrThrottled":
		if self.present != nil {
			if _, ok := self.present["cpusNrThrottled"]; ok {
				return self.CpusNrThrottled, nil
			}
		}
		return nil, fmt.Errorf("Field CpusNrThrottled no set on CpusNrThrottled %+v", self)

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		if self.present != nil {
			if _, ok := self.present["cpusSystemTimeSecs"]; ok {
				return self.CpusSystemTimeSecs, nil
			}
		}
		return nil, fmt.Errorf("Field CpusSystemTimeSecs no set on CpusSystemTimeSecs %+v", self)

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		if self.present != nil {
			if _, ok := self.present["cpusThrottledTimeSecs"]; ok {
				return self.CpusThrottledTimeSecs, nil
			}
		}
		return nil, fmt.Errorf("Field CpusThrottledTimeSecs no set on CpusThrottledTimeSecs %+v", self)

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		if self.present != nil {
			if _, ok := self.present["cpusUserTimeSecs"]; ok {
				return self.CpusUserTimeSecs, nil
			}
		}
		return nil, fmt.Errorf("Field CpusUserTimeSecs no set on CpusUserTimeSecs %+v", self)

	case "memAnonBytes", "MemAnonBytes":
		if self.present != nil {
			if _, ok := self.present["memAnonBytes"]; ok {
				return self.MemAnonBytes, nil
			}
		}
		return nil, fmt.Errorf("Field MemAnonBytes no set on MemAnonBytes %+v", self)

	case "memFileBytes", "MemFileBytes":
		if self.present != nil {
			if _, ok := self.present["memFileBytes"]; ok {
				return self.MemFileBytes, nil
			}
		}
		return nil, fmt.Errorf("Field MemFileBytes no set on MemFileBytes %+v", self)

	case "memLimitBytes", "MemLimitBytes":
		if self.present != nil {
			if _, ok := self.present["memLimitBytes"]; ok {
				return self.MemLimitBytes, nil
			}
		}
		return nil, fmt.Errorf("Field MemLimitBytes no set on MemLimitBytes %+v", self)

	case "memMappedFileBytes", "MemMappedFileBytes":
		if self.present != nil {
			if _, ok := self.present["memMappedFileBytes"]; ok {
				return self.MemMappedFileBytes, nil
			}
		}
		return nil, fmt.Errorf("Field MemMappedFileBytes no set on MemMappedFileBytes %+v", self)

	case "memRssBytes", "MemRssBytes":
		if self.present != nil {
			if _, ok := self.present["memRssBytes"]; ok {
				return self.MemRssBytes, nil
			}
		}
		return nil, fmt.Errorf("Field MemRssBytes no set on MemRssBytes %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *MesosTaskStatisticsObject) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on MesosTaskStatisticsObject", name)

	case "cpusLimit", "CpusLimit":
		self.present["cpusLimit"] = false

	case "cpusNrPeriods", "CpusNrPeriods":
		self.present["cpusNrPeriods"] = false

	case "cpusNrThrottled", "CpusNrThrottled":
		self.present["cpusNrThrottled"] = false

	case "cpusSystemTimeSecs", "CpusSystemTimeSecs":
		self.present["cpusSystemTimeSecs"] = false

	case "cpusThrottledTimeSecs", "CpusThrottledTimeSecs":
		self.present["cpusThrottledTimeSecs"] = false

	case "cpusUserTimeSecs", "CpusUserTimeSecs":
		self.present["cpusUserTimeSecs"] = false

	case "memAnonBytes", "MemAnonBytes":
		self.present["memAnonBytes"] = false

	case "memFileBytes", "MemFileBytes":
		self.present["memFileBytes"] = false

	case "memLimitBytes", "MemLimitBytes":
		self.present["memLimitBytes"] = false

	case "memMappedFileBytes", "MemMappedFileBytes":
		self.present["memMappedFileBytes"] = false

	case "memRssBytes", "MemRssBytes":
		self.present["memRssBytes"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

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
