package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskHistoryUpdateExtendedTaskState string

const (
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LAUNCHED        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LAUNCHED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_STAGING         SingularityTaskHistoryUpdateExtendedTaskState = "TASK_STAGING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_STARTING        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_STARTING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_RUNNING         SingularityTaskHistoryUpdateExtendedTaskState = "TASK_RUNNING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_CLEANING        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_CLEANING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FINISHED        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FINISHED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FAILED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FAILED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_KILLED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_KILLED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST            SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_ERROR           SingularityTaskHistoryUpdateExtendedTaskState = "TASK_ERROR"
)

type SingularityTaskHistoryUpdate struct {
	present map[string]bool

	StatusMessage string `json:"statusMessage,omitempty"`

	StatusReason string `json:"statusReason,omitempty"`

	TaskId *SingularityTaskId `json:"taskId"`

	TaskState SingularityTaskHistoryUpdateExtendedTaskState `json:"taskState"`

	Timestamp int64 `json:"timestamp"`
}

func (self *SingularityTaskHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskHistoryUpdate) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHistoryUpdate); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHistoryUpdate cannot copy the values from %#v", other)
}

func (self *SingularityTaskHistoryUpdate) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskHistoryUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskHistoryUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskHistoryUpdate) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskHistoryUpdate) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "statusMessage", "StatusMessage":
		v, ok := value.(string)
		if ok {
			self.StatusMessage = v
			self.present["statusMessage"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusMessage/StatusMessage: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "statusReason", "StatusReason":
		v, ok := value.(string)
		if ok {
			self.StatusReason = v
			self.present["statusReason"] = true
			return nil
		} else {
			return fmt.Errorf("Field statusReason/StatusReason: value %v(%T) couldn't be cast to type string", value, value)
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

	case "taskState", "TaskState":
		v, ok := value.(SingularityTaskHistoryUpdateExtendedTaskState)
		if ok {
			self.TaskState = v
			self.present["taskState"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskState/TaskState: value %v(%T) couldn't be cast to type SingularityTaskHistoryUpdateExtendedTaskState", value, value)
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

func (self *SingularityTaskHistoryUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "statusMessage", "StatusMessage":
		if self.present != nil {
			if _, ok := self.present["statusMessage"]; ok {
				return self.StatusMessage, nil
			}
		}
		return nil, fmt.Errorf("Field StatusMessage no set on StatusMessage %+v", self)

	case "statusReason", "StatusReason":
		if self.present != nil {
			if _, ok := self.present["statusReason"]; ok {
				return self.StatusReason, nil
			}
		}
		return nil, fmt.Errorf("Field StatusReason no set on StatusReason %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "taskState", "TaskState":
		if self.present != nil {
			if _, ok := self.present["taskState"]; ok {
				return self.TaskState, nil
			}
		}
		return nil, fmt.Errorf("Field TaskState no set on TaskState %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityTaskHistoryUpdate) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "statusMessage", "StatusMessage":
		self.present["statusMessage"] = false

	case "statusReason", "StatusReason":
		self.present["statusReason"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "taskState", "TaskState":
		self.present["taskState"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityTaskHistoryUpdate) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskHistoryUpdateList []*SingularityTaskHistoryUpdate

func (self *SingularityTaskHistoryUpdateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskHistoryUpdateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskHistoryUpdateList cannot copy the values from %#v", other)
}

func (list *SingularityTaskHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskHistoryUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskHistoryUpdateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
