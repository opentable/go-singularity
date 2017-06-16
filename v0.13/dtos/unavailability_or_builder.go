package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type UnavailabilityOrBuilder struct {
	present map[string]bool

	Duration *DurationInfo `json:"duration"`

	DurationOrBuilder *DurationInfoOrBuilder `json:"durationOrBuilder"`

	Start *TimeInfo `json:"start"`

	StartOrBuilder *TimeInfoOrBuilder `json:"startOrBuilder"`
}

func (self *UnavailabilityOrBuilder) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *UnavailabilityOrBuilder) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*UnavailabilityOrBuilder); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A UnavailabilityOrBuilder cannot absorb the values from %v", other)
}

func (self *UnavailabilityOrBuilder) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *UnavailabilityOrBuilder) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *UnavailabilityOrBuilder) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *UnavailabilityOrBuilder) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *UnavailabilityOrBuilder) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on UnavailabilityOrBuilder", name)

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

	}
}

func (self *UnavailabilityOrBuilder) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on UnavailabilityOrBuilder", name)

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

	}
}

func (self *UnavailabilityOrBuilder) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on UnavailabilityOrBuilder", name)

	case "duration", "Duration":
		self.present["duration"] = false

	case "durationOrBuilder", "DurationOrBuilder":
		self.present["durationOrBuilder"] = false

	case "start", "Start":
		self.present["start"] = false

	case "startOrBuilder", "StartOrBuilder":
		self.present["startOrBuilder"] = false

	}

	return nil
}

func (self *UnavailabilityOrBuilder) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type UnavailabilityOrBuilderList []*UnavailabilityOrBuilder

func (self *UnavailabilityOrBuilderList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*UnavailabilityOrBuilderList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A UnavailabilityOrBuilder cannot absorb the values from %v", other)
}

func (list *UnavailabilityOrBuilderList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *UnavailabilityOrBuilderList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *UnavailabilityOrBuilderList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
