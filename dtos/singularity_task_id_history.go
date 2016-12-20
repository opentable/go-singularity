package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskIdHistory struct {
	present map[string]bool

	// LastTaskState *ExtendedTaskState `json:"lastTaskState"`

	RunId string `json:"runId,omitempty"`

	TaskId *SingularityTaskId `json:"taskId"`

	UpdatedAt int64 `json:"updatedAt"`
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

func (self *SingularityTaskIdHistory) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskIdHistory) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskIdHistory) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskIdHistory) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskIdHistory) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = v
			self.present["runId"] = true
			return nil
		} else {
			return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)
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

	case "updatedAt", "UpdatedAt":
		v, ok := value.(int64)
		if ok {
			self.UpdatedAt = v
			self.present["updatedAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field updatedAt/UpdatedAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityTaskIdHistory) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "runId", "RunId":
		if self.present != nil {
			if _, ok := self.present["runId"]; ok {
				return self.RunId, nil
			}
		}
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "taskId", "TaskId":
		if self.present != nil {
			if _, ok := self.present["taskId"]; ok {
				return self.TaskId, nil
			}
		}
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	case "updatedAt", "UpdatedAt":
		if self.present != nil {
			if _, ok := self.present["updatedAt"]; ok {
				return self.UpdatedAt, nil
			}
		}
		return nil, fmt.Errorf("Field UpdatedAt no set on UpdatedAt %+v", self)

	}
}

func (self *SingularityTaskIdHistory) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskIdHistory", name)

	case "runId", "RunId":
		self.present["runId"] = false

	case "taskId", "TaskId":
		self.present["taskId"] = false

	case "updatedAt", "UpdatedAt":
		self.present["updatedAt"] = false

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
