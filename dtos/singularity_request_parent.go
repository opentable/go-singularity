package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestParentRequestState string

const (
	SingularityRequestParentRequestStateACTIVE               SingularityRequestParentRequestState = "ACTIVE"
	SingularityRequestParentRequestStateDELETING             SingularityRequestParentRequestState = "DELETING"
	SingularityRequestParentRequestStateDELETED              SingularityRequestParentRequestState = "DELETED"
	SingularityRequestParentRequestStatePAUSED               SingularityRequestParentRequestState = "PAUSED"
	SingularityRequestParentRequestStateSYSTEM_COOLDOWN      SingularityRequestParentRequestState = "SYSTEM_COOLDOWN"
	SingularityRequestParentRequestStateFINISHED             SingularityRequestParentRequestState = "FINISHED"
	SingularityRequestParentRequestStateDEPLOYING_TO_UNPAUSE SingularityRequestParentRequestState = "DEPLOYING_TO_UNPAUSE"
)

type SingularityRequestParent struct {
	Request                  *SingularityRequest                   `json:"request,omitempty"`
	RequestDeployState       *SingularityRequestDeployState        `json:"requestDeployState,omitempty"`
	PendingDeploy            *SingularityDeploy                    `json:"pendingDeploy,omitempty"`
	PendingDeployState       *SingularityPendingDeploy             `json:"pendingDeployState,omitempty"`
	ExpiringBounce           *SingularityExpiringBounce            `json:"expiringBounce,omitempty"`
	ExpiringScale            *SingularityExpiringScale             `json:"expiringScale,omitempty"`
	ExpiringSkipHealthchecks *SingularityExpiringSkipHealthchecks  `json:"expiringSkipHealthchecks,omitempty"`
	State                    *SingularityRequestParentRequestState `json:"state,omitempty"`
	ActiveDeploy             *SingularityDeploy                    `json:"activeDeploy,omitempty"`
	ExpiringPause            *SingularityExpiringPause             `json:"expiringPause,omitempty"`
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

func (self *SingularityRequestParent) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestParent) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestParent) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "request", "Request":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.Request = v
			return nil
		}
		return fmt.Errorf("Field request/Request: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)

	case "requestDeployState", "RequestDeployState":
		v, ok := value.(*SingularityRequestDeployState)
		if ok {
			self.RequestDeployState = v
			return nil
		}
		return fmt.Errorf("Field requestDeployState/RequestDeployState: value %v(%T) couldn't be cast to type *SingularityRequestDeployState", value, value)

	case "pendingDeploy", "PendingDeploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.PendingDeploy = v
			return nil
		}
		return fmt.Errorf("Field pendingDeploy/PendingDeploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "pendingDeployState", "PendingDeployState":
		v, ok := value.(*SingularityPendingDeploy)
		if ok {
			self.PendingDeployState = v
			return nil
		}
		return fmt.Errorf("Field pendingDeployState/PendingDeployState: value %v(%T) couldn't be cast to type *SingularityPendingDeploy", value, value)

	case "expiringBounce", "ExpiringBounce":
		v, ok := value.(*SingularityExpiringBounce)
		if ok {
			self.ExpiringBounce = v
			return nil
		}
		return fmt.Errorf("Field expiringBounce/ExpiringBounce: value %v(%T) couldn't be cast to type *SingularityExpiringBounce", value, value)

	case "expiringScale", "ExpiringScale":
		v, ok := value.(*SingularityExpiringScale)
		if ok {
			self.ExpiringScale = v
			return nil
		}
		return fmt.Errorf("Field expiringScale/ExpiringScale: value %v(%T) couldn't be cast to type *SingularityExpiringScale", value, value)

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		v, ok := value.(*SingularityExpiringSkipHealthchecks)
		if ok {
			self.ExpiringSkipHealthchecks = v
			return nil
		}
		return fmt.Errorf("Field expiringSkipHealthchecks/ExpiringSkipHealthchecks: value %v(%T) couldn't be cast to type *SingularityExpiringSkipHealthchecks", value, value)

	case "state", "State":
		v, ok := value.(SingularityRequestParentRequestState)
		if ok {
			self.State = &v
			return nil
		}
		return fmt.Errorf("Field state/State: value %v(%T) couldn't be cast to type SingularityRequestParentRequestState", value, value)

	case "activeDeploy", "ActiveDeploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.ActiveDeploy = v
			return nil
		}
		return fmt.Errorf("Field activeDeploy/ActiveDeploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "expiringPause", "ExpiringPause":
		v, ok := value.(*SingularityExpiringPause)
		if ok {
			self.ExpiringPause = v
			return nil
		}
		return fmt.Errorf("Field expiringPause/ExpiringPause: value %v(%T) couldn't be cast to type *SingularityExpiringPause", value, value)

	}
}

func (self *SingularityRequestParent) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "request", "Request":
		return self.Request, nil
		return nil, fmt.Errorf("Field Request no set on Request %+v", self)

	case "requestDeployState", "RequestDeployState":
		return self.RequestDeployState, nil
		return nil, fmt.Errorf("Field RequestDeployState no set on RequestDeployState %+v", self)

	case "pendingDeploy", "PendingDeploy":
		return self.PendingDeploy, nil
		return nil, fmt.Errorf("Field PendingDeploy no set on PendingDeploy %+v", self)

	case "pendingDeployState", "PendingDeployState":
		return self.PendingDeployState, nil
		return nil, fmt.Errorf("Field PendingDeployState no set on PendingDeployState %+v", self)

	case "expiringBounce", "ExpiringBounce":
		return self.ExpiringBounce, nil
		return nil, fmt.Errorf("Field ExpiringBounce no set on ExpiringBounce %+v", self)

	case "expiringScale", "ExpiringScale":
		return self.ExpiringScale, nil
		return nil, fmt.Errorf("Field ExpiringScale no set on ExpiringScale %+v", self)

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		return self.ExpiringSkipHealthchecks, nil
		return nil, fmt.Errorf("Field ExpiringSkipHealthchecks no set on ExpiringSkipHealthchecks %+v", self)

	case "state", "State":
		return *self.State, nil
		return nil, fmt.Errorf("Field State no set on State %+v", self)

	case "activeDeploy", "ActiveDeploy":
		return self.ActiveDeploy, nil
		return nil, fmt.Errorf("Field ActiveDeploy no set on ActiveDeploy %+v", self)

	case "expiringPause", "ExpiringPause":
		return self.ExpiringPause, nil
		return nil, fmt.Errorf("Field ExpiringPause no set on ExpiringPause %+v", self)

	}
}

func (self *SingularityRequestParent) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestParent", name)

	case "request", "Request":
		self.Request = nil

	case "requestDeployState", "RequestDeployState":
		self.RequestDeployState = nil

	case "pendingDeploy", "PendingDeploy":
		self.PendingDeploy = nil

	case "pendingDeployState", "PendingDeployState":
		self.PendingDeployState = nil

	case "expiringBounce", "ExpiringBounce":
		self.ExpiringBounce = nil

	case "expiringScale", "ExpiringScale":
		self.ExpiringScale = nil

	case "expiringSkipHealthchecks", "ExpiringSkipHealthchecks":
		self.ExpiringSkipHealthchecks = nil

	case "state", "State":
		self.State = nil

	case "activeDeploy", "ActiveDeploy":
		self.ActiveDeploy = nil

	case "expiringPause", "ExpiringPause":
		self.ExpiringPause = nil

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
