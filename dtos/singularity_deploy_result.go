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
	DeployState    *SingularityDeployResultDeployState `json:"deployState,omitempty"`
	LbUpdate       *SingularityLoadBalancerUpdate      `json:"lbUpdate,omitempty"`
	Message        *string                             `json:"message,omitempty"`
	DeployFailures *SingularityDeployFailureList       `json:"deployFailures,omitempty"`
	Timestamp      *int64                              `json:"timestamp,omitempty"`
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

func (self *SingularityDeployResult) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployResult) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployResult) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployState", "DeployState":
		v, ok := value.(SingularityDeployResultDeployState)
		if ok {
			self.DeployState = &v
			return nil
		}
		return fmt.Errorf("Field deployState/DeployState: value %v(%T) couldn't be cast to type SingularityDeployResultDeployState", value, value)

	case "lbUpdate", "LbUpdate":
		v, ok := value.(*SingularityLoadBalancerUpdate)
		if ok {
			self.LbUpdate = v
			return nil
		}
		return fmt.Errorf("Field lbUpdate/LbUpdate: value %v(%T) couldn't be cast to type *SingularityLoadBalancerUpdate", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "deployFailures", "DeployFailures":
		v, ok := value.(SingularityDeployFailureList)
		if ok {
			self.DeployFailures = &v
			return nil
		}
		return fmt.Errorf("Field deployFailures/DeployFailures: value %v(%T) couldn't be cast to type SingularityDeployFailureList", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityDeployResult) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployState", "DeployState":
		return *self.DeployState, nil
		return nil, fmt.Errorf("Field DeployState no set on DeployState %+v", self)

	case "lbUpdate", "LbUpdate":
		return self.LbUpdate, nil
		return nil, fmt.Errorf("Field LbUpdate no set on LbUpdate %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "deployFailures", "DeployFailures":
		return *self.DeployFailures, nil
		return nil, fmt.Errorf("Field DeployFailures no set on DeployFailures %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityDeployResult) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployResult", name)

	case "deployState", "DeployState":
		self.DeployState = nil

	case "lbUpdate", "LbUpdate":
		self.LbUpdate = nil

	case "message", "Message":
		self.Message = nil

	case "deployFailures", "DeployFailures":
		self.DeployFailures = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

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
