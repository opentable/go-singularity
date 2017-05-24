package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDisabledActionRequestSingularityAction string

const (
	SingularityDisabledActionRequestSingularityActionBOUNCE_REQUEST              SingularityDisabledActionRequestSingularityAction = "BOUNCE_REQUEST"
	SingularityDisabledActionRequestSingularityActionSCALE_REQUEST               SingularityDisabledActionRequestSingularityAction = "SCALE_REQUEST"
	SingularityDisabledActionRequestSingularityActionREMOVE_REQUEST              SingularityDisabledActionRequestSingularityAction = "REMOVE_REQUEST"
	SingularityDisabledActionRequestSingularityActionCREATE_REQUEST              SingularityDisabledActionRequestSingularityAction = "CREATE_REQUEST"
	SingularityDisabledActionRequestSingularityActionUPDATE_REQUEST              SingularityDisabledActionRequestSingularityAction = "UPDATE_REQUEST"
	SingularityDisabledActionRequestSingularityActionVIEW_REQUEST                SingularityDisabledActionRequestSingularityAction = "VIEW_REQUEST"
	SingularityDisabledActionRequestSingularityActionPAUSE_REQUEST               SingularityDisabledActionRequestSingularityAction = "PAUSE_REQUEST"
	SingularityDisabledActionRequestSingularityActionKILL_TASK                   SingularityDisabledActionRequestSingularityAction = "KILL_TASK"
	SingularityDisabledActionRequestSingularityActionBOUNCE_TASK                 SingularityDisabledActionRequestSingularityAction = "BOUNCE_TASK"
	SingularityDisabledActionRequestSingularityActionRUN_SHELL_COMMAND           SingularityDisabledActionRequestSingularityAction = "RUN_SHELL_COMMAND"
	SingularityDisabledActionRequestSingularityActionADD_METADATA                SingularityDisabledActionRequestSingularityAction = "ADD_METADATA"
	SingularityDisabledActionRequestSingularityActionDEPLOY                      SingularityDisabledActionRequestSingularityAction = "DEPLOY"
	SingularityDisabledActionRequestSingularityActionCANCEL_DEPLOY               SingularityDisabledActionRequestSingularityAction = "CANCEL_DEPLOY"
	SingularityDisabledActionRequestSingularityActionADD_WEBHOOK                 SingularityDisabledActionRequestSingularityAction = "ADD_WEBHOOK"
	SingularityDisabledActionRequestSingularityActionREMOVE_WEBHOOK              SingularityDisabledActionRequestSingularityAction = "REMOVE_WEBHOOK"
	SingularityDisabledActionRequestSingularityActionVIEW_WEBHOOKS               SingularityDisabledActionRequestSingularityAction = "VIEW_WEBHOOKS"
	SingularityDisabledActionRequestSingularityActionTASK_RECONCILIATION         SingularityDisabledActionRequestSingularityAction = "TASK_RECONCILIATION"
	SingularityDisabledActionRequestSingularityActionSTARTUP_TASK_RECONCILIATION SingularityDisabledActionRequestSingularityAction = "STARTUP_TASK_RECONCILIATION"
	SingularityDisabledActionRequestSingularityActionRUN_HEALTH_CHECKS           SingularityDisabledActionRequestSingularityAction = "RUN_HEALTH_CHECKS"
	SingularityDisabledActionRequestSingularityActionADD_DISASTER                SingularityDisabledActionRequestSingularityAction = "ADD_DISASTER"
	SingularityDisabledActionRequestSingularityActionREMOVE_DISASTER             SingularityDisabledActionRequestSingularityAction = "REMOVE_DISASTER"
	SingularityDisabledActionRequestSingularityActionDISABLE_ACTION              SingularityDisabledActionRequestSingularityAction = "DISABLE_ACTION"
	SingularityDisabledActionRequestSingularityActionENABLE_ACTION               SingularityDisabledActionRequestSingularityAction = "ENABLE_ACTION"
	SingularityDisabledActionRequestSingularityActionVIEW_DISASTERS              SingularityDisabledActionRequestSingularityAction = "VIEW_DISASTERS"
	SingularityDisabledActionRequestSingularityActionFREEZE_SLAVE                SingularityDisabledActionRequestSingularityAction = "FREEZE_SLAVE"
	SingularityDisabledActionRequestSingularityActionACTIVATE_SLAVE              SingularityDisabledActionRequestSingularityAction = "ACTIVATE_SLAVE"
	SingularityDisabledActionRequestSingularityActionDECOMMISSION_SLAVE          SingularityDisabledActionRequestSingularityAction = "DECOMMISSION_SLAVE"
	SingularityDisabledActionRequestSingularityActionVIEW_SLAVES                 SingularityDisabledActionRequestSingularityAction = "VIEW_SLAVES"
	SingularityDisabledActionRequestSingularityActionFREEZE_RACK                 SingularityDisabledActionRequestSingularityAction = "FREEZE_RACK"
	SingularityDisabledActionRequestSingularityActionACTIVATE_RACK               SingularityDisabledActionRequestSingularityAction = "ACTIVATE_RACK"
	SingularityDisabledActionRequestSingularityActionDECOMMISSION_RACK           SingularityDisabledActionRequestSingularityAction = "DECOMMISSION_RACK"
	SingularityDisabledActionRequestSingularityActionVIEW_RACKS                  SingularityDisabledActionRequestSingularityAction = "VIEW_RACKS"
	SingularityDisabledActionRequestSingularityActionSEND_EMAIL                  SingularityDisabledActionRequestSingularityAction = "SEND_EMAIL"
	SingularityDisabledActionRequestSingularityActionPROCESS_OFFERS              SingularityDisabledActionRequestSingularityAction = "PROCESS_OFFERS"
	SingularityDisabledActionRequestSingularityActionCACHE_OFFERS                SingularityDisabledActionRequestSingularityAction = "CACHE_OFFERS"
	SingularityDisabledActionRequestSingularityActionEXPENSIVE_API_CALLS         SingularityDisabledActionRequestSingularityAction = "EXPENSIVE_API_CALLS"
	SingularityDisabledActionRequestSingularityActionRUN_CLEANUP_POLLER          SingularityDisabledActionRequestSingularityAction = "RUN_CLEANUP_POLLER"
	SingularityDisabledActionRequestSingularityActionRUN_DEPLOY_POLLER           SingularityDisabledActionRequestSingularityAction = "RUN_DEPLOY_POLLER"
	SingularityDisabledActionRequestSingularityActionRUN_SCHEDULER_POLLER        SingularityDisabledActionRequestSingularityAction = "RUN_SCHEDULER_POLLER"
	SingularityDisabledActionRequestSingularityActionRUN_EXPIRING_ACTION_POLLER  SingularityDisabledActionRequestSingularityAction = "RUN_EXPIRING_ACTION_POLLER"
)

