package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPauseRequest struct {
	KillTasks                 *bool                    `json:"killTasks,omitempty"`
	RunShellCommandBeforeKill *SingularityShellCommand `json:"runShellCommandBeforeKill,omitempty"`
	DurationMillis            *int64                   `json:"durationMillis,omitempty"`
	ActionId                  *string                  `json:"actionId,omitempty"`
	Message                   *string                  `json:"message,omitempty"`
}

func (self *SingularityPauseRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPauseRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPauseRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPauseRequest cannot copy the values from %#v", other)
}

func (self *SingularityPauseRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPauseRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPauseRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = &v
			return nil
		}
		return fmt.Errorf("Field killTasks/KillTasks: value %v(%T) couldn't be cast to type bool", value, value)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		v, ok := value.(*SingularityShellCommand)
		if ok {
			self.RunShellCommandBeforeKill = v
			return nil
		}
		return fmt.Errorf("Field runShellCommandBeforeKill/RunShellCommandBeforeKill: value %v(%T) couldn't be cast to type *SingularityShellCommand", value, value)

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = &v
			return nil
		}
		return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityPauseRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "killTasks", "KillTasks":
		return *self.KillTasks, nil
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		return self.RunShellCommandBeforeKill, nil
		return nil, fmt.Errorf("Field RunShellCommandBeforeKill no set on RunShellCommandBeforeKill %+v", self)

	case "durationMillis", "DurationMillis":
		return *self.DurationMillis, nil
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityPauseRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPauseRequest", name)

	case "killTasks", "KillTasks":
		self.KillTasks = nil

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		self.RunShellCommandBeforeKill = nil

	case "durationMillis", "DurationMillis":
		self.DurationMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "message", "Message":
		self.Message = nil

	}

	return nil
}

func (self *SingularityPauseRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPauseRequestList []*SingularityPauseRequest

func (self *SingularityPauseRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPauseRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPauseRequestList cannot copy the values from %#v", other)
}

func (list *SingularityPauseRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPauseRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPauseRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
