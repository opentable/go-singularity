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
	DeployId         *string                               `json:"deployId,omitempty"`
	Timestamp        *int64                                `json:"timestamp,omitempty"`
	PendingType      *SingularityPendingRequestPendingType `json:"pendingType,omitempty"`
	RunId            *string                               `json:"runId,omitempty"`
	Message          *string                               `json:"message,omitempty"`
	Resources        *Resources                            `json:"resources,omitempty"`
	RequestId        *string                               `json:"requestId,omitempty"`
	User             *string                               `json:"user,omitempty"`
	CmdLineArgsList  *swaggering.StringList                `json:"cmdLineArgsList,omitempty"`
	SkipHealthchecks *bool                                 `json:"skipHealthchecks,omitempty"`
	ActionId         *string                               `json:"actionId,omitempty"`
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

func (self *SingularityPendingRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = &v
			return nil
		}
		return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "pendingType", "PendingType":
		v, ok := value.(SingularityPendingRequestPendingType)
		if ok {
			self.PendingType = &v
			return nil
		}
		return fmt.Errorf("Field pendingType/PendingType: value %v(%T) couldn't be cast to type SingularityPendingRequestPendingType", value, value)

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = &v
			return nil
		}
		return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			return nil
		}
		return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = &v
			return nil
		}
		return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "cmdLineArgsList", "CmdLineArgsList":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CmdLineArgsList = &v
			return nil
		}
		return fmt.Errorf("Field cmdLineArgsList/CmdLineArgsList: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityPendingRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "deployId", "DeployId":
		return *self.DeployId, nil
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "pendingType", "PendingType":
		return *self.PendingType, nil
		return nil, fmt.Errorf("Field PendingType no set on PendingType %+v", self)

	case "runId", "RunId":
		return *self.RunId, nil
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "resources", "Resources":
		return self.Resources, nil
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	case "requestId", "RequestId":
		return *self.RequestId, nil
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "cmdLineArgsList", "CmdLineArgsList":
		return *self.CmdLineArgsList, nil
		return nil, fmt.Errorf("Field CmdLineArgsList no set on CmdLineArgsList %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	}
}

func (self *SingularityPendingRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingRequest", name)

	case "deployId", "DeployId":
		self.DeployId = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "pendingType", "PendingType":
		self.PendingType = nil

	case "runId", "RunId":
		self.RunId = nil

	case "message", "Message":
		self.Message = nil

	case "resources", "Resources":
		self.Resources = nil

	case "requestId", "RequestId":
		self.RequestId = nil

	case "user", "User":
		self.User = nil

	case "cmdLineArgsList", "CmdLineArgsList":
		self.CmdLineArgsList = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "actionId", "ActionId":
		self.ActionId = nil

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
