package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityScaleRequest struct {
	present map[string]bool

	ActionId string `json:"actionId,omitempty"`

	DurationMillis int64 `json:"durationMillis"`

	Instances int32 `json:"instances"`

	Message string `json:"message,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`
}

func (self *SingularityScaleRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityScaleRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityScaleRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityScaleRequest cannot copy the values from %#v", other)
}

func (self *SingularityScaleRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityScaleRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityScaleRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityScaleRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityScaleRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = v
			self.present["actionId"] = true
			return nil
		} else {
			return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = v
			self.present["durationMillis"] = true
			return nil
		} else {
			return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "instances", "Instances":
		v, ok := value.(int32)
		if ok {
			self.Instances = v
			self.present["instances"] = true
			return nil
		} else {
			return fmt.Errorf("Field instances/Instances: value %v(%T) couldn't be cast to type int32", value, value)
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

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = v
			self.present["skipHealthchecks"] = true
			return nil
		} else {
			return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)
		}

	}
}

func (self *SingularityScaleRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "actionId", "ActionId":
		if self.present != nil {
			if _, ok := self.present["actionId"]; ok {
				return self.ActionId, nil
			}
		}
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "durationMillis", "DurationMillis":
		if self.present != nil {
			if _, ok := self.present["durationMillis"]; ok {
				return self.DurationMillis, nil
			}
		}
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "instances", "Instances":
		if self.present != nil {
			if _, ok := self.present["instances"]; ok {
				return self.Instances, nil
			}
		}
		return nil, fmt.Errorf("Field Instances no set on Instances %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityScaleRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "actionId", "ActionId":
		self.present["actionId"] = false

	case "durationMillis", "DurationMillis":
		self.present["durationMillis"] = false

	case "instances", "Instances":
		self.present["instances"] = false

	case "message", "Message":
		self.present["message"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

	}

	return nil
}

func (self *SingularityScaleRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityScaleRequestList []*SingularityScaleRequest

func (self *SingularityScaleRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityScaleRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityScaleRequestList cannot copy the values from %#v", other)
}

func (list *SingularityScaleRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityScaleRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityScaleRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
