package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDeployUpdateDeployEventType string

const (
	SingularityDeployUpdateDeployEventTypeSTARTING SingularityDeployUpdateDeployEventType = "STARTING"
	SingularityDeployUpdateDeployEventTypeFINISHED SingularityDeployUpdateDeployEventType = "FINISHED"
)

type SingularityDeployUpdate struct {
	DeployMarker *SingularityDeployMarker                `json:"deployMarker,omitempty"`
	Deploy       *SingularityDeploy                      `json:"deploy,omitempty"`
	EventType    *SingularityDeployUpdateDeployEventType `json:"eventType,omitempty"`
	DeployResult *SingularityDeployResult                `json:"deployResult,omitempty"`
}

func (self *SingularityDeployUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDeployUpdate) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployUpdate); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployUpdate cannot copy the values from %#v", other)
}

func (self *SingularityDeployUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDeployUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDeployUpdate) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployUpdate", name)

	case "deployMarker", "DeployMarker":
		v, ok := value.(*SingularityDeployMarker)
		if ok {
			self.DeployMarker = v
			return nil
		}
		return fmt.Errorf("Field deployMarker/DeployMarker: value %v(%T) couldn't be cast to type *SingularityDeployMarker", value, value)

	case "deploy", "Deploy":
		v, ok := value.(*SingularityDeploy)
		if ok {
			self.Deploy = v
			return nil
		}
		return fmt.Errorf("Field deploy/Deploy: value %v(%T) couldn't be cast to type *SingularityDeploy", value, value)

	case "eventType", "EventType":
		v, ok := value.(SingularityDeployUpdateDeployEventType)
		if ok {
			self.EventType = &v
			return nil
		}
		return fmt.Errorf("Field eventType/EventType: value %v(%T) couldn't be cast to type SingularityDeployUpdateDeployEventType", value, value)

	case "deployResult", "DeployResult":
		v, ok := value.(*SingularityDeployResult)
		if ok {
			self.DeployResult = v
			return nil
		}
		return fmt.Errorf("Field deployResult/DeployResult: value %v(%T) couldn't be cast to type *SingularityDeployResult", value, value)

	}
}

func (self *SingularityDeployUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDeployUpdate", name)

	case "deployMarker", "DeployMarker":
		return self.DeployMarker, nil
		return nil, fmt.Errorf("Field DeployMarker no set on DeployMarker %+v", self)

	case "deploy", "Deploy":
		return self.Deploy, nil
		return nil, fmt.Errorf("Field Deploy no set on Deploy %+v", self)

	case "eventType", "EventType":
		return *self.EventType, nil
		return nil, fmt.Errorf("Field EventType no set on EventType %+v", self)

	case "deployResult", "DeployResult":
		return self.DeployResult, nil
		return nil, fmt.Errorf("Field DeployResult no set on DeployResult %+v", self)

	}
}

func (self *SingularityDeployUpdate) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDeployUpdate", name)

	case "deployMarker", "DeployMarker":
		self.DeployMarker = nil

	case "deploy", "Deploy":
		self.Deploy = nil

	case "eventType", "EventType":
		self.EventType = nil

	case "deployResult", "DeployResult":
		self.DeployResult = nil

	}

	return nil
}

func (self *SingularityDeployUpdate) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDeployUpdateList []*SingularityDeployUpdate

func (self *SingularityDeployUpdateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDeployUpdateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDeployUpdateList cannot copy the values from %#v", other)
}

func (list *SingularityDeployUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDeployUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDeployUpdateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
