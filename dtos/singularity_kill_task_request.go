package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityKillTaskRequest struct {
	Message                   *string                  `json:"message,omitempty"`
	Override                  *bool                    `json:"override,omitempty"`
	ActionId                  *string                  `json:"actionId,omitempty"`
	WaitForReplacementTask    *bool                    `json:"waitForReplacementTask,omitempty"`
	RunShellCommandBeforeKill *SingularityShellCommand `json:"runShellCommandBeforeKill,omitempty"`
}

func (self *SingularityKillTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityKillTaskRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityKillTaskRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityKillTaskRequest cannot copy the values from %#v", other)
}

func (self *SingularityKillTaskRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityKillTaskRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityKillTaskRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityKillTaskRequest", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "override", "Override":
		v, ok := value.(bool)
		if ok {
			self.Override = &v
			return nil
		}
		return fmt.Errorf("Field override/Override: value %v(%T) couldn't be cast to type bool", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "waitForReplacementTask", "WaitForReplacementTask":
		v, ok := value.(bool)
		if ok {
			self.WaitForReplacementTask = &v
			return nil
		}
		return fmt.Errorf("Field waitForReplacementTask/WaitForReplacementTask: value %v(%T) couldn't be cast to type bool", value, value)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		v, ok := value.(*SingularityShellCommand)
		if ok {
			self.RunShellCommandBeforeKill = v
			return nil
		}
		return fmt.Errorf("Field runShellCommandBeforeKill/RunShellCommandBeforeKill: value %v(%T) couldn't be cast to type *SingularityShellCommand", value, value)

	}
}

func (self *SingularityKillTaskRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityKillTaskRequest", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "override", "Override":
		return *self.Override, nil
		return nil, fmt.Errorf("Field Override no set on Override %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "waitForReplacementTask", "WaitForReplacementTask":
		return *self.WaitForReplacementTask, nil
		return nil, fmt.Errorf("Field WaitForReplacementTask no set on WaitForReplacementTask %+v", self)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		return self.RunShellCommandBeforeKill, nil
		return nil, fmt.Errorf("Field RunShellCommandBeforeKill no set on RunShellCommandBeforeKill %+v", self)

	}
}

func (self *SingularityKillTaskRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityKillTaskRequest", name)

	case "message", "Message":
		self.Message = nil

	case "override", "Override":
		self.Override = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "waitForReplacementTask", "WaitForReplacementTask":
		self.WaitForReplacementTask = nil

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		self.RunShellCommandBeforeKill = nil

	}

	return nil
}

func (self *SingularityKillTaskRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityKillTaskRequestList []*SingularityKillTaskRequest

func (self *SingularityKillTaskRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityKillTaskRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityKillTaskRequestList cannot copy the values from %#v", other)
}

func (list *SingularityKillTaskRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityKillTaskRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityKillTaskRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
