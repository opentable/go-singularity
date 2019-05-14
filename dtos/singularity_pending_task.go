package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPendingTask struct {
	Resources        *Resources                `json:"resources,omitempty"`
	ActionId         *string                   `json:"actionId,omitempty"`
	PendingTaskId    *SingularityPendingTaskId `json:"pendingTaskId,omitempty"`
	CmdLineArgsList  *swaggering.StringList    `json:"cmdLineArgsList,omitempty"`
	User             *string                   `json:"user,omitempty"`
	RunId            *string                   `json:"runId,omitempty"`
	SkipHealthchecks *bool                     `json:"skipHealthchecks,omitempty"`
	Message          *string                   `json:"message,omitempty"`
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

func (self *SingularityPendingTask) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingTask) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingTask) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			return nil
		}
		return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "pendingTaskId", "PendingTaskId":
		v, ok := value.(*SingularityPendingTaskId)
		if ok {
			self.PendingTaskId = v
			return nil
		}
		return fmt.Errorf("Field pendingTaskId/PendingTaskId: value %v(%T) couldn't be cast to type *SingularityPendingTaskId", value, value)

	case "cmdLineArgsList", "CmdLineArgsList":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CmdLineArgsList = &v
			return nil
		}
		return fmt.Errorf("Field cmdLineArgsList/CmdLineArgsList: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = &v
			return nil
		}
		return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityPendingTask) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "resources", "Resources":
		return self.Resources, nil
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "pendingTaskId", "PendingTaskId":
		return self.PendingTaskId, nil
		return nil, fmt.Errorf("Field PendingTaskId no set on PendingTaskId %+v", self)

	case "cmdLineArgsList", "CmdLineArgsList":
		return *self.CmdLineArgsList, nil
		return nil, fmt.Errorf("Field CmdLineArgsList no set on CmdLineArgsList %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "runId", "RunId":
		return *self.RunId, nil
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityPendingTask) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingTask", name)

	case "resources", "Resources":
		self.Resources = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "pendingTaskId", "PendingTaskId":
		self.PendingTaskId = nil

	case "cmdLineArgsList", "CmdLineArgsList":
		self.CmdLineArgsList = nil

	case "user", "User":
		self.User = nil

	case "runId", "RunId":
		self.RunId = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "message", "Message":
		self.Message = nil

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
