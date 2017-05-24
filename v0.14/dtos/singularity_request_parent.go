package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestParentRequestState string

const (
	SingularityRequestParentRequestStateACTIVE               SingularityRequestParentRequestState = "ACTIVE"
	SingularityRequestParentRequestStateDELETED              SingularityRequestParentRequestState = "DELETED"
	SingularityRequestParentRequestStatePAUSED               SingularityRequestParentRequestState = "PAUSED"
	SingularityRequestParentRequestStateSYSTEM_COOLDOWN      SingularityRequestParentRequestState = "SYSTEM_COOLDOWN"
	SingularityRequestParentRequestStateFINISHED             SingularityRequestParentRequestState = "FINISHED"
	SingularityRequestParentRequestStateDEPLOYING_TO_UNPAUSE SingularityRequestParentRequestState = "DEPLOYING_TO_UNPAUSE"
)

type SingularityRequestParent struct {
	present map[string]bool

	ActiveDeploy *SingularityDeploy `json:"activeDeploy"`

	ExpiringBounce *SingularityExpiringBounce `json:"expiringBounce"`

	ExpiringPause *SingularityExpiringPause `json:"expiringPause"`

	ExpiringScale *SingularityExpiringScale `json:"expiringScale"`

	ExpiringSkipHealthchecks *SingularityExpiringSkipHealthchecks `json:"expiringSkipHealthchecks"`

	PendingDeploy *SingularityDeploy `json:"pendingDeploy"`

	PendingDeployState *SingularityPendingDeploy `json:"pendingDeployState"`

	Request *SingularityRequest `json:"request"`

	RequestDeployState *SingularityRequestDeployState `json:"requestDeployState"`

	State SingularityRequestParentRequestState `json:"state"`
}

func (self *SingularityRequestParent) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequestParent) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestParent); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestParent cannot copy the values from %#v", other)
}

func (self *SingularityRequestParent) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequestParent) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestParent) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestParent) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequestParent) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "activeDeploy", "ActiveDeploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.ActiveDeploy = v
			self.present["activeDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeDeploy/ActiveDeploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)
		}

	case "expiringBounce", "ExpiringBounce":
		v, ok := value.(*SingularityExpiringBounce)
		if ok {
			self.ExpiringBounce = v
			self.present["expiringBounce"] = true
			return nil
		} else {
			return fmt.Errorf("Field expiringBounce/ExpiringBounce: value %v(%T) couldn't be cast to type *SingularityExpiringBounce", value, value)
		}

	case "expiringPause", "ExpiringPause":
		v, ok := value.(*SingularityExpiringPause)
		if ok {
			self.ExpiringPause = v
			self.present["expiringPause"] = true
			return nil
		} else {
			return fmt.Errorf("Field expiringPause/ExpiringPause: value %v(%T) couldn't be cast to type *SingularityExpiringPause", value, value)
		}

	case "expiringScale", "ExpiringScale":
		v, ok := value.(*SingularityExpiringScale)
		if ok {
			self.ExpiringScale = v
			self.present["expiringScale"] = true
			return nil
		} else {
			return fmt.Errorf("Field expiringScale/ExpiringScale: value %v(%T) couldn't be cast to type *SingularityExpiringScale", value, value)
		}

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		v, ok := value.(*SingularityExpiringSkipHealthchecks)
		if ok {
			self.ExpiringSkipHealthchecks = v
			self.present["expiringSkipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field expiringSkipHealthchecks/ExpiringSkipHealthchecks: value %v(%T) couldn't be cast to type *SingularityExpiringSkipHealthchecks", value, value)
		}

	case "pendingDeploy", "PendingDeploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.PendingDeploy = v
			self.present["pendingDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingDeploy/PendingDeploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)
		}

	case "pendingDeployState", "PendingDeployState":
		v, ok := value.(*SingularityPendingDeploy)
		if ok {
			self.PendingDeployState = v
			self.present["pendingDeployState"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingDeployState/PendingDeployState: value %v(%T) couldn't be cast to type *SingularityPendingDeploy", value, value)
		}

	case "request", "Request":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.Request = v
			self.present["request"] = true
			return nil
		} else {
			return fmt.Errorf("Field request/Request: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)
		}

	case "requestDeployState", "RequestDeployState":
		v, ok := value.(*SingularityRequestDeployState)
		if ok {
			self.RequestDeployState = v
			self.present["requestDeployState"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestDeployState/RequestDeployState: value %v(%T) couldn't be cast to type *SingularityRequestDeployState", value, value)
		}

	case "state", "State":
		v, ok := value.(SingularityRequestParentRequestState)
		if ok {
			self.State = v
			self.present["state"] = true
			return nil
		} else {
			return fmt.Errorf("Field state/State: value %v(%T) couldn't be cast to type SingularityRequestParentRequestState", value, value)
		}

	}
}

