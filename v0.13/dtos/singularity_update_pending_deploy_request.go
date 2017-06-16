package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityUpdatePendingDeployRequest struct {
	present map[string]bool

	DeployId string `json:"deployId,omitempty"`

	RequestId string `json:"requestId,omitempty"`

	TargetActiveInstances int32 `json:"targetActiveInstances"`
}

func (self *SingularityUpdatePendingDeployRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityUpdatePendingDeployRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityUpdatePendingDeployRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityUpdatePendingDeployRequest cannot copy the values from %#v", other)
}

func (self *SingularityUpdatePendingDeployRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityUpdatePendingDeployRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityUpdatePendingDeployRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityUpdatePendingDeployRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityUpdatePendingDeployRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityUpdatePendingDeployRequest", name)

	case "deployId", "DeployId":
		v, ok := value.(string)
		if ok {
			self.DeployId = v
			self.present["deployId"] = true
			return nil
		} else {
			return fmt.Errorf("Field deployId/DeployId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "requestId", "RequestId":
		v, ok := value.(string)
		if ok {
			self.RequestId = v
			self.present["requestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field requestId/RequestId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "targetActiveInstances", "TargetActiveInstances":
		v, ok := value.(int32)
		if ok {
			self.TargetActiveInstances = v
			self.present["targetActiveInstances"] = true
			return nil
		} else {
			return fmt.Errorf("Field targetActiveInstances/TargetActiveInstances: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *SingularityUpdatePendingDeployRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityUpdatePendingDeployRequest", name)

	case "deployId", "DeployId":
		if self.present != nil {
			if _, ok := self.present["deployId"]; ok {
				return self.DeployId, nil
			}
		}
		return nil, fmt.Errorf("Field DeployId no set on DeployId %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	case "targetActiveInstances", "TargetActiveInstances":
		if self.present != nil {
			if _, ok := self.present["targetActiveInstances"]; ok {
				return self.TargetActiveInstances, nil
			}
		}
		return nil, fmt.Errorf("Field TargetActiveInstances no set on TargetActiveInstances %+v", self)

	}
}

func (self *SingularityUpdatePendingDeployRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityUpdatePendingDeployRequest", name)

	case "deployId", "DeployId":
		self.present["deployId"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	case "targetActiveInstances", "TargetActiveInstances":
		self.present["targetActiveInstances"] = false

	}

	return nil
}

func (self *SingularityUpdatePendingDeployRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityUpdatePendingDeployRequestList []*SingularityUpdatePendingDeployRequest

func (self *SingularityUpdatePendingDeployRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityUpdatePendingDeployRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityUpdatePendingDeployRequestList cannot copy the values from %#v", other)
}

func (list *SingularityUpdatePendingDeployRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityUpdatePendingDeployRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityUpdatePendingDeployRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
