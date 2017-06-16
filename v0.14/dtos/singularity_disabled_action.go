package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDisabledActionSingularityAction string

const (
	SingularityDisabledActionSingularityActionBOUNCE_REQUEST      SingularityDisabledActionSingularityAction = "BOUNCE_REQUEST"
	SingularityDisabledActionSingularityActionSCALE_REQUEST       SingularityDisabledActionSingularityAction = "SCALE_REQUEST"
	SingularityDisabledActionSingularityActionREMOVE_REQUEST      SingularityDisabledActionSingularityAction = "REMOVE_REQUEST"
	SingularityDisabledActionSingularityActionCREATE_REQUEST      SingularityDisabledActionSingularityAction = "CREATE_REQUEST"
	SingularityDisabledActionSingularityActionUPDATE_REQUEST      SingularityDisabledActionSingularityAction = "UPDATE_REQUEST"
	SingularityDisabledActionSingularityActionVIEW_REQUEST        SingularityDisabledActionSingularityAction = "VIEW_REQUEST"
	SingularityDisabledActionSingularityActionPAUSE_REQUEST       SingularityDisabledActionSingularityAction = "PAUSE_REQUEST"
	SingularityDisabledActionSingularityActionKILL_TASK           SingularityDisabledActionSingularityAction = "KILL_TASK"
	SingularityDisabledActionSingularityActionBOUNCE_TASK         SingularityDisabledActionSingularityAction = "BOUNCE_TASK"
	SingularityDisabledActionSingularityActionRUN_SHELL_COMMAND   SingularityDisabledActionSingularityAction = "RUN_SHELL_COMMAND"
	SingularityDisabledActionSingularityActionADD_METADATA        SingularityDisabledActionSingularityAction = "ADD_METADATA"
	SingularityDisabledActionSingularityActionDEPLOY              SingularityDisabledActionSingularityAction = "DEPLOY"
	SingularityDisabledActionSingularityActionCANCEL_DEPLOY       SingularityDisabledActionSingularityAction = "CANCEL_DEPLOY"
	SingularityDisabledActionSingularityActionADD_WEBHOOK         SingularityDisabledActionSingularityAction = "ADD_WEBHOOK"
	SingularityDisabledActionSingularityActionREMOVE_WEBHOOK      SingularityDisabledActionSingularityAction = "REMOVE_WEBHOOK"
	SingularityDisabledActionSingularityActionVIEW_WEBHOOKS       SingularityDisabledActionSingularityAction = "VIEW_WEBHOOKS"
	SingularityDisabledActionSingularityActionTASK_RECONCILIATION SingularityDisabledActionSingularityAction = "TASK_RECONCILIATION"
	SingularityDisabledActionSingularityActionRUN_HEALTH_CHECKS   SingularityDisabledActionSingularityAction = "RUN_HEALTH_CHECKS"
	SingularityDisabledActionSingularityActionADD_DISASTER        SingularityDisabledActionSingularityAction = "ADD_DISASTER"
	SingularityDisabledActionSingularityActionREMOVE_DISASTER     SingularityDisabledActionSingularityAction = "REMOVE_DISASTER"
	SingularityDisabledActionSingularityActionDISABLE_ACTION      SingularityDisabledActionSingularityAction = "DISABLE_ACTION"
	SingularityDisabledActionSingularityActionENABLE_ACTION       SingularityDisabledActionSingularityAction = "ENABLE_ACTION"
	SingularityDisabledActionSingularityActionVIEW_DISASTERS      SingularityDisabledActionSingularityAction = "VIEW_DISASTERS"
	SingularityDisabledActionSingularityActionFREEZE_SLAVE        SingularityDisabledActionSingularityAction = "FREEZE_SLAVE"
	SingularityDisabledActionSingularityActionACTIVATE_SLAVE      SingularityDisabledActionSingularityAction = "ACTIVATE_SLAVE"
	SingularityDisabledActionSingularityActionDECOMMISSION_SLAVE  SingularityDisabledActionSingularityAction = "DECOMMISSION_SLAVE"
	SingularityDisabledActionSingularityActionVIEW_SLAVES         SingularityDisabledActionSingularityAction = "VIEW_SLAVES"
	SingularityDisabledActionSingularityActionFREEZE_RACK         SingularityDisabledActionSingularityAction = "FREEZE_RACK"
	SingularityDisabledActionSingularityActionACTIVATE_RACK       SingularityDisabledActionSingularityAction = "ACTIVATE_RACK"
	SingularityDisabledActionSingularityActionDECOMMISSION_RACK   SingularityDisabledActionSingularityAction = "DECOMMISSION_RACK"
	SingularityDisabledActionSingularityActionVIEW_RACKS          SingularityDisabledActionSingularityAction = "VIEW_RACKS"
	SingularityDisabledActionSingularityActionSEND_EMAIL          SingularityDisabledActionSingularityAction = "SEND_EMAIL"
	SingularityDisabledActionSingularityActionPROCESS_OFFERS      SingularityDisabledActionSingularityAction = "PROCESS_OFFERS"
	SingularityDisabledActionSingularityActionCACHE_OFFERS        SingularityDisabledActionSingularityAction = "CACHE_OFFERS"
)

