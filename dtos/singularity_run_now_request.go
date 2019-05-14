package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRunNowRequest struct {
	Message          *string                `json:"message,omitempty"`
	RunId            *string                `json:"runId,omitempty"`
	CommandLineArgs  *swaggering.StringList `json:"commandLineArgs,omitempty"`
	SkipHealthchecks *bool                  `json:"skipHealthchecks,omitempty"`
	Resources        *Resources             `json:"resources,omitempty"`
}

func (self *SingularityRunNowRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityRunNowRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRunNowRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRunNowRequest cannot copy the values from %#v", other)
}

func (self *SingularityRunNowRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRunNowRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRunNowRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = &v
			return nil
		}
		return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)

	case "commandLineArgs", "CommandLineArgs":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CommandLineArgs = &v
			return nil
		}
		return fmt.Errorf("Field commandLineArgs/CommandLineArgs: value %v(%T) couldn't be cast to type swaggering.StringList", value, value)

	case "skipHealthchecks", "SkipHealthchecks":
		v, ok := value.(bool)
		if ok {
			self.SkipHealthchecks = &v
			return nil
		}
		return fmt.Errorf("Field skipHealthchecks/SkipHealthchecks: value %v(%T) couldn't be cast to type bool", value, value)

	case "resources", "Resources":
		v, ok := value.(*Resources)
		if ok {
			self.Resources = v
			return nil
		}
		return fmt.Errorf("Field resources/Resources: value %v(%T) couldn't be cast to type *Resources", value, value)

	}
}

func (self *SingularityRunNowRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "runId", "RunId":
		return *self.RunId, nil
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "commandLineArgs", "CommandLineArgs":
		return *self.CommandLineArgs, nil
		return nil, fmt.Errorf("Field CommandLineArgs no set on CommandLineArgs %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		return *self.SkipHealthchecks, nil
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	case "resources", "Resources":
		return self.Resources, nil
		return nil, fmt.Errorf("Field Resources no set on Resources %+v", self)

	}
}

func (self *SingularityRunNowRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "message", "Message":
		self.Message = nil

	case "runId", "RunId":
		self.RunId = nil

	case "commandLineArgs", "CommandLineArgs":
		self.CommandLineArgs = nil

	case "skipHealthchecks", "SkipHealthchecks":
		self.SkipHealthchecks = nil

	case "resources", "Resources":
		self.Resources = nil

	}

	return nil
}

func (self *SingularityRunNowRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityRunNowRequestList []*SingularityRunNowRequest

func (self *SingularityRunNowRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityRunNowRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityRunNowRequestList cannot copy the values from %#v", other)
}

func (list *SingularityRunNowRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityRunNowRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRunNowRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
