package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPendingRequestPendingType string

const (
	SingularityPendingRequestPendingTypeIMMEDIATE                   SingularityPendingRequestPendingType = "IMMEDIATE"
	SingularityPendingRequestPendingTypeONEOFF                      SingularityPendingRequestPendingType = "ONEOFF"
	SingularityPendingRequestPendingTypeBOUNCE                      SingularityPendingRequestPendingType = "BOUNCE"
	SingularityPendingRequestPendingTypeNEW_DEPLOY                  SingularityPendingRequestPendingType = "NEW_DEPLOY"
	SingularityPendingRequestPendingTypeNEXT_DEPLOY_STEP            SingularityPendingRequestPendingType = "NEXT_DEPLOY_STEP"
	SingularityPendingRequestPendingTypeUNPAUSED                    SingularityPendingRequestPendingType = "UNPAUSED"
	SingularityPendingRequestPendingTypeRETRY                       SingularityPendingRequestPendingType = "RETRY"
	SingularityPendingRequestPendingTypeUPDATED_REQUEST             SingularityPendingRequestPendingType = "UPDATED_REQUEST"
	SingularityPendingRequestPendingTypeDECOMISSIONED_SLAVE_OR_RACK SingularityPendingRequestPendingType = "DECOMISSIONED_SLAVE_OR_RACK"
	SingularityPendingRequestPendingTypeTASK_DONE                   SingularityPendingRequestPendingType = "TASK_DONE"
	SingularityPendingRequestPendingTypeSTARTUP                     SingularityPendingRequestPendingType = "STARTUP"
	SingularityPendingRequestPendingTypeCANCEL_BOUNCE               SingularityPendingRequestPendingType = "CANCEL_BOUNCE"
	SingularityPendingRequestPendingTypeTASK_BOUNCE                 SingularityPendingRequestPendingType = "TASK_BOUNCE"
	SingularityPendingRequestPendingTypeDEPLOY_CANCELLED            SingularityPendingRequestPendingType = "DEPLOY_CANCELLED"
	SingularityPendingRequestPendingTypeDEPLOY_FAILED               SingularityPendingRequestPendingType = "DEPLOY_FAILED"
)

type SingularityPendingRequest struct {
	present map[string]bool

	RequestId string `json:"requestId,omitempty"`

	DeployId string `json:"deployId,omitempty"`

	User string `json:"user,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`

	Message string `json:"message,omitempty"`

	ActionId string `json:"actionId,omitempty"`

	Resources *Resources `json:"resources"`

	Timestamp int64 `json:"timestamp"`

	PendingType SingularityPendingRequestPendingType `json:"pendingType"`

	CmdLineArgsList swaggering.StringList `json:"cmdLineArgsList"`

	RunId string `json:"runId,omitempty"`
}

func (self *SingularityPendingRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPendingRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingRequest cannot copy the values from %#v", other)
}

func (self *SingularityPendingRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPendingRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPendingRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
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

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = v
			self.present["user"] = true
			return nil
		} else {
			return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)
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

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			self.present["resources"] = true
			return nil
		} else {
			return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)
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

	case "pendingType", "PendingType":
		v, ok := value.(SingularityPendingRequestPendingType)
		if ok {
			self.PendingType = v
			self.present["pendingType"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingType/PendingType: value %v(%T) couldn't be cast to type SingularityPendingRequestPendingType", value, value)
		}

	case "cmdLineArgsList", "CmdLineArgsList":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CmdLineArgsList = v
			self.present["cmdLineArgsList"] = true
			return nil
		} else {
			return fmt.Errorf("Field cmdLineArgsList/CmdLineArgsList: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)
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

	}
}

func (self *SingularityPendingRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "resources", "Resources":
		if self.present != nil {
			if _, ok := self.present["resources"]; ok {
				return self.Resources, nil
			}
		}
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "pendingType", "PendingType":
		if self.present != nil {
			if _, ok := self.present["pendingType"]; ok {
				return self.PendingType, nil
			}
		}
		return nil, fmt.Errorf("Field PendingType no set on PendingType %+v", self)

	case "cmdLineArgsList", "CmdLineArgsList":
		if self.present != nil {
			if _, ok := self.present["cmdLineArgsList"]; ok {
				return self.CmdLineArgsList, nil
			}
		}
		return nil, fmt.Errorf("Field CmdLineArgsList no set on CmdLineArgsList %+v", self)

	case "runId", "RunId":
		if self.present != nil {
			if _, ok := self.present["runId"]; ok {
				return self.RunId, nil
			}
		}
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	}
}

func (self *SingularityPendingRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "user", "User":
		self.present["user"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

	case "message", "Message":
		self.present["message"] = false

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "resources", "Resources":
		self.present["resources"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "pendingType", "PendingType":
		self.present["pendingType"] = false

	case "cmdLineArgsList", "CmdLineArgsList":
		self.present["cmdLineArgsList"] = false

	case "runId", "RunId":
		self.present["runId"] = false

	}

	return nil
}

func (self *SingularityPendingRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPendingRequestList []*SingularityPendingRequest

func (self *SingularityPendingRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingRequestList cannot copy the values from %#v", other)
}

func (list *SingularityPendingRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
