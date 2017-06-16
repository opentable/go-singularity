package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRequestDeployState struct {
	present map[string]bool

	ActiveDeploy *SingularityDeployMarker `json:"activeDeploy"`

	PendingDeploy *SingularityDeployMarker `json:"pendingDeploy"`

	RequestId string `json:"requestId,omitempty"`
}

func (self *SingularityRequestDeployState) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRequestDeployState) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestDeployState); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestDeployState cannot copy the values from %#v", other)
}

func (self *SingularityRequestDeployState) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRequestDeployState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRequestDeployState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRequestDeployState) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRequestDeployState) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestDeployState", name)

	case "activeDeploy", "ActiveDeploy":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.ActiveDeploy = v
			self.present["activeDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field activeDeploy/ActiveDeploy: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)
		}

	case "pendingDeploy", "PendingDeploy":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.PendingDeploy = v
			self.present["pendingDeploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingDeploy/PendingDeploy: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)
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

	}
}

func (self *SingularityRequestDeployState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRequestDeployState", name)

	case "activeDeploy", "ActiveDeploy":
		if self.present != nil {
			if _, ok := self.present["activeDeploy"]; ok {
				return self.ActiveDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field ActiveDeploy no set on ActiveDeploy %+v", self)

	case "pendingDeploy", "PendingDeploy":
		if self.present != nil {
			if _, ok := self.present["pendingDeploy"]; ok {
				return self.PendingDeploy, nil
			}
		}
		return nil, fmt.Errorf("Field PendingDeploy no set on PendingDeploy %+v", self)

	case "requestId", "RequestId":
		if self.present != nil {
			if _, ok := self.present["requestId"]; ok {
				return self.RequestId, nil
			}
		}
		return nil, fmt.Errorf("Field RequestId no set on RequestId %+v", self)

	}
}

func (self *SingularityRequestDeployState) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRequestDeployState", name)

	case "activeDeploy", "ActiveDeploy":
		self.present["activeDeploy"] = false

	case "pendingDeploy", "PendingDeploy":
		self.present["pendingDeploy"] = false

	case "requestId", "RequestId":
		self.present["requestId"] = false

	}

	return nil
}

func (self *SingularityRequestDeployState) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRequestDeployStateList []*SingularityRequestDeployState

func (self *SingularityRequestDeployStateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRequestDeployStateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRequestDeployStateList cannot copy the values from %#v", other)
}

func (list *SingularityRequestDeployStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRequestDeployStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRequestDeployStateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
