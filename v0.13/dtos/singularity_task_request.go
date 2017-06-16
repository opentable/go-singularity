package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskRequest struct {
	present map[string]bool

	Deploy *SingularityDeploy `json:"deploy"`

	PendingTask *SingularityPendingTask `json:"pendingTask"`

	Request *SingularityRequest `json:"request"`
}

func (self *SingularityTaskRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskRequest) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskRequest); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskRequest cannot copy the values from %#v", other)
}

func (self *SingularityTaskRequest) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskRequest) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			self.present["deploy"] = true
			return nil
		} else {
			return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)
		}

	case "pendingTask", "PendingTask":
		v, ok := value.(*SingularityPendingTask)
		if ok {
			self.PendingTask = v
			self.present["pendingTask"] = true
			return nil
		} else {
			return fmt.Errorf("Field pendingTask/PendingTask: value %v(%T) couldn't be cast to type *SingularityPendingTask", value, value)
		}

	case "request", "Request":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.Request = v
			self.present["request"] = true
			return nil
		} else {
			return fmt.Errorf("Field request/Request: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)
		}

	}
}

func (self *SingularityTaskRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "deploy", "Deploy":
		if self.present != nil {
			if _, ok := self.present["deploy"]; ok {
				return self.Deploy, nil
			}
		}
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "pendingTask", "PendingTask":
		if self.present != nil {
			if _, ok := self.present["pendingTask"]; ok {
				return self.PendingTask, nil
			}
		}
		return nil, fmt.Errorf("Field PendingTask no set on PendingTask %+v", self)

	case "request", "Request":
		if self.present != nil {
			if _, ok := self.present["request"]; ok {
				return self.Request, nil
			}
		}
		return nil, fmt.Errorf("Field Request no set on Request %+v", self)

	}
}

func (self *SingularityTaskRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "deploy", "Deploy":
		self.present["deploy"] = false

	case "pendingTask", "PendingTask":
		self.present["pendingTask"] = false

	case "request", "Request":
		self.present["request"] = false

	}

	return nil
}

func (self *SingularityTaskRequest) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskRequestList []*SingularityTaskRequest

func (self *SingularityTaskRequestList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskRequestList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskRequestList cannot copy the values from %#v", other)
}

func (list *SingularityTaskRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskRequestList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
