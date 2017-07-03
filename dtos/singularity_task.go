package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTask struct {
	present map[string]bool

	TaskRequest *SingularityTaskRequest `json:"taskRequest"`

	RackId string `json:"rackId,omitempty"`

	TaskId *SingularityTaskId `json:"taskId"`
}

func (self *SingularityTask) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTask) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTask); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTask cannot copy the values from %#v", other)
}

func (self *SingularityTask) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTask) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTask) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTask) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTask) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTask", name)

	case "taskRequest", "TaskRequest":
		v, ok := value.(*SingularityTaskRequest)
		if ok {
			self.TaskRequest = v
			self.present["taskRequest"] = true
			return nil
		} else {
			return fmt.Errorf("Field taskRequest/TaskRequest: value %v(%T) couldn't be cast to type *SingularityTaskRequest", value, value)
		}

	case "rackId", "RackId":
		v, ok := value.(string)
		if ok {
			self.RackId = v
			self.present["rackId"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackId/RackId: value %v(%T) couldn't be cast to type string", value, value)
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

	}
}

func (self *SingularityTask) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTask", name)

	case "taskRequest", "TaskRequest":
		if self.present != nil {
			if _, ok := self.present["taskRequest"]; ok {
				return self.TaskRequest, nil
			}
		}
		return nil, fmt.Errorf("Field TaskRequest no set on TaskRequest %+v", self)

	case "rackId", "RackId":
		if self.present != nil {
			if _, ok := self.present["rackId"]; ok {
				return self.RackId, nil
			}
		}
		return nil, fmt.Errorf("Field RackId no set on RackId %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	}
}

func (self *SingularityTask) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTask", name)

	case "taskRequest", "TaskRequest":
		self.present["taskRequest"] = false

	case "rackId", "RackId":
		self.present["rackId"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	}

	return nil
}

func (self *SingularityTask) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskList []*SingularityTask

func (self *SingularityTaskList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskList cannot copy the values from %#v", other)
}

func (list *SingularityTaskList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
