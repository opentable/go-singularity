package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityPendingDeployDeployState string

const (
	SingularityPendingDeployDeployStateSUCCEEDED             SingularityPendingDeployDeployState = "SUCCEEDED"
	SingularityPendingDeployDeployStateFAILED_INTERNAL_STATE SingularityPendingDeployDeployState = "FAILED_INTERNAL_STATE"
	SingularityPendingDeployDeployStateCANCELING             SingularityPendingDeployDeployState = "CANCELING"
	SingularityPendingDeployDeployStateWAITING               SingularityPendingDeployDeployState = "WAITING"
	SingularityPendingDeployDeployStateOVERDUE               SingularityPendingDeployDeployState = "OVERDUE"
	SingularityPendingDeployDeployStateFAILED                SingularityPendingDeployDeployState = "FAILED"
	SingularityPendingDeployDeployStateCANCELED              SingularityPendingDeployDeployState = "CANCELED"
)

type SingularityPendingDeploy struct {
	present map[string]bool

	CurrentDeployState SingularityPendingDeployDeployState `json:"currentDeployState"`

	DeployMarker *SingularityDeployMarker `json:"deployMarker"`

	DeployProgress *SingularityDeployProgress `json:"deployProgress"`

	LastLoadBalancerUpdate *SingularityLoadBalancerUpdate `json:"lastLoadBalancerUpdate"`
}

func (self *SingularityPendingDeploy) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityPendingDeploy) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingDeploy); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingDeploy cannot copy the values from %#v", other)
}

func (self *SingularityPendingDeploy) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityPendingDeploy) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityPendingDeploy) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityPendingDeploy) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityPendingDeploy) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "currentDeployState", "CurrentDeployState":
		v, ok := value.(SingularityPendingDeployDeployState)
		if ok {
			self.CurrentDeployState = v
			self.present["currentDeployState"] = true
			return nil
		} else {
			return fmt.Errorf("Field currentDeployState/CurrentDeployState: value %v(%T) couldn't be cast to type SingularityPendingDeployDeployState", value, value)
		}

	case "deployMarker", "DeployMarker":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.DeployMarker = v
			self.present["deployMarker"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployMarker/DeployMarker: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)
		}

	case "deployProgress", "DeployProgress":
		v, ok := value.(*SingularityDeployProgress)
		if ok {
			self.DeployProgress = v
			self.present["deployProgress"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployProgress/DeployProgress: value %v(%T) couldn't be cast to type *SingularityDeployProgress", value, value)
		}

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		v, ok := value.(*SingularityLoadBalancerUpdate)
		if ok {
			self.LastLoadBalancerUpdate = v
			self.present["lastLoadBalancerUpdate"] = true
			return nil
		} else {
			return fmt.Errorf("Field lastLoadBalancerUpdate/LastLoadBalancerUpdate: value %v(%T) couldn't be cast to type *SingularityLoadBalancerUpdate", value, value)
		}

	}
}

func (self *SingularityPendingDeploy) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "currentDeployState", "CurrentDeployState":
		if self.present != nil {
			if _, ok := self.present["currentDeployState"]; ok {
				return self.CurrentDeployState, nil
			}
		}
		return nil, fmt.Errorf("Field CurrentDeployState no set on CurrentDeployState %+v", self)

	case "deployMarker", "DeployMarker":
		if self.present != nil {
			if _, ok := self.present["deployMarker"]; ok {
				return self.DeployMarker, nil
			}
		}
		return nil, fmt.Errorf("Field DeployMarker no set on DeployMarker %+v", self)

	case "deployProgress", "DeployProgress":
		if self.present != nil {
			if _, ok := self.present["deployProgress"]; ok {
				return self.DeployProgress, nil
			}
		}
		return nil, fmt.Errorf("Field DeployProgress no set on DeployProgress %+v", self)

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		if self.present != nil {
			if _, ok := self.present["lastLoadBalancerUpdate"]; ok {
				return self.LastLoadBalancerUpdate, nil
			}
		}
		return nil, fmt.Errorf("Field LastLoadBalancerUpdate no set on LastLoadBalancerUpdate %+v", self)

	}
}

func (self *SingularityPendingDeploy) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityPendingDeploy", name)

	case "currentDeployState", "CurrentDeployState":
		self.present["currentDeployState"] = false

	case "deployMarker", "DeployMarker":
		self.present["deployMarker"] = false

	case "deployProgress", "DeployProgress":
		self.present["deployProgress"] = false

	case "lastLoadBalancerUpdate", "LastLoadBalancerUpdate":
		self.present["lastLoadBalancerUpdate"] = false

	}

	return nil
}

func (self *SingularityPendingDeploy) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityPendingDeployList []*SingularityPendingDeploy

func (self *SingularityPendingDeployList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityPendingDeployList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityPendingDeployList cannot copy the values from %#v", other)
}

func (list *SingularityPendingDeployList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityPendingDeployList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityPendingDeployList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
