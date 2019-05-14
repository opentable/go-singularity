package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestCleanupRequestCleanupType string

const (
	SingularityRequestCleanupRequestCleanupTypeDELETING           SingularityRequestCleanupRequestCleanupType = "DELETING"
	SingularityRequestCleanupRequestCleanupTypePAUSING            SingularityRequestCleanupRequestCleanupType = "PAUSING"
	SingularityRequestCleanupRequestCleanupTypeBOUNCE             SingularityRequestCleanupRequestCleanupType = "BOUNCE"
	SingularityRequestCleanupRequestCleanupTypeINCREMENTAL_BOUNCE SingularityRequestCleanupRequestCleanupType = "INCREMENTAL_BOUNCE"
)

type SingularityRequestCleanup struct {
	Timestamp                 *int64                                       `json:"timestamp,omitempty"`
	RequestId                 *string                                      `json:"requestId,omitempty"`
	ActionId                  *string                                      `json:"actionId,omitempty"`
	RunShellCommandBeforeKill *SingularityShellCommand                     `json:"runShellCommandBeforeKill,omitempty"`
	User                      *string                                      `json:"user,omitempty"`
	KillTasks                 *bool                                        `json:"killTasks,omitempty"`
	DeployId                  *string                                      `json:"deployId,omitempty"`
	Message                   *string                                      `json:"message,omitempty"`
	CleanupType               *SingularityRequestCleanupRequestCleanupType `json:"cleanupType,omitempty"`
	SkipHealthchecks          *bool                                        `json:"skipHealthchecks,omitempty"`
}

func (self *SingularityRequestCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequestCleanup) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestCleanup); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestCleanup cannot copy the values from %#v", other)
}

func (self *SingularityRequestCleanup) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestCleanup) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestCleanup) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		v, ok := value.(*SingularityShellCommand)
		if ok {
			self.RunShellCommandBeforeKill = v
			return nil
		}
		return fmt.Errorf("Field runShellCommandBeforeKill/RunShellCommandBeforeKill: value %v(%T) couldn't be cast to type *SingularityShellCommand", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = &v
			return nil
		}
		return fmt.Errorf("Field killTasks/KillTasks: value %v(%T) couldn't be cast to type bool", value, value)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "cleanupType", "CleanupType":
		v, ok := value.(SingularityRequestCleanupRequestCleanupType)
		if ok {
			self.CleanupType = &v
			return nil
		}
		return fmt.Errorf("Field cleanupType/CleanupType: value %v(%T) couldn't be cast to type SingularityRequestCleanupRequestCleanupType", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	}
}

func (self *SingularityRequestCleanup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		return self.RunShellCommandBeforeKill, nil
		return nil, fmt.Errorf("Field RunShellCommandBeforeKill no set on RunShellCommandBeforeKill %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "killTasks", "KillTasks":
		return *self.KillTasks, nil
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "cleanupType", "CleanupType":
		return *self.CleanupType, nil
		return nil, fmt.Errorf("Field CleanupType no set on CleanupType %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityRequestCleanup) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "runShellCommandBeforeKill", "RunShellCommandBeforeKill":
		self.RunShellCommandBeforeKill = nil

	case "user", "User":
		self.User = nil

	case "killTasks", "KillTasks":
		self.KillTasks = nil

	case "deployId", "DeployId":
		self.DeployId = nil

	case "message", "Message":
		self.Message = nil

	case "cleanupType", "CleanupType":
		self.CleanupType = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	}

	return nil
}

func (self *SingularityRequestCleanup) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestCleanupList []*SingularityRequestCleanup

func (self *SingularityRequestCleanupList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestCleanupList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestCleanupList cannot copy the values from %#v", other)
}

func (list *SingularityRequestCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestCleanupList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
