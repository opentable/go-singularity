package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployResultDeployState string

const (
	SingularityDeployResultDeployStateSUCCEEDED             SingularityDeployResultDeployState = "SUCCEEDED"
	SingularityDeployResultDeployStateFAILED_INTERNAL_STATE SingularityDeployResultDeployState = "FAILED_INTERNAL_STATE"
	SingularityDeployResultDeployStateCANCELING             SingularityDeployResultDeployState = "CANCELING"
	SingularityDeployResultDeployStateWAITING               SingularityDeployResultDeployState = "WAITING"
	SingularityDeployResultDeployStateOVERDUE               SingularityDeployResultDeployState = "OVERDUE"
	SingularityDeployResultDeployStateFAILED                SingularityDeployResultDeployState = "FAILED"
	SingularityDeployResultDeployStateCANCELED              SingularityDeployResultDeployState = "CANCELED"
)

type SingularityDeployResult struct {
	present map[string]bool

	DeployFailures SingularityDeployFailureList `json:"deployFailures"`

	DeployState SingularityDeployResultDeployState `json:"deployState"`

	LbUpdate *SingularityLoadBalancerUpdate `json:"lbUpdate"`

	Message string `json:"message,omitempty"`

	Timestamp int64 `json:"timestamp"`
}

func (self *SingularityDeployResult) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployResult) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployResult); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployResult cannot copy the values from %#v", other)
}

func (self *SingularityDeployResult) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDeployResult) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployResult) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployResult) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDeployResult) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployFailures", "DeployFailures":
		v, ok := value.(SingularityDeployFailureList)
		if ok {
			self.DeployFailures = v
			self.present["deployFailures"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployFailures/DeployFailures: value %v(%T) couldn't be cast to type SingularityDeployFailureList", value, value)
		}

	case "deployState", "DeployState":
		v, ok := value.(SingularityDeployResultDeployState)
		if ok {
			self.DeployState = v
			self.present["deployState"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployState/DeployState: value %v(%T) couldn't be cast to type SingularityDeployResultDeployState", value, value)
		}

	case "lbUpdate", "LbUpdate":
		v, ok := value.(*SingularityLoadBalancerUpdate)
		if ok {
			self.LbUpdate = v
			self.present["lbUpdate"] = true
			return nil
		} else {
			return fmt.Errorf("Field lbUpdate/LbUpdate: value %v(%T) couldn't be cast to type *SingularityLoadBalancerUpdate", value, value)
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

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityDeployResult) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployFailures", "DeployFailures":
		if self.present != nil {
			if _, ok := self.present["deployFailures"]; ok {
				return self.DeployFailures, nil
			}
		}
		return nil, fmt.Errorf("Field DeployFailures no set on DeployFailures %+v", self)

	case "deployState", "DeployState":
		if self.present != nil {
			if _, ok := self.present["deployState"]; ok {
				return self.DeployState, nil
			}
		}
		return nil, fmt.Errorf("Field DeployState no set on DeployState %+v", self)

	case "lbUpdate", "LbUpdate":
		if self.present != nil {
			if _, ok := self.present["lbUpdate"]; ok {
				return self.LbUpdate, nil
			}
		}
		return nil, fmt.Errorf("Field LbUpdate no set on LbUpdate %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityDeployResult) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployFailures", "DeployFailures":
		self.present["deployFailures"] = false

	case "deployState", "DeployState":
		self.present["deployState"] = false

	case "lbUpdate", "LbUpdate":
		self.present["lbUpdate"] = false

	case "message", "Message":
		self.present["message"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityDeployResult) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployResultList []*SingularityDeployResult

func (self *SingularityDeployResultList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployResultList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployResultList cannot copy the values from %#v", other)
}

func (list *SingularityDeployResultList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployResultList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployResultList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
