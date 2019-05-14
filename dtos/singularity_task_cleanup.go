package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskCleanupTaskCleanupType string

const (
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED               SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED"
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED_TASK_BOUNCE   SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED_TASK_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDECOMISSIONING               SingularityTaskCleanupTaskCleanupType = "DECOMISSIONING"
	SingularityTaskCleanupTaskCleanupTypeSCALING_DOWN                 SingularityTaskCleanupTaskCleanupType = "SCALING_DOWN"
	SingularityTaskCleanupTaskCleanupTypeBOUNCING                     SingularityTaskCleanupTaskCleanupType = "BOUNCING"
	SingularityTaskCleanupTaskCleanupTypeINCREMENTAL_BOUNCE           SingularityTaskCleanupTaskCleanupType = "INCREMENTAL_BOUNCE"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_FAILED                SingularityTaskCleanupTaskCleanupType = "DEPLOY_FAILED"
	SingularityTaskCleanupTaskCleanupTypeNEW_DEPLOY_SUCCEEDED         SingularityTaskCleanupTaskCleanupType = "NEW_DEPLOY_SUCCEEDED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_STEP_FINISHED         SingularityTaskCleanupTaskCleanupType = "DEPLOY_STEP_FINISHED"
	SingularityTaskCleanupTaskCleanupTypeDEPLOY_CANCELED              SingularityTaskCleanupTaskCleanupType = "DEPLOY_CANCELED"
	SingularityTaskCleanupTaskCleanupTypeTASK_EXCEEDED_TIME_LIMIT     SingularityTaskCleanupTaskCleanupType = "TASK_EXCEEDED_TIME_LIMIT"
	SingularityTaskCleanupTaskCleanupTypeUNHEALTHY_NEW_TASK           SingularityTaskCleanupTaskCleanupType = "UNHEALTHY_NEW_TASK"
	SingularityTaskCleanupTaskCleanupTypeOVERDUE_NEW_TASK             SingularityTaskCleanupTaskCleanupType = "OVERDUE_NEW_TASK"
	SingularityTaskCleanupTaskCleanupTypeUSER_REQUESTED_DESTROY       SingularityTaskCleanupTaskCleanupType = "USER_REQUESTED_DESTROY"
	SingularityTaskCleanupTaskCleanupTypeINCREMENTAL_DEPLOY_FAILED    SingularityTaskCleanupTaskCleanupType = "INCREMENTAL_DEPLOY_FAILED"
	SingularityTaskCleanupTaskCleanupTypeINCREMENTAL_DEPLOY_CANCELLED SingularityTaskCleanupTaskCleanupType = "INCREMENTAL_DEPLOY_CANCELLED"
	SingularityTaskCleanupTaskCleanupTypePRIORITY_KILL                SingularityTaskCleanupTaskCleanupType = "PRIORITY_KILL"
	SingularityTaskCleanupTaskCleanupTypeREBALANCE_RACKS              SingularityTaskCleanupTaskCleanupType = "REBALANCE_RACKS"
	SingularityTaskCleanupTaskCleanupTypePAUSING                      SingularityTaskCleanupTaskCleanupType = "PAUSING"
	SingularityTaskCleanupTaskCleanupTypePAUSE                        SingularityTaskCleanupTaskCleanupType = "PAUSE"
	SingularityTaskCleanupTaskCleanupTypeDECOMMISSION_TIMEOUT         SingularityTaskCleanupTaskCleanupType = "DECOMMISSION_TIMEOUT"
	SingularityTaskCleanupTaskCleanupTypeREQUEST_DELETING             SingularityTaskCleanupTaskCleanupType = "REQUEST_DELETING"
)

type SingularityTaskCleanup struct {
	Message         *string                                `json:"message,omitempty"`
	ActionId        *string                                `json:"actionId,omitempty"`
	RunBeforeKillId *SingularityTaskShellCommandRequestId  `json:"runBeforeKillId,omitempty"`
	User            *string                                `json:"user,omitempty"`
	CleanupType     *SingularityTaskCleanupTaskCleanupType `json:"cleanupType,omitempty"`
	Timestamp       *int64                                 `json:"timestamp,omitempty"`
	TaskId          *SingularityTaskId                     `json:"taskId,omitempty"`
}

func (self *SingularityTaskCleanup) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskCleanup) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskCleanup); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskCleanup cannot copy the values from %#v", other)
}

func (self *SingularityTaskCleanup) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskCleanup) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskCleanup) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "runBeforeKillId", "RunBeforeKillId":
		v, ok := value.(*SingularityTaskShellCommandRequestId)
		if ok {
			self.RunBeforeKillId = v
			return nil
		}
		return fmt.Errorf("Field runBeforeKillId/RunBeforeKillId: value %v(%T) couldn't be cast to type *SingularityTaskShellCommandRequestId", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "cleanupType", "CleanupType":
		v, ok := value.(SingularityTaskCleanupTaskCleanupType)
		if ok {
			self.CleanupType = &v
			return nil
		}
		return fmt.Errorf("Field cleanupType/CleanupType: value %v(%T) couldn't be cast to type SingularityTaskCleanupTaskCleanupType", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "taskId", "TaskId":
		v, ok := value.(*SingularityTaskId)
		if ok {
			self.TaskId = v
			return nil
		}
		return fmt.Errorf("Field taskId/TaskId: value %v(%T) couldn't be cast to type *SingularityTaskId", value, value)

	}
}

func (self *SingularityTaskCleanup) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "runBeforeKillId", "RunBeforeKillId":
		return self.RunBeforeKillId, nil
		return nil, fmt.Errorf("Field RunBeforeKillId no set on RunBeforeKillId %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "cleanupType", "CleanupType":
		return *self.CleanupType, nil
		return nil, fmt.Errorf("Field CleanupType no set on CleanupType %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "taskId", "TaskId":
		return self.TaskId, nil
		return nil, fmt.Errorf("Field TaskId no set on TaskId %+v", self)

	}
}

func (self *SingularityTaskCleanup) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskCleanup", name)

	case "message", "Message":
		self.Message = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "runBeforeKillId", "RunBeforeKillId":
		self.RunBeforeKillId = nil

	case "user", "User":
		self.User = nil

	case "cleanupType", "CleanupType":
		self.CleanupType = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "taskId", "TaskId":
		self.TaskId = nil

	}

	return nil
}

func (self *SingularityTaskCleanup) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskCleanupList []*SingularityTaskCleanup

func (self *SingularityTaskCleanupList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskCleanupList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskCleanupList cannot copy the values from %#v", other)
}

func (list *SingularityTaskCleanupList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskCleanupList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskCleanupList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
