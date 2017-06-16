package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityKilledTaskIdRecord struct {
	present map[string]bool

	OriginalTimestamp int64 `json:"originalTimestamp"`

	// RequestCleanupType *RequestCleanupType `json:"requestCleanupType"`

	Retries int32 `json:"retries"`

	// TaskCleanupType *TaskCleanupType `json:"taskCleanupType"`

	TaskId *SingularityTaskId `json:"taskId"`

	Timestamp int64 `json:"timestamp"`
}

func (self *SingularityKilledTaskIdRecord) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityKilledTaskIdRecord) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityKilledTaskIdRecord); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityKilledTaskIdRecord cannot copy the values from %#v", other)
}

func (self *SingularityKilledTaskIdRecord) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityKilledTaskIdRecord) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityKilledTaskIdRecord) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityKilledTaskIdRecord) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityKilledTaskIdRecord) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityKilledTaskIdRecord", name)

	case "originalTimestamp", "OriginalTimestamp":
		v, ok := value.(int64)
		if ok {
			self.OriginalTimestamp = v
			self.present["originalTimestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field originalTimestamp/OriginalTimestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "retries", "Retries":
		v, ok := value.(int32)
		if ok {
			self.Retries = v
			self.present["retries"] = true
			return nil
		} else {
			return fmt.Errorf("Field retries/Retries: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			self.present["taskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)
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

func (self *SingularityKilledTaskIdRecord) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityKilledTaskIdRecord", name)

	case "originalTimestamp", "OriginalTimestamp":
		if self.present != nil {
			if _, ok := self.present["originalTimestamp"]; ok {
				return self.OriginalTimestamp, nil
			}
		}
		return nil, fmt.Errorf("Field OriginalTimestamp no set on OriginalTimestamp %+v", self)

	case "retries", "Retries":
		if self.present != nil {
			if _, ok := self.present["retries"]; ok {
				return self.Retries, nil
			}
		}
		return nil, fmt.Errorf("Field Retries no set on Retries %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityKilledTaskIdRecord) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityKilledTaskIdRecord", name)

	case "originalTimestamp", "OriginalTimestamp":
		self.present["originalTimestamp"] = false

	case "retries", "Retries":
		self.present["retries"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityKilledTaskIdRecord) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityKilledTaskIdRecordList []*SingularityKilledTaskIdRecord

func (self *SingularityKilledTaskIdRecordList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityKilledTaskIdRecordList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityKilledTaskIdRecordList cannot copy the values from %#v", other)
}

func (list *SingularityKilledTaskIdRecordList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityKilledTaskIdRecordList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityKilledTaskIdRecordList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
