package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskShellCommandRequestId struct {
	TaskId    *SingularityTaskId `json:"taskId,omitempty"`
	Name      *string            `json:"name,omitempty"`
	Timestamp *int64             `json:"timestamp,omitempty"`
}

func (self *SingularityTaskShellCommandRequestId) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskShellCommandRequestId) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandRequestId); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandRequestId cannot copy the values from %#v", other)
}

func (self *SingularityTaskShellCommandRequestId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandRequestId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandRequestId) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = &v
			return nil
		}
		return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityTaskShellCommandRequestId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "name", "Name":
		return *self.Name, nil
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityTaskShellCommandRequestId) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "taskId", "TaskId":
		self.TaskId = nil

	case "name", "Name":
		self.Name = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	}

	return nil
}

func (self *SingularityTaskShellCommandRequestId) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskShellCommandRequestIdList []*SingularityTaskShellCommandRequestId

func (self *SingularityTaskShellCommandRequestIdList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandRequestIdList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandRequestIdList cannot copy the values from %#v", other)
}

func (list *SingularityTaskShellCommandRequestIdList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskShellCommandRequestIdList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskShellCommandRequestIdList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
