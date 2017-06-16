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
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	CleanupType SingularityRequestCleanupRequestCleanupType `json:"cleanupType"`

	DeployId string `json:"deployId,omitempty"`

	KillTasks bool `json:"killTasks"`

	Message string `json:"message,omitempty"`

	RequestId string `json:"requestId,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`

	Timestamp int64 `json:"timestamp"`

	User string `json:"user,omitempty"`
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

func (self *SingularityRequestCleanup) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequestCleanup) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestCleanup) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestCleanup) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequestCleanup) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "cleanupType", "CleanupType":
		v, ok := value.(SingularityRequestCleanupRequestCleanupType)
		if ok {
			self.CleanupType = v
			self.present["cleanupType"] = true
			return nil
		} else {
			return fmt.Errorf("Field cleanupType/CleanupType: value %v(%T) couldn't be cast to type SingularityRequestCleanupRequestCleanupType", value, value)
		}

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "killTasks", "KillTasks":
		v, ok := value.(bool)
		if ok {
			self.KillTasks = v
			self.present["killTasks"] = true
			return nil
		} else {
			return fmt.Errorf("Field killTasks/KillTasks: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
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

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
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

func (self *SingularityRequestCleanup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "cleanupType", "CleanupType":
		if self.present != nil {
			if _, ok := self.present["cleanupType"]; ok {
				return self.CleanupType, nil
			}
		}
		return nil, fmt.Errorf("Field CleanupType no set on CleanupType %+v", self)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "killTasks", "KillTasks":
		if self.present != nil {
			if _, ok := self.present["killTasks"]; ok {
				return self.KillTasks, nil
			}
		}
		return nil, fmt.Errorf("Field KillTasks no set on KillTasks %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityRequestCleanup) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestCleanup", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "cleanupType", "CleanupType":
		self.present["cleanupType"] = false

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "killTasks", "KillTasks":
		self.present["killTasks"] = false

	case "message", "Message":
		self.present["message"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "user", "User":
		self.present["user"] = false

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
