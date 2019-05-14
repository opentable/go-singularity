package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskIdHistoryExtendedTaskState string

const (
	SingularityTaskIdHistoryExtendedTaskStateTASK_LAUNCHED        SingularityTaskIdHistoryExtendedTaskState = "TASK_LAUNCHED"
	SingularityTaskIdHistoryExtendedTaskStateTASK_STAGING         SingularityTaskIdHistoryExtendedTaskState = "TASK_STAGING"
	SingularityTaskIdHistoryExtendedTaskStateTASK_STARTING        SingularityTaskIdHistoryExtendedTaskState = "TASK_STARTING"
	SingularityTaskIdHistoryExtendedTaskStateTASK_RUNNING         SingularityTaskIdHistoryExtendedTaskState = "TASK_RUNNING"
	SingularityTaskIdHistoryExtendedTaskStateTASK_CLEANING        SingularityTaskIdHistoryExtendedTaskState = "TASK_CLEANING"
	SingularityTaskIdHistoryExtendedTaskStateTASK_KILLING         SingularityTaskIdHistoryExtendedTaskState = "TASK_KILLING"
	SingularityTaskIdHistoryExtendedTaskStateTASK_FINISHED        SingularityTaskIdHistoryExtendedTaskState = "TASK_FINISHED"
	SingularityTaskIdHistoryExtendedTaskStateTASK_FAILED          SingularityTaskIdHistoryExtendedTaskState = "TASK_FAILED"
	SingularityTaskIdHistoryExtendedTaskStateTASK_KILLED          SingularityTaskIdHistoryExtendedTaskState = "TASK_KILLED"
	SingularityTaskIdHistoryExtendedTaskStateTASK_LOST            SingularityTaskIdHistoryExtendedTaskState = "TASK_LOST"
	SingularityTaskIdHistoryExtendedTaskStateTASK_LOST_WHILE_DOWN SingularityTaskIdHistoryExtendedTaskState = "TASK_LOST_WHILE_DOWN"
	SingularityTaskIdHistoryExtendedTaskStateTASK_ERROR           SingularityTaskIdHistoryExtendedTaskState = "TASK_ERROR"
)

type SingularityTaskIdHistory struct {
	TaskId        *SingularityTaskId                         `json:"taskId,omitempty"`
	UpdatedAt     *int64                                     `json:"updatedAt,omitempty"`
	LastTaskState *SingularityTaskIdHistoryExtendedTaskState `json:"lastTaskState,omitempty"`
	RunId         *string                                    `json:"runId,omitempty"`
}

func (self *SingularityTaskIdHistory) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskIdHistory) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskIdHistory); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskIdHistory cannot copy the values from %#v", other)
}

func (self *SingularityTaskIdHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskIdHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskIdHistory) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "updatedAt", "UpdatedAt":
		v, ok := value.(int64)
		if ok {
			self.UpdatedAt = &v
			return nil
		}
		return fmt.Errorf("Field updatedAt/UpdatedAt: value %v(%T) couldn't be cast to type int64", value, value)

	case "lastTaskState", "LastTaskState":
		v, ok := value.(SingularityTaskIdHistoryExtendedTaskState)
		if ok {
			self.LastTaskState = &v
			return nil
		}
		return fmt.Errorf("Field lastTaskState/LastTaskState: value %v(%T) couldn't be cast to type SingularityTaskIdHistoryExtendedTaskState", value, value)

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = &v
			return nil
		}
		return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityTaskIdHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "updatedAt", "UpdatedAt":
		return *self.UpdatedAt, nil
		return nil, fmt.Errorf("Field UpdatedAt no set on UpdatedAt %+v", self)

	case "lastTaskState", "LastTaskState":
		return *self.LastTaskState, nil
		return nil, fmt.Errorf("Field LastTaskState no set on LastTaskState %+v", self)

	case "runId", "RunId":
		return *self.RunId, nil
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	}
}

func (self *SingularityTaskIdHistory) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "taskId", "TaskId":
		self.TaskId = nil

	case "updatedAt", "UpdatedAt":
		self.UpdatedAt = nil

	case "lastTaskState", "LastTaskState":
		self.LastTaskState = nil

	case "runId", "RunId":
		self.RunId = nil

	}

	return nil
}

func (self *SingularityTaskIdHistory) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskIdHistoryList []*SingularityTaskIdHistory

func (self *SingularityTaskIdHistoryList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskIdHistoryList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskIdHistoryList cannot copy the values from %#v", other)
}

func (list *SingularityTaskIdHistoryList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskIdHistoryList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskIdHistoryList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
