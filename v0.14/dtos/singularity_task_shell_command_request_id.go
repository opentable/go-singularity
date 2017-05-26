package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskShellCommandRequestId struct {
	present map[string]bool

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	TaskId *SingularityTaskId `json:"taskId"`

	Timestamp int64 `json:"timestamp"`
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

func (self *SingularityTaskShellCommandRequestId) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskShellCommandRequestId) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandRequestId) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandRequestId) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskShellCommandRequestId) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "name", "Name":
		v, ok := value.(string)
		if ok {
			self.Name = v
			self.present["name"] = true
			return nil
		} else {
			return fmt.Errorf("Field name/Name: value %v(%T) couldn't be cast to type string", value, value)
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

func (self *SingularityTaskShellCommandRequestId) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "name", "Name":
		if self.present != nil {
			if _, ok := self.present["name"]; ok {
				return self.Name, nil
			}
		}
		return nil, fmt.Errorf("Field Name no set on Name %+v", self)

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

func (self *SingularityTaskShellCommandRequestId) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandRequestId", name)

	case "id", "Id":
		self.present["id"] = false

	case "name", "Name":
		self.present["name"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

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
