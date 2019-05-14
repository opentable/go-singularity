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
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_KILLING         SingularityTaskHistoryUpdateExtendedTaskState = "TASK_KILLING"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FINISHED        SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FINISHED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_FAILED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_FAILED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_KILLED          SingularityTaskHistoryUpdateExtendedTaskState = "TASK_KILLED"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST            SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityTaskHistoryUpdateExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityTaskHistoryUpdateExtendedTaskStateTASK_ERROR           SingularityTaskHistoryUpdateExtendedTaskState = "TASK_ERROR"
)

type SingularityTaskHistoryUpdate struct {
	Previous      *SingularityTaskHistoryUpdateList              `json:"previous,omitempty"`
	TaskId        *SingularityTaskId                             `json:"taskId,omitempty"`
	Timestamp     *int64                                         `json:"timestamp,omitempty"`
	TaskState     *SingularityTaskHistoryUpdateExtendedTaskState `json:"taskState,omitempty"`
	StatusMessage *string                                        `json:"statusMessage,omitempty"`
	StatusReason  *string                                        `json:"statusReason,omitempty"`
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

func (self *SingularityTaskHistoryUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskHistoryUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskHistoryUpdate) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "previous", "Previous":
		v, ok := value.(SingularityTaskHistoryUpdateList)
		if ok {
			self.Previous = &v
			return nil
		}
		return fmt.Errorf("Field previous/Previous: value %v(%T) couldn't be cast to type SingularityTaskHistoryUpdateList", value, value)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskState", "TaskState":
		v, ok := value.(SingularityTaskHistoryUpdateExtendedTaskState)
		if ok {
			self.TaskState = &v
			return nil
		}
		return fmt.Errorf("Field taskState/TaskState: value %v(%T) couldn't be cast to type SingularityTaskHistoryUpdateExtendedTaskState", value, value)

	case "statusMessage", "StatusMessage":
		v, ok := value.(string)
		if ok {
			self.StatusMessage = &v
			return nil
		}
		return fmt.Errorf("Field statusMessage/StatusMessage: value %v(%T) couldn't be cast to type string", value, value)

	case "statusReason", "StatusReason":
		v, ok := value.(string)
		if ok {
			self.StatusReason = &v
			return nil
		}
		return fmt.Errorf("Field statusReason/StatusReason: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityTaskHistoryUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "previous", "Previous":
		return *self.Previous, nil
		return nil, fmt.Errorf("Field Previous no set on Previous %+v", self)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "taskState", "TaskState":
		return *self.TaskState, nil
		return nil, fmt.Errorf("Field TaskState no set on TaskState %+v", self)

	case "statusMessage", "StatusMessage":
		return *self.StatusMessage, nil
		return nil, fmt.Errorf("Field StatusMessage no set on StatusMessage %+v", self)

	case "statusReason", "StatusReason":
		return *self.StatusReason, nil
		return nil, fmt.Errorf("Field StatusReason no set on StatusReason %+v", self)

	}
}

func (self *SingularityTaskHistoryUpdate) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskHistoryUpdate", name)

	case "previous", "Previous":
		self.Previous = nil

	case "taskId", "TaskId":
		self.TaskId = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "taskState", "TaskState":
		self.TaskState = nil

	case "statusMessage", "StatusMessage":
		self.StatusMessage = nil

	case "statusReason", "StatusReason":
		self.StatusReason = nil

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
