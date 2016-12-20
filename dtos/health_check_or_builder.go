package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type HealthCheckOrBuilder struct {
	present map[string]bool

	Command *CommandInfo `json:"command"`

	CommandOrBuilder *CommandInfoOrBuilder `json:"commandOrBuilder"`

	ConsecutiveFailures int32 `json:"consecutiveFailures"`

	DelaySeconds float64 `json:"delaySeconds"`

	GracePeriodSeconds float64 `json:"gracePeriodSeconds"`

	Http *HTTP `json:"http"`

	HttpOrBuilder *HTTPOrBuilder `json:"httpOrBuilder"`

	IntervalSeconds float64 `json:"intervalSeconds"`

	TimeoutSeconds float64 `json:"timeoutSeconds"`
}

func (self *HealthCheckOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *HealthCheckOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthCheckOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthCheckOrBuilder cannot copy the values from %#v", other)
}

func (self *HealthCheckOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *HealthCheckOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HealthCheckOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HealthCheckOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *HealthCheckOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthCheckOrBuilder", name)

	case "command", "Command":
		v, ok := value.(*CommandInfo)
		if ok {
			self.Command = v
			self.present["command"] = true
			return nil
		} else {
			return fmt.Errorf("Field command/Command: value %v(%T) couldn't be cast to type *CommandInfo", value, value)
		}

	case "commandOrBuilder", "CommandOrBuilder":
		v, ok := value.(*CommandInfoOrBuilder)
		if ok {
			self.CommandOrBuilder = v
			self.present["commandOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field commandOrBuilder/CommandOrBuilder: value %v(%T) couldn't be cast to type *CommandInfoOrBuilder", value, value)
		}

	case "consecutiveFailures", "ConsecutiveFailures":
		v, ok := value.(int32)
		if ok {
			self.ConsecutiveFailures = v
			self.present["consecutiveFailures"] = true
			return nil
		} else {
			return fmt.Errorf("Field consecutiveFailures/ConsecutiveFailures: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "delaySeconds", "DelaySeconds":
		v, ok := value.(float64)
		if ok {
			self.DelaySeconds = v
			self.present["delaySeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field delaySeconds/DelaySeconds: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "gracePeriodSeconds", "GracePeriodSeconds":
		v, ok := value.(float64)
		if ok {
			self.GracePeriodSeconds = v
			self.present["gracePeriodSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field gracePeriodSeconds/GracePeriodSeconds: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "http", "Http":
		v, ok := value.(*HTTP)
		if ok {
			self.Http = v
			self.present["http"] = true
			return nil
		} else {
			return fmt.Errorf("Field http/Http: value %v(%T) couldn't be cast to type *HTTP", value, value)
		}

	case "httpOrBuilder", "HttpOrBuilder":
		v, ok := value.(*HTTPOrBuilder)
		if ok {
			self.HttpOrBuilder = v
			self.present["httpOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field httpOrBuilder/HttpOrBuilder: value %v(%T) couldn't be cast to type *HTTPOrBuilder", value, value)
		}

	case "intervalSeconds", "IntervalSeconds":
		v, ok := value.(float64)
		if ok {
			self.IntervalSeconds = v
			self.present["intervalSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field intervalSeconds/IntervalSeconds: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "timeoutSeconds", "TimeoutSeconds":
		v, ok := value.(float64)
		if ok {
			self.TimeoutSeconds = v
			self.present["timeoutSeconds"] = true
			return nil
		} else {
			return fmt.Errorf("Field timeoutSeconds/TimeoutSeconds: value %v(%T) couldn't be cast to type float64", value, value)
		}

	}
}

func (self *HealthCheckOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HealthCheckOrBuilder", name)

	case "command", "Command":
		if self.present != nil {
			if _, ok := self.present["command"]; ok {
				return self.Command, nil
			}
		}
		return nil, fmt.Errorf("Field Command no set on Command %+v", self)

	case "commandOrBuilder", "CommandOrBuilder":
		if self.present != nil {
			if _, ok := self.present["commandOrBuilder"]; ok {
				return self.CommandOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field CommandOrBuilder no set on CommandOrBuilder %+v", self)

	case "consecutiveFailures", "ConsecutiveFailures":
		if self.present != nil {
			if _, ok := self.present["consecutiveFailures"]; ok {
				return self.ConsecutiveFailures, nil
			}
		}
		return nil, fmt.Errorf("Field ConsecutiveFailures no set on ConsecutiveFailures %+v", self)

	case "delaySeconds", "DelaySeconds":
		if self.present != nil {
			if _, ok := self.present["delaySeconds"]; ok {
				return self.DelaySeconds, nil
			}
		}
		return nil, fmt.Errorf("Field DelaySeconds no set on DelaySeconds %+v", self)

	case "gracePeriodSeconds", "GracePeriodSeconds":
		if self.present != nil {
			if _, ok := self.present["gracePeriodSeconds"]; ok {
				return self.GracePeriodSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field GracePeriodSeconds no set on GracePeriodSeconds %+v", self)

	case "http", "Http":
		if self.present != nil {
			if _, ok := self.present["http"]; ok {
				return self.Http, nil
			}
		}
		return nil, fmt.Errorf("Field Http no set on Http %+v", self)

	case "httpOrBuilder", "HttpOrBuilder":
		if self.present != nil {
			if _, ok := self.present["httpOrBuilder"]; ok {
				return self.HttpOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field HttpOrBuilder no set on HttpOrBuilder %+v", self)

	case "intervalSeconds", "IntervalSeconds":
		if self.present != nil {
			if _, ok := self.present["intervalSeconds"]; ok {
				return self.IntervalSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field IntervalSeconds no set on IntervalSeconds %+v", self)

	case "timeoutSeconds", "TimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["timeoutSeconds"]; ok {
				return self.TimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field TimeoutSeconds no set on TimeoutSeconds %+v", self)

	}
}

func (self *HealthCheckOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthCheckOrBuilder", name)

	case "command", "Command":
		self.present["command"] = false

	case "commandOrBuilder", "CommandOrBuilder":
		self.present["commandOrBuilder"] = false

	case "consecutiveFailures", "ConsecutiveFailures":
		self.present["consecutiveFailures"] = false

	case "delaySeconds", "DelaySeconds":
		self.present["delaySeconds"] = false

	case "gracePeriodSeconds", "GracePeriodSeconds":
		self.present["gracePeriodSeconds"] = false

	case "http", "Http":
		self.present["http"] = false

	case "httpOrBuilder", "HttpOrBuilder":
		self.present["httpOrBuilder"] = false

	case "intervalSeconds", "IntervalSeconds":
		self.present["intervalSeconds"] = false

	case "timeoutSeconds", "TimeoutSeconds":
		self.present["timeoutSeconds"] = false

	}

	return nil
}

func (self *HealthCheckOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type HealthCheckOrBuilderList []*HealthCheckOrBuilder

func (self *HealthCheckOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthCheckOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthCheckOrBuilderList cannot copy the values from %#v", other)
}

func (list *HealthCheckOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *HealthCheckOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HealthCheckOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