type SingularityDisabledActionRequest struct {
	present map[string]bool

	Message string `json:"message,omitempty"`

	Type SingularityDisabledActionRequestSingularityAction `json:"type"`
}

func (self *SingularityDisabledActionRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDisabledActionRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisabledActionRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisabledActionRequest cannot copy the values from %#v", other)
}

func (self *SingularityDisabledActionRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDisabledActionRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDisabledActionRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDisabledActionRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDisabledActionRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisabledActionRequest", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "type", "Type":
		v, ok := value.(SingularityDisabledActionRequestSingularityAction)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityDisabledActionRequestSingularityAction", value, value)
		}

	}
}

func (self *SingularityDisabledActionRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDisabledActionRequest", name)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "type", "Type":
		if self.present != nil {
			if _, ok := self.present["type"]; ok {
				return self.Type, nil
			}
		}
		return nil, fmt.Errorf("Field Type no set on Type %+v", self)

	}
}

func (self *SingularityDisabledActionRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisabledActionRequest", name)

	case "message", "Message":
		self.present["message"] = false

	case "type", "Type":
		self.present["type"] = false

	}

	return nil
}

func (self *SingularityDisabledActionRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDisabledActionRequestList []*SingularityDisabledActionRequest

func (self *SingularityDisabledActionRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisabledActionRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisabledActionRequestList cannot copy the values from %#v", other)
}

func (list *SingularityDisabledActionRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDisabledActionRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDisabledActionRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
