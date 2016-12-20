package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type HealthCheck struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	Command *CommandInfo `json:"command"`

	CommandOrBuilder *CommandInfoOrBuilder `json:"commandOrBuilder"`

	ConsecutiveFailures int32 `json:"consecutiveFailures"`

	DefaultInstanceForType *HealthCheck `json:"defaultInstanceForType"`

	DelaySeconds float64 `json:"delaySeconds"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	GracePeriodSeconds float64 `json:"gracePeriodSeconds"`

	Http *HTTP `json:"http"`

	HttpOrBuilder *HTTPOrBuilder `json:"httpOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	IntervalSeconds float64 `json:"intervalSeconds"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$HealthCheck> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	TimeoutSeconds float64 `json:"timeoutSeconds"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *HealthCheck) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *HealthCheck) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthCheck); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthCheck cannot copy the values from %#v", other)
}

func (self *HealthCheck) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *HealthCheck) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *HealthCheck) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *HealthCheck) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *HealthCheck) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthCheck", name)

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

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*HealthCheck)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *HealthCheck", value, value)
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

	case "descriptorForType", "DescriptorForType":
		v, ok := value.(*Descriptor)
		if ok {
			self.DescriptorForType = v
			self.present["descriptorForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field descriptorForType/DescriptorForType: value %v(%T) couldn't be cast to type *Descriptor", value, value)
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

	case "initializationErrorString", "InitializationErrorString":
		v, ok := value.(string)
		if ok {
			self.InitializationErrorString = v
			self.present["initializationErrorString"] = true
			return nil
		} else {
			return fmt.Errorf("Field initializationErrorString/InitializationErrorString: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "initialized", "Initialized":
		v, ok := value.(bool)
		if ok {
			self.Initialized = v
			self.present["initialized"] = true
			return nil
		} else {
			return fmt.Errorf("Field initialized/Initialized: value %v(%T) couldn't be cast to type bool", value, value)
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

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
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

	case "unknownFields", "UnknownFields":
		v, ok := value.(*UnknownFieldSet)
		if ok {
			self.UnknownFields = v
			self.present["unknownFields"] = true
			return nil
		} else {
			return fmt.Errorf("Field unknownFields/UnknownFields: value %v(%T) couldn't be cast to type *UnknownFieldSet", value, value)
		}

	}
}

func (self *HealthCheck) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on HealthCheck", name)

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

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "delaySeconds", "DelaySeconds":
		if self.present != nil {
			if _, ok := self.present["delaySeconds"]; ok {
				return self.DelaySeconds, nil
			}
		}
		return nil, fmt.Errorf("Field DelaySeconds no set on DelaySeconds %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

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

	case "initializationErrorString", "InitializationErrorString":
		if self.present != nil {
			if _, ok := self.present["initializationErrorString"]; ok {
				return self.InitializationErrorString, nil
			}
		}
		return nil, fmt.Errorf("Field InitializationErrorString no set on InitializationErrorString %+v", self)

	case "initialized", "Initialized":
		if self.present != nil {
			if _, ok := self.present["initialized"]; ok {
				return self.Initialized, nil
			}
		}
		return nil, fmt.Errorf("Field Initialized no set on Initialized %+v", self)

	case "intervalSeconds", "IntervalSeconds":
		if self.present != nil {
			if _, ok := self.present["intervalSeconds"]; ok {
				return self.IntervalSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field IntervalSeconds no set on IntervalSeconds %+v", self)

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "timeoutSeconds", "TimeoutSeconds":
		if self.present != nil {
			if _, ok := self.present["timeoutSeconds"]; ok {
				return self.TimeoutSeconds, nil
			}
		}
		return nil, fmt.Errorf("Field TimeoutSeconds no set on TimeoutSeconds %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *HealthCheck) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on HealthCheck", name)

	case "command", "Command":
		self.present["command"] = false

	case "commandOrBuilder", "CommandOrBuilder":
		self.present["commandOrBuilder"] = false

	case "consecutiveFailures", "ConsecutiveFailures":
		self.present["consecutiveFailures"] = false

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "delaySeconds", "DelaySeconds":
		self.present["delaySeconds"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "gracePeriodSeconds", "GracePeriodSeconds":
		self.present["gracePeriodSeconds"] = false

	case "http", "Http":
		self.present["http"] = false

	case "httpOrBuilder", "HttpOrBuilder":
		self.present["httpOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "intervalSeconds", "IntervalSeconds":
		self.present["intervalSeconds"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "timeoutSeconds", "TimeoutSeconds":
		self.present["timeoutSeconds"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *HealthCheck) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type HealthCheckList []*HealthCheck

func (self *HealthCheckList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*HealthCheckList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A HealthCheckList cannot copy the values from %#v", other)
}

func (list *HealthCheckList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *HealthCheckList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *HealthCheckList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
