package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityScaleRequest struct {
	Instances        *int32  `json:"instances,omitempty"`
	SkipHealthchecks *bool   `json:"skipHealthchecks,omitempty"`
	Bounce           *bool   `json:"bounce,omitempty"`
	Incremental      *bool   `json:"incremental,omitempty"`
	DurationMillis   *int64  `json:"durationMillis,omitempty"`
	ActionId         *string `json:"actionId,omitempty"`
	Message          *string `json:"message,omitempty"`
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

func (self *SingularityScaleRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityScaleRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityScaleRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "instances", "Instances":
		v, ok := value.(int32)
		if ok {
			self.Instances = &v
			return nil
		}
		return fmt.Errorf("Field instances/Instances: value %v(%T) couldn't be cast to type int32", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "bounce", "Bounce":
		v, ok := value.(bool)
		if ok {
			self.Bounce = &v
			return nil
		}
		return fmt.Errorf("Field bounce/Bounce: value %v(%T) couldn't be cast to type bool", value, value)

	case "incremental", "Incremental":
		v, ok := value.(bool)
		if ok {
			self.Incremental = &v
			return nil
		}
		return fmt.Errorf("Field incremental/Incremental: value %v(%T) couldn't be cast to type bool", value, value)

	case "durationMillis", "DurationMillis":
		v, ok := value.(int64)
		if ok {
			self.DurationMillis = &v
			return nil
		}
		return fmt.Errorf("Field durationMillis/DurationMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	}
}

func (self *SingularityScaleRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "instances", "Instances":
		return *self.Instances, nil
		return nil, fmt.Errorf("Field Instances no set on Instances %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "bounce", "Bounce":
		return *self.Bounce, nil
		return nil, fmt.Errorf("Field Bounce no set on Bounce %+v", self)

	case "incremental", "Incremental":
		return *self.Incremental, nil
		return nil, fmt.Errorf("Field Incremental no set on Incremental %+v", self)

	case "durationMillis", "DurationMillis":
		return *self.DurationMillis, nil
		return nil, fmt.Errorf("Field DurationMillis no set on DurationMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityScaleRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityScaleRequest", name)

	case "instances", "Instances":
		self.Instances = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "bounce", "Bounce":
		self.Bounce = nil

	case "incremental", "Incremental":
		self.Incremental = nil

	case "durationMillis", "DurationMillis":
		self.DurationMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "message", "Message":
		self.Message = nil

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