type SingularityDisabledAction struct {
	present map[string]bool

	AutomaticallyClearable bool `json:"automaticallyClearable"`

	ExpiresAt int64 `json:"expiresAt"`

	Message string `json:"message,omitempty"`

	Type SingularityDisabledActionSingularityAction `json:"type"`

	User string `json:"user,omitempty"`
}

func (self *SingularityDisabledAction) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDisabledAction) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisabledAction); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisabledAction cannot copy the values from %#v", other)
}

func (self *SingularityDisabledAction) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDisabledAction) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDisabledAction) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDisabledAction) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDisabledAction) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisabledAction", name)

	case "automaticallyClearable", "AutomaticallyClearable":
		v, ok := value.(bool)
		if ok {
			self.AutomaticallyClearable = v
			self.present["automaticallyClearable"] = true
			return nil
		} else {
			return fmt.Errorf("Field automaticallyClearable/AutomaticallyClearable: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "expiresAt", "ExpiresAt":
		v, ok := value.(int64)
		if ok {
			self.ExpiresAt = v
			self.present["expiresAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field expiresAt/ExpiresAt: value %v(%T) couldn't be cast to type int64", value, value)
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

	case "type", "Type":
		v, ok := value.(SingularityDisabledActionSingularityAction)
		if ok {
			self.Type = v
			self.present["type"] = true
			return nil
		} else {
			return fmt.Errorf("Field type/Type: value %v(%T) couldn't be cast to type SingularityDisabledActionSingularityAction", value, value)
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

func (self *SingularityDisabledAction) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDisabledAction", name)

	case "automaticallyClearable", "AutomaticallyClearable":
		if self.present != nil {
			if _, ok := self.present["automaticallyClearable"]; ok {
				return self.AutomaticallyClearable, nil
			}
		}
		return nil, fmt.Errorf("Field AutomaticallyClearable no set on AutomaticallyClearable %+v", self)

	case "expiresAt", "ExpiresAt":
		if self.present != nil {
			if _, ok := self.present["expiresAt"]; ok {
				return self.ExpiresAt, nil
			}
		}
		return nil, fmt.Errorf("Field ExpiresAt no set on ExpiresAt %+v", self)

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

	case "user", "User":
		if self.present != nil {
			if _, ok := self.present["user"]; ok {
				return self.User, nil
			}
		}
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	}
}

func (self *SingularityDisabledAction) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisabledAction", name)

	case "automaticallyClearable", "AutomaticallyClearable":
		self.present["automaticallyClearable"] = false

	case "expiresAt", "ExpiresAt":
		self.present["expiresAt"] = false

	case "message", "Message":
		self.present["message"] = false

	case "type", "Type":
		self.present["type"] = false

	case "user", "User":
		self.present["user"] = false

	}

	return nil
}

func (self *SingularityDisabledAction) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDisabledActionList []*SingularityDisabledAction

func (self *SingularityDisabledActionList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisabledActionList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisabledActionList cannot copy the values from %#v", other)
}

func (list *SingularityDisabledActionList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDisabledActionList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDisabledActionList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