func (self *SingularityRequestParent) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "activeDeploy", "ActiveDeploy":
		if self.present != nil {
			if _, ok := self.present["activeDeploy"]; ok {
				return self.ActiveDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveDeploy no set on ActiveDeploy %+v", self)

	case "expiringBounce", "ExpiringBounce":
		if self.present != nil {
			if _, ok := self.present["expiringBounce"]; ok {
				return self.ExpiringBounce, nil
			}
		}
		return nil, fmt.Errorf("Field ExpiringBounce no set on ExpiringBounce %+v", self)

	case "expiringPause", "ExpiringPause":
		if self.present != nil {
			if _, ok := self.present["expiringPause"]; ok {
				return self.ExpiringPause, nil
			}
		}
		return nil, fmt.Errorf("Field ExpiringPause no set on ExpiringPause %+v", self)

	case "expiringScale", "ExpiringScale":
		if self.present != nil {
			if _, ok := self.present["expiringScale"]; ok {
				return self.ExpiringScale, nil
			}
		}
		return nil, fmt.Errorf("Field ExpiringScale no set on ExpiringScale %+v", self)

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["expiringSkipHealthchecks"]; ok {
				return self.ExpiringSkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field ExpiringSkipHealthchecks no set on ExpiringSkipHealthchecks %+v", self)

	case "pendingDeploy", "PendingDeploy":
		if self.present != nil {
			if _, ok := self.present["pendingDeploy"]; ok {
				return self.PendingDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field PendingDeploy no set on PendingDeploy %+v", self)

	case "pendingDeployState", "PendingDeployState":
		if self.present != nil {
			if _, ok := self.present["pendingDeployState"]; ok {
				return self.PendingDeployState, nil
			}
		}
		return nil, fmt.Errorf("Field PendingDeployState no set on PendingDeployState %+v", self)

	case "request", "Request":
		if self.present != nil {
			if _, ok := self.present["request"]; ok {
				return self.Request, nil
			}
		}
		return nil, fmt.Errorf("Field Request no set on Request %+v", self)

	case "requestDeployState", "RequestDeployState":
		if self.present != nil {
			if _, ok := self.present["requestDeployState"]; ok {
				return self.RequestDeployState, nil
			}
		}
		return nil, fmt.Errorf("Field RequestDeployState no set on RequestDeployState %+v", self)

	case "state", "State":
		if self.present != nil {
			if _, ok := self.present["state"]; ok {
				return self.State, nil
			}
		}
		return nil, fmt.Errorf("Field State no set on State %+v", self)

	}
}

func (self *SingularityRequestParent) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "activeDeploy", "ActiveDeploy":
		self.present["activeDeploy"] = false

	case "expiringBounce", "ExpiringBounce":
		self.present["expiringBounce"] = false

	case "expiringPause", "ExpiringPause":
		self.present["expiringPause"] = false

	case "expiringScale", "ExpiringScale":
		self.present["expiringScale"] = false

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		self.present["expiringSkipHealthchecks"] = false

	case "pendingDeploy", "PendingDeploy":
		self.present["pendingDeploy"] = false

	case "pendingDeployState", "PendingDeployState":
		self.present["pendingDeployState"] = false

	case "request", "Request":
		self.present["request"] = false

	case "requestDeployState", "RequestDeployState":
		self.present["requestDeployState"] = false

	case "state", "State":
		self.present["state"] = false

	}

	return nil
}

func (self *SingularityRequestParent) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestParentList []*SingularityRequestParent

func (self *SingularityRequestParentList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestParentList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestParentList cannot copy the values from %#v", other)
}

func (list *SingularityRequestParentList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestParentList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestParentList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
