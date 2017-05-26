package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDisasterDataPoint struct {
	present map[string]bool

	AvgTaskLagMillis int64 `json:"avgTaskLagMillis"`

	NumActiveSlaves int32 `json:"numActiveSlaves"`

	NumActiveTasks int32 `json:"numActiveTasks"`

	NumLateTasks int32 `json:"numLateTasks"`

	NumLostSlaves int32 `json:"numLostSlaves"`

	NumLostTasks int32 `json:"numLostTasks"`

	NumPendingTasks int32 `json:"numPendingTasks"`

	Timestamp int64 `json:"timestamp"`
}

func (self *SingularityDisasterDataPoint) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDisasterDataPoint) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisasterDataPoint); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisasterDataPoint cannot copy the values from %#v", other)
}

func (self *SingularityDisasterDataPoint) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDisasterDataPoint) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDisasterDataPoint) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDisasterDataPoint) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDisasterDataPoint) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisasterDataPoint", name)

	case "avgTaskLagMillis", "AvgTaskLagMillis":
		v, ok := value.(int64)
		if ok {
			self.AvgTaskLagMillis = v
			self.present["avgTaskLagMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field avgTaskLagMillis/AvgTaskLagMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "numActiveSlaves", "NumActiveSlaves":
		v, ok := value.(int32)
		if ok {
			self.NumActiveSlaves = v
			self.present["numActiveSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field numActiveSlaves/NumActiveSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numActiveTasks", "NumActiveTasks":
		v, ok := value.(int32)
		if ok {
			self.NumActiveTasks = v
			self.present["numActiveTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field numActiveTasks/NumActiveTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numLateTasks", "NumLateTasks":
		v, ok := value.(int32)
		if ok {
			self.NumLateTasks = v
			self.present["numLateTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field numLateTasks/NumLateTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numLostSlaves", "NumLostSlaves":
		v, ok := value.(int32)
		if ok {
			self.NumLostSlaves = v
			self.present["numLostSlaves"] = true
			return nil
		} else {
			return fmt.Errorf("Field numLostSlaves/NumLostSlaves: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numLostTasks", "NumLostTasks":
		v, ok := value.(int32)
		if ok {
			self.NumLostTasks = v
			self.present["numLostTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field numLostTasks/NumLostTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "numPendingTasks", "NumPendingTasks":
		v, ok := value.(int32)
		if ok {
			self.NumPendingTasks = v
			self.present["numPendingTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field numPendingTasks/NumPendingTasks: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityDisasterDataPoint) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDisasterDataPoint", name)

	case "avgTaskLagMillis", "AvgTaskLagMillis":
		if self.present != nil {
			if _, ok := self.present["avgTaskLagMillis"]; ok {
				return self.AvgTaskLagMillis, nil
			}
		}
		return nil, fmt.Errorf("Field AvgTaskLagMillis no set on AvgTaskLagMillis %+v", self)

	case "numActiveSlaves", "NumActiveSlaves":
		if self.present != nil {
			if _, ok := self.present["numActiveSlaves"]; ok {
				return self.NumActiveSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field NumActiveSlaves no set on NumActiveSlaves %+v", self)

	case "numActiveTasks", "NumActiveTasks":
		if self.present != nil {
			if _, ok := self.present["numActiveTasks"]; ok {
				return self.NumActiveTasks, nil
			}
		}
		return nil, fmt.Errorf("Field NumActiveTasks no set on NumActiveTasks %+v", self)

	case "numLateTasks", "NumLateTasks":
		if self.present != nil {
			if _, ok := self.present["numLateTasks"]; ok {
				return self.NumLateTasks, nil
			}
		}
		return nil, fmt.Errorf("Field NumLateTasks no set on NumLateTasks %+v", self)

	case "numLostSlaves", "NumLostSlaves":
		if self.present != nil {
			if _, ok := self.present["numLostSlaves"]; ok {
				return self.NumLostSlaves, nil
			}
		}
		return nil, fmt.Errorf("Field NumLostSlaves no set on NumLostSlaves %+v", self)

	case "numLostTasks", "NumLostTasks":
		if self.present != nil {
			if _, ok := self.present["numLostTasks"]; ok {
				return self.NumLostTasks, nil
			}
		}
		return nil, fmt.Errorf("Field NumLostTasks no set on NumLostTasks %+v", self)

	case "numPendingTasks", "NumPendingTasks":
		if self.present != nil {
			if _, ok := self.present["numPendingTasks"]; ok {
				return self.NumPendingTasks, nil
			}
		}
		return nil, fmt.Errorf("Field NumPendingTasks no set on NumPendingTasks %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityDisasterDataPoint) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisasterDataPoint", name)

	case "avgTaskLagMillis", "AvgTaskLagMillis":
		self.present["avgTaskLagMillis"] = false

	case "numActiveSlaves", "NumActiveSlaves":
		self.present["numActiveSlaves"] = false

	case "numActiveTasks", "NumActiveTasks":
		self.present["numActiveTasks"] = false

	case "numLateTasks", "NumLateTasks":
		self.present["numLateTasks"] = false

	case "numLostSlaves", "NumLostSlaves":
		self.present["numLostSlaves"] = false

	case "numLostTasks", "NumLostTasks":
		self.present["numLostTasks"] = false

	case "numPendingTasks", "NumPendingTasks":
		self.present["numPendingTasks"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityDisasterDataPoint) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDisasterDataPointList []*SingularityDisasterDataPoint

func (self *SingularityDisasterDataPointList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisasterDataPointList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisasterDataPointList cannot copy the values from %#v", other)
}

func (list *SingularityDisasterDataPointList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDisasterDataPointList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDisasterDataPointList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
