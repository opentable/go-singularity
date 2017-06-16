package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Unavailability struct {
	present map[string]bool

	// AllFields *Map[FieldDescriptor,Object] `json:"allFields"`

	DefaultInstanceForType *Unavailability `json:"defaultInstanceForType"`

	DescriptorForType *Descriptor `json:"descriptorForType"`

	Duration *DurationInfo `json:"duration"`

	DurationOrBuilder *DurationInfoOrBuilder `json:"durationOrBuilder"`

	InitializationErrorString string `json:"initializationErrorString,omitempty"`

	Initialized bool `json:"initialized"`

	// ParserForType *com.google.protobuf.Parser<org.apache.mesos.Protos$Unavailability> `json:"parserForType"`

	SerializedSize int32 `json:"serializedSize"`

	Start *TimeInfo `json:"start"`

	StartOrBuilder *TimeInfoOrBuilder `json:"startOrBuilder"`

	UnknownFields *UnknownFieldSet `json:"unknownFields"`
}

func (self *Unavailability) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Unavailability) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Unavailability); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Unavailability cannot absorb the values from %v", other)
}

func (self *Unavailability) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Unavailability) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Unavailability) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Unavailability) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Unavailability) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Unavailability", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		v, ok := value.(*Unavailability)
		if ok {
			self.DefaultInstanceForType = v
			self.present["defaultInstanceForType"] = true
			return nil
		} else {
			return fmt.Errorf("Field defaultInstanceForType/DefaultInstanceForType: value %v(%T) couldn't be cast to type *Unavailability", value, value)
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

	case "duration", "Duration":
		v, ok := value.(*DurationInfo)
		if ok {
			self.Duration = v
			self.present["duration"] = true
			return nil
		} else {
			return fmt.Errorf("Field duration/Duration: value %v(%T) couldn't be cast to type *DurationInfo", value, value)
		}

	case "durationOrBuilder", "DurationOrBuilder":
		v, ok := value.(*DurationInfoOrBuilder)
		if ok {
			self.DurationOrBuilder = v
			self.present["durationOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field durationOrBuilder/DurationOrBuilder: value %v(%T) couldn't be cast to type *DurationInfoOrBuilder", value, value)
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

	case "serializedSize", "SerializedSize":
		v, ok := value.(int32)
		if ok {
			self.SerializedSize = v
			self.present["serializedSize"] = true
			return nil
		} else {
			return fmt.Errorf("Field serializedSize/SerializedSize: value %v(%T) couldn't be cast to type int32", value, value)
		}

	case "start", "Start":
		v, ok := value.(*TimeInfo)
		if ok {
			self.Start = v
			self.present["start"] = true
			return nil
		} else {
			return fmt.Errorf("Field start/Start: value %v(%T) couldn't be cast to type *TimeInfo", value, value)
		}

	case "startOrBuilder", "StartOrBuilder":
		v, ok := value.(*TimeInfoOrBuilder)
		if ok {
			self.StartOrBuilder = v
			self.present["startOrBuilder"] = true
			return nil
		} else {
			return fmt.Errorf("Field startOrBuilder/StartOrBuilder: value %v(%T) couldn't be cast to type *TimeInfoOrBuilder", value, value)
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

func (self *Unavailability) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Unavailability", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		if self.present != nil {
			if _, ok := self.present["defaultInstanceForType"]; ok {
				return self.DefaultInstanceForType, nil
			}
		}
		return nil, fmt.Errorf("Field DefaultInstanceForType no set on DefaultInstanceForType %+v", self)

	case "descriptorForType", "DescriptorForType":
		if self.present != nil {
			if _, ok := self.present["descriptorForType"]; ok {
				return self.DescriptorForType, nil
			}
		}
		return nil, fmt.Errorf("Field DescriptorForType no set on DescriptorForType %+v", self)

	case "duration", "Duration":
		if self.present != nil {
			if _, ok := self.present["duration"]; ok {
				return self.Duration, nil
			}
		}
		return nil, fmt.Errorf("Field Duration no set on Duration %+v", self)

	case "durationOrBuilder", "DurationOrBuilder":
		if self.present != nil {
			if _, ok := self.present["durationOrBuilder"]; ok {
				return self.DurationOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field DurationOrBuilder no set on DurationOrBuilder %+v", self)

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

	case "serializedSize", "SerializedSize":
		if self.present != nil {
			if _, ok := self.present["serializedSize"]; ok {
				return self.SerializedSize, nil
			}
		}
		return nil, fmt.Errorf("Field SerializedSize no set on SerializedSize %+v", self)

	case "start", "Start":
		if self.present != nil {
			if _, ok := self.present["start"]; ok {
				return self.Start, nil
			}
		}
		return nil, fmt.Errorf("Field Start no set on Start %+v", self)

	case "startOrBuilder", "StartOrBuilder":
		if self.present != nil {
			if _, ok := self.present["startOrBuilder"]; ok {
				return self.StartOrBuilder, nil
			}
		}
		return nil, fmt.Errorf("Field StartOrBuilder no set on StartOrBuilder %+v", self)

	case "unknownFields", "UnknownFields":
		if self.present != nil {
			if _, ok := self.present["unknownFields"]; ok {
				return self.UnknownFields, nil
			}
		}
		return nil, fmt.Errorf("Field UnknownFields no set on UnknownFields %+v", self)

	}
}

func (self *Unavailability) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Unavailability", name)

	case "defaultInstanceForType", "DefaultInstanceForType":
		self.present["defaultInstanceForType"] = false

	case "descriptorForType", "DescriptorForType":
		self.present["descriptorForType"] = false

	case "duration", "Duration":
		self.present["duration"] = false

	case "durationOrBuilder", "DurationOrBuilder":
		self.present["durationOrBuilder"] = false

	case "initializationErrorString", "InitializationErrorString":
		self.present["initializationErrorString"] = false

	case "initialized", "Initialized":
		self.present["initialized"] = false

	case "serializedSize", "SerializedSize":
		self.present["serializedSize"] = false

	case "start", "Start":
		self.present["start"] = false

	case "startOrBuilder", "StartOrBuilder":
		self.present["startOrBuilder"] = false

	case "unknownFields", "UnknownFields":
		self.present["unknownFields"] = false

	}

	return nil
}

func (self *Unavailability) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type UnavailabilityList []*Unavailability

func (self *UnavailabilityList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*UnavailabilityList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Unavailability cannot absorb the values from %v", other)
}

func (list *UnavailabilityList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *UnavailabilityList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *UnavailabilityList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
