package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityRunNowRequest struct {
	present map[string]bool

	CommandLineArgs swaggering.StringList `json:"commandLineArgs"`

	Message string `json:"message,omitempty"`

	RunId string `json:"runId,omitempty"`

	SkipHealthchecks bool `json:"skipHealthchecks"`
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

func (self *SingularityRunNowRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityRunNowRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityRunNowRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityRunNowRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityRunNowRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "commandLineArgs", "CommandLineArgs":
		v, ok := value.(swaggering.StringList)
		if ok {
			self.CommandLineArgs = v
			self.present["commandLineArgs"] = true
			return nil
		} else {
			return fmt.Errorf("Field commandLineArgs/CommandLineArgs: value %v(%T) couldn't be cast to type StringList", value, value)
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

	case "runId", "RunId":
		v, ok := value.(string)
		if ok {
			self.RunId = v
			self.present["runId"] = true
			return nil
		} else {
			return fmt.Errorf("Field runId/RunId: value %v(%T) couldn't be cast to type string", value, value)
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

func (self *SingularityRunNowRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "commandLineArgs", "CommandLineArgs":
		if self.present != nil {
			if _, ok := self.present["commandLineArgs"]; ok {
				return self.CommandLineArgs, nil
			}
		}
		return nil, fmt.Errorf("Field CommandLineArgs no set on CommandLineArgs %+v", self)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "runId", "RunId":
		if self.present != nil {
			if _, ok := self.present["runId"]; ok {
				return self.RunId, nil
			}
		}
		return nil, fmt.Errorf("Field RunId no set on RunId %+v", self)

	case "skipHealthchecks", "SkipHealthchecks":
		if self.present != nil {
			if _, ok := self.present["skipHealthchecks"]; ok {
				return self.SkipHealthchecks, nil
			}
		}
		return nil, fmt.Errorf("Field SkipHealthchecks no set on SkipHealthchecks %+v", self)

	}
}

func (self *SingularityRunNowRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRunNowRequest", name)

	case "commandLineArgs", "CommandLineArgs":
		self.present["commandLineArgs"] = false

	case "message", "Message":
		self.present["message"] = false

	case "runId", "RunId":
		self.present["runId"] = false

	case "skipHealthchecks", "SkipHealthchecks":
		self.present["skipHealthchecks"] = false

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
