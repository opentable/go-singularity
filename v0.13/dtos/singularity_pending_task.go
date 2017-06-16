package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPendingTask struct {
	present map[string]bool

	CmdLineArgsList swaggering.StringList `json:"cmdLineArgsList"`

	Message string `json:"message,omitempty"`

	PendingTaskId *SingularityPendingTaskId `json:"pendingTaskId"`

	RunId string `json:"runId,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`

	User string `json:"user,omitempty"`
}

func (self *SingularityPendingTask) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPendingTask) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingTask); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingTask cannot copy the values from %#v", other)
}

func (self *SingularityPendingTask) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPendingTask) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingTask) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingTask) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPendingTask) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "cmdLineArgsList", "CmdLineArgsList":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CmdLineArgsList = v
			self.present["cmdLineArgsList"] = true
			return nil
		} else {
			return fmt.Errorf("Field cmdLineArgsList/CmdLineArgsList: value %v(%T) couldn't be cast to type StringList", value, value)
		}

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "pendingTaskId", "PendingTaskId":
		v, ok := value.(*SingularityPendingTaskId)
		if ok {
			self.PendingTaskId = v
			self.present["pendingTaskId"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingTaskId/PendingTaskId: value %v(%T) couldn't be cast to type *SingularityPendingTaskId", value, value)
		}

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = v
			self.present["runId"] = true
			return nil
		} else {
			return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = v
			self.present["skipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityPendingTask) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "cmdLineArgsList", "CmdLineArgsList":
		if self.present != nil {
			if _, ok := self.present["cmdLineArgsList"]; ok {
				return self.CmdLineArgsList, nil
			}
		}
		return nil, fmt.Errorf("Field CmdLineArgsList no set on CmdLineArgsList %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "pendingTaskId", "PendingTaskId":
		if self.present != nil {
			if _, ok := self.present["pendingTaskId"]; ok {
				return self.PendingTaskId, nil
			}
		}
		return nil, fmt.Errorf("Field PendingTaskId no set on PendingTaskId %+v", self)

	case "runId", "RunId":
		if self.present != nil {
			if _, ok := self.present["runId"]; ok {
				return self.RunId, nil
			}
		}
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityPendingTask) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "cmdLineArgsList", "CmdLineArgsList":
		self.present["cmdLineArgsList"] = false

	case "message", "Message":
		self.present["message"] = false

	case "pendingTaskId", "PendingTaskId":
		self.present["pendingTaskId"] = false

	case "runId", "RunId":
		self.present["runId"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityPendingTask) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPendingTaskList []*SingularityPendingTask

func (self *SingularityPendingTaskList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingTaskList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingTaskList cannot copy the values from %#v", other)
}

func (list *SingularityPendingTaskList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingTaskList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingTaskList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
