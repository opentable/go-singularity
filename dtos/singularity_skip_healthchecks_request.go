package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularitySkipHealthchecksRequest struct {
	SkipHealthchecks *bool   `json:"skipHealthchecks,omitempty"`
	DurationMillis   *int64  `json:"durationMillis,omitempty"`
	ActionId         *string `json:"actionId,omitempty"`
	Message          *string `json:"message,omitempty"`
}

func (self *SingularitySkipHealthchecksRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularitySkipHealthchecksRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySkipHealthchecksRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySkipHealthchecksRequest cannot copy the values from %#v", other)
}

func (self *SingularitySkipHealthchecksRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularitySkipHealthchecksRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularitySkipHealthchecksRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySkipHealthchecksRequest", name)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

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

func (self *SingularitySkipHealthchecksRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularitySkipHealthchecksRequest", name)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

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

func (self *SingularitySkipHealthchecksRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySkipHealthchecksRequest", name)

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "durationMillis", "DurationMillis":
		self.DurationMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "message", "Message":
		self.Message = nil

	}

	return nil
}

func (self *SingularitySkipHealthchecksRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularitySkipHealthchecksRequestList []*SingularitySkipHealthchecksRequest

func (self *SingularitySkipHealthchecksRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularitySkipHealthchecksRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularitySkipHealthchecksRequestList cannot copy the values from %#v", other)
}

func (list *SingularitySkipHealthchecksRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularitySkipHealthchecksRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySkipHealthchecksRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
