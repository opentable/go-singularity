package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskRequest struct {
	Request     *SingularityRequest     `json:"request,omitempty"`
	Deploy      *SingularityDeploy      `json:"deploy,omitempty"`
	PendingTask *SingularityPendingTask `json:"pendingTask,omitempty"`
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

func (self *SingularityTaskRequest) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskRequest) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskRequest) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "request", "Request":
		v, ok := value.(*SingularityRequest)
		if ok {
			self.Request = v
			return nil
		}
		return fmt.Errorf("Field request/Request: value %v(%T) couldn't be cast to type *SingularityRequest", value, value)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			return nil
		}
		return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "pendingTask", "PendingTask":
		v, ok := value.(*SingularityPendingTask)
		if ok {
			self.PendingTask = v
			return nil
		}
		return fmt.Errorf("Field pendingTask/PendingTask: value %v(%T) couldn't be cast to type *SingularityPendingTask", value, value)

	}
}

func (self *SingularityTaskRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "request", "Request":
		return self.Request, nil
		return nil, fmt.Errorf("Field Request no set on Request %+v", self)

	case "deploy", "Deploy":
		return self.Deploy, nil
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "pendingTask", "PendingTask":
		return self.PendingTask, nil
		return nil, fmt.Errorf("Field PendingTask no set on PendingTask %+v", self)

	}
}

func (self *SingularityTaskRequest) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskRequest", name)

	case "request", "Request":
		self.Request = nil

	case "deploy", "Deploy":
		self.Deploy = nil

	case "pendingTask", "PendingTask":
		self.PendingTask = nil

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
