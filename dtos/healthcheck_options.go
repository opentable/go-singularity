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
	present map[string]bool

	PortNumber int64 `json:"portNumber"`

	StartupTimeoutSeconds int32 `json:"startupTimeoutSeconds"`

	IntervalSeconds int32 `json:"intervalSeconds"`

	FailureStatusCodes []int32 `json:"failureStatusCodes"`

	MaxRetries int32 `json:"maxRetries"`

	Uri string `json:"uri,omitempty"`

	PortIndex int32 `json:"portIndex"`

	Protocol HealthcheckOptionsHealthcheckProtocol `json:"protocol"`

	StartupDelaySeconds int32 `json:"startupDelaySeconds"`

	StartupIntervalSeconds int32 `json:"startupIntervalSeconds"`

	ResponseTimeoutSeconds int32 `json:"responseTimeoutSeconds"`
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

func (self *HealthcheckOptions) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *HealthcheckOptions) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HealthcheckOptions) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HealthcheckOptions) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *HealthcheckOptions) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "portNumber", "PortNumber":
		v, ok := value.(int64)
		if ok {
			self.PortNumber = v
			self.present["portNumber"] = true
			return nil
		} else {
			return fmt.Errorf("Field portNumber/PortNumber: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupTimeoutSeconds = v
			self.present["startupTimeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field startupTimeoutSeconds/StartupTimeoutSeconds: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "intervalSeconds", "IntervalSeconds":
		v, ok := value.(int32)
		if ok {
			self.IntervalSeconds = v
			self.present["intervalSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field intervalSeconds/IntervalSeconds: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "failureStatusCodes", "FailureStatusCodes":
		v, ok := value.([]int32)
		if ok {
			self.FailureStatusCodes = v
			self.present["failureStatusCodes"] = true
			return nil
		} else {
			return fmt.Errorf("Field failureStatusCodes/FailureStatusCodes: value %v(%T) couldn't be cast to type []int32", value, value)
		}

	case "maxRetries", "MaxRetries":
		v, ok := value.(int32)
		if ok {
			self.MaxRetries = v
			self.present["maxRetries"] = true
			return nil
		} else {
			return fmt.Errorf("Field maxRetries/MaxRetries: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "uri", "Uri":
		v, ok := value.(string)
		if ok {
			self.Uri = v
			self.present["uri"] = true
			return nil
		} else {
			return fmt.Errorf("Field uri/Uri: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "portIndex", "PortIndex":
		v, ok := value.(int32)
		if ok {
			self.PortIndex = v
			self.present["portIndex"] = true
			return nil
		} else {
			return fmt.Errorf("Field portIndex/PortIndex: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "protocol", "Protocol":
		v, ok := value.(HealthcheckOptionsHealthcheckProtocol)
		if ok {
			self.Protocol = v
			self.present["protocol"] = true
			return nil
		} else {
			return fmt.Errorf("Field protocol/Protocol: value %v(%T) couldn't be cast to type HealthcheckOptionsHealthcheckProtocol", value, value)
		}

	case "startupDelaySeconds", "StartupDelaySeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupDelaySeconds = v
			self.present["startupDelaySeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field startupDelaySeconds/StartupDelaySeconds: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		v, ok := value.(int32)
		if ok {
			self.StartupIntervalSeconds = v
			self.present["startupIntervalSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field startupIntervalSeconds/StartupIntervalSeconds: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		v, ok := value.(int32)
		if ok {
			self.ResponseTimeoutSeconds = v
			self.present["responseTimeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field responseTimeoutSeconds/ResponseTimeoutSeconds: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *HealthcheckOptions) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "portNumber", "PortNumber":
		if self.present != nil {
			if _, ok := self.present["portNumber"]; ok {
				return self.PortNumber, nil
			}
		}
		return nil, fmt.Errorf("Field PortNumber no set on PortNumber %+v", self)

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["startupTimeoutSeconds"]; ok {
				return self.StartupTimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field StartupTimeoutSeconds no set on StartupTimeoutSeconds %+v", self)

	case "intervalSeconds", "IntervalSeconds":
		if self.present != nil {
			if _, ok := self.present["intervalSeconds"]; ok {
				return self.IntervalSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field IntervalSeconds no set on IntervalSeconds %+v", self)

	case "failureStatusCodes", "FailureStatusCodes":
		if self.present != nil {
			if _, ok := self.present["failureStatusCodes"]; ok {
				return self.FailureStatusCodes, nil
			}
		}
		return nil, fmt.Errorf("Field FailureStatusCodes no set on FailureStatusCodes %+v", self)

	case "maxRetries", "MaxRetries":
		if self.present != nil {
			if _, ok := self.present["maxRetries"]; ok {
				return self.MaxRetries, nil
			}
		}
		return nil, fmt.Errorf("Field MaxRetries no set on MaxRetries %+v", self)

	case "uri", "Uri":
		if self.present != nil {
			if _, ok := self.present["uri"]; ok {
				return self.Uri, nil
			}
		}
		return nil, fmt.Errorf("Field Uri no set on Uri %+v", self)

	case "portIndex", "PortIndex":
		if self.present != nil {
			if _, ok := self.present["portIndex"]; ok {
				return self.PortIndex, nil
			}
		}
		return nil, fmt.Errorf("Field PortIndex no set on PortIndex %+v", self)

	case "protocol", "Protocol":
		if self.present != nil {
			if _, ok := self.present["protocol"]; ok {
				return self.Protocol, nil
			}
		}
		return nil, fmt.Errorf("Field Protocol no set on Protocol %+v", self)

	case "startupDelaySeconds", "StartupDelaySeconds":
		if self.present != nil {
			if _, ok := self.present["startupDelaySeconds"]; ok {
				return self.StartupDelaySeconds, nil
			}
		}
		return nil, fmt.Errorf("Field StartupDelaySeconds no set on StartupDelaySeconds %+v", self)

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		if self.present != nil {
			if _, ok := self.present["startupIntervalSeconds"]; ok {
				return self.StartupIntervalSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field StartupIntervalSeconds no set on StartupIntervalSeconds %+v", self)

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["responseTimeoutSeconds"]; ok {
				return self.ResponseTimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field ResponseTimeoutSeconds no set on ResponseTimeoutSeconds %+v", self)

	}
}

func (self *HealthcheckOptions) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthcheckOptions", name)

	case "portNumber", "PortNumber":
		self.present["portNumber"] = false

	case "startupTimeoutSeconds", "StartupTimeoutSeconds":
		self.present["startupTimeoutSeconds"] = false

	case "intervalSeconds", "IntervalSeconds":
		self.present["intervalSeconds"] = false

	case "failureStatusCodes", "FailureStatusCodes":
		self.present["failureStatusCodes"] = false

	case "maxRetries", "MaxRetries":
		self.present["maxRetries"] = false

	case "uri", "Uri":
		self.present["uri"] = false

	case "portIndex", "PortIndex":
		self.present["portIndex"] = false

	case "protocol", "Protocol":
		self.present["protocol"] = false

	case "startupDelaySeconds", "StartupDelaySeconds":
		self.present["startupDelaySeconds"] = false

	case "startupIntervalSeconds", "StartupIntervalSeconds":
		self.present["startupIntervalSeconds"] = false

	case "responseTimeoutSeconds", "ResponseTimeoutSeconds":
		self.present["responseTimeoutSeconds"] = false

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
