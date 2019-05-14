package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type HealthcheckOptionsHealthcheckProtocol string

const (
	HealthcheckOptionsHealthcheckProtocolhttp  HealthcheckOptionsHealthcheckProtocol = "http"
	HealthcheckOptionsHealthcheckProtocolhttps HealthcheckOptionsHealthcheckProtocol = "https"
)

type HealthcheckOptions struct {
	StartupDelaySeconds    *int32                                 `json:"startupDelaySeconds,omitempty"`
	StartupIntervalSeconds *int32                                 `json:"startupIntervalSeconds,omitempty"`
	IntervalSeconds        *int32                                 `json:"intervalSeconds,omitempty"`
	ResponseTimeoutSeconds *int32                                 `json:"responseTimeoutSeconds,omitempty"`
	MaxRetries             *int32                                 `json:"maxRetries,omitempty"`
	Protocol               *HealthcheckOptionsHealthcheckProtocol `json:"protocol,omitempty"`
	StartupTimeoutSeconds  *int32                                 `json:"startupTimeoutSeconds,omitempty"`
	PortNumber             *int64                                 `json:"portNumber,omitempty"`
	FailureStatusCodes     *[]int32                               `json:"failureStatusCodes,omitempty"`
	Uri                    *string                                `json:"uri,omitempty"`
	PortIndex              *int32                                 `json:"portIndex,omitempty"`
}

func (self *HealthcheckOptions) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *HealthcheckOptions) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthcheckOptions); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthcheckOptions cannot copy the values from %#v", other)
}

func (self *HealthcheckOptions) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HealthcheckOptions) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HealthcheckOptions) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "startupDelaySeconds", "StartupDelaySeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupDelaySeconds = &v
			return nil
		}
		return fmt.Errorf("Field startupDelaySeconds/StartupDelaySeconds: value %v(%T) couldn't be cast to type int32", value, value)

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupIntervalSeconds = &v
			return nil
		}
		return fmt.Errorf("Field startupIntervalSeconds/StartupIntervalSeconds: value %v(%T) couldn't be cast to type int32", value, value)

	case "intervalSeconds", "IntervalSeconds":
		v, ok := value.(int32)
		if ok {
			self.IntervalSeconds = &v
			return nil
		}
		return fmt.Errorf("Field intervalSeconds/IntervalSeconds: value %v(%T) couldn't be cast to type int32", value, value)

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		v, ok := value.(int32)
		if ok {
			self.ResponseTimeoutSeconds = &v
			return nil
		}
		return fmt.Errorf("Field responseTimeoutSeconds/ResponseTimeoutSeconds: value %v(%T) couldn't be cast to type int32", value, value)

	case "maxRetries", "MaxRetries":
		v, ok := value.(int32)
		if ok {
			self.MaxRetries = &v
			return nil
		}
		return fmt.Errorf("Field maxRetries/MaxRetries: value %v(%T) couldn't be cast to type int32", value, value)

	case "protocol", "Protocol":
		v, ok := value.(HealthcheckOptionsHealthcheckProtocol)
		if ok {
			self.Protocol = &v
			return nil
		}
		return fmt.Errorf("Field protocol/Protocol: value %v(%T) couldn't be cast to type HealthcheckOptionsHealthcheckProtocol", value, value)

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupTimeoutSeconds = &v
			return nil
		}
		return fmt.Errorf("Field startupTimeoutSeconds/StartupTimeoutSeconds: value %v(%T) couldn't be cast to type int32", value, value)

	case "portNumber", "PortNumber":
		v, ok := value.(int64)
		if ok {
			self.PortNumber = &v
			return nil
		}
		return fmt.Errorf("Field portNumber/PortNumber: value %v(%T) couldn't be cast to type int64", value, value)

	case "failureStatusCodes", "FailureStatusCodes":
		v, ok := value.([]int32)
		if ok {
			self.FailureStatusCodes = &v
			return nil
		}
		return fmt.Errorf("Field failureStatusCodes/FailureStatusCodes: value %v(%T) couldn't be cast to type []int32", value, value)

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = &v
			return nil
		}
		return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)

	case "portIndex", "PortIndex":
		v, ok := value.(int32)
		if ok {
			self.PortIndex = &v
			return nil
		}
		return fmt.Errorf("Field portIndex/PortIndex: value %v(%T) couldn't be cast to type int32", value, value)

	}
}

func (self *HealthcheckOptions) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "startupDelaySeconds", "StartupDelaySeconds":
		return *self.StartupDelaySeconds, nil
		return nil, fmt.Errorf("Field StartupDelaySeconds no set on StartupDelaySeconds %+v", self)

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		return *self.StartupIntervalSeconds, nil
		return nil, fmt.Errorf("Field StartupIntervalSeconds no set on StartupIntervalSeconds %+v", self)

	case "intervalSeconds", "IntervalSeconds":
		return *self.IntervalSeconds, nil
		return nil, fmt.Errorf("Field IntervalSeconds no set on IntervalSeconds %+v", self)

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		return *self.ResponseTimeoutSeconds, nil
		return nil, fmt.Errorf("Field ResponseTimeoutSeconds no set on ResponseTimeoutSeconds %+v", self)

	case "maxRetries", "MaxRetries":
		return *self.MaxRetries, nil
		return nil, fmt.Errorf("Field MaxRetries no set on MaxRetries %+v", self)

	case "protocol", "Protocol":
		return *self.Protocol, nil
		return nil, fmt.Errorf("Field Protocol no set on Protocol %+v", self)

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		return *self.StartupTimeoutSeconds, nil
		return nil, fmt.Errorf("Field StartupTimeoutSeconds no set on StartupTimeoutSeconds %+v", self)

	case "portNumber", "PortNumber":
		return *self.PortNumber, nil
		return nil, fmt.Errorf("Field PortNumber no set on PortNumber %+v", self)

	case "failureStatusCodes", "FailureStatusCodes":
		return *self.FailureStatusCodes, nil
		return nil, fmt.Errorf("Field FailureStatusCodes no set on FailureStatusCodes %+v", self)

	case "uri", "Uri":
		return *self.Uri, nil
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "portIndex", "PortIndex":
		return *self.PortIndex, nil
		return nil, fmt.Errorf("Field PortIndex no set on PortIndex %+v", self)

	}
}

func (self *HealthcheckOptions) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "startupDelaySeconds", "StartupDelaySeconds":
		self.StartupDelaySeconds = nil

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		self.StartupIntervalSeconds = nil

	case "intervalSeconds", "IntervalSeconds":
		self.IntervalSeconds = nil

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		self.ResponseTimeoutSeconds = nil

	case "maxRetries", "MaxRetries":
		self.MaxRetries = nil

	case "protocol", "Protocol":
		self.Protocol = nil

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		self.StartupTimeoutSeconds = nil

	case "portNumber", "PortNumber":
		self.PortNumber = nil

	case "failureStatusCodes", "FailureStatusCodes":
		self.FailureStatusCodes = nil

	case "uri", "Uri":
		self.Uri = nil

	case "portIndex", "PortIndex":
		self.PortIndex = nil

	}

	return nil
}

func (self *HealthcheckOptions) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type HealthcheckOptionsList []*HealthcheckOptions

func (self *HealthcheckOptionsList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthcheckOptionsList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthcheckOptionsList cannot copy the values from %#v", other)
}

func (list *HealthcheckOptionsList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *HealthcheckOptionsList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HealthcheckOptionsList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
