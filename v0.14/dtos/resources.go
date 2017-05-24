package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Resources struct {
	present map[string]bool

	Cpus float64 `json:"cpus"`

	MemoryMb float64 `json:"memoryMb"`

	NumPorts int32 `json:"numPorts"`
}

func (self *Resources) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *Resources) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*Resources); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A Resources cannot copy the values from %#v", other)
}

func (self *Resources) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *Resources) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Resources) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Resources) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *Resources) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Resources", name)

	case "cpus", "Cpus":
		v, ok := value.(float64)
		if ok {
			self.Cpus = v
			self.present["cpus"] = true
			return nil
		} else {
			return fmt.Errorf("Field cpus/Cpus: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "memoryMb", "MemoryMb":
		v, ok := value.(float64)
		if ok {
			self.MemoryMb = v
			self.present["memoryMb"] = true
			return nil
		} else {
			return fmt.Errorf("Field memoryMb/MemoryMb: value %v(%T) couldn't be cast to type float64", value, value)
		}

	case "numPorts", "NumPorts":
		v, ok := value.(int32)
		if ok {
			self.NumPorts = v
			self.present["numPorts"] = true
			return nil
		} else {
			return fmt.Errorf("Field numPorts/NumPorts: value %v(%T) couldn't be cast to type int32", value, value)
		}

	}
}

func (self *Resources) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Resources", name)

	case "cpus", "Cpus":
		if self.present != nil {
			if _, ok := self.present["cpus"]; ok {
				return self.Cpus, nil
			}
		}
		return nil, fmt.Errorf("Field Cpus no set on Cpus %+v", self)

	case "memoryMb", "MemoryMb":
		if self.present != nil {
			if _, ok := self.present["memoryMb"]; ok {
				return self.MemoryMb, nil
			}
		}
		return nil, fmt.Errorf("Field MemoryMb no set on MemoryMb %+v", self)

	case "numPorts", "NumPorts":
		if self.present != nil {
			if _, ok := self.present["numPorts"]; ok {
				return self.NumPorts, nil
			}
		}
		return nil, fmt.Errorf("Field NumPorts no set on NumPorts %+v", self)

	}
}

func (self *Resources) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on Resources", name)

	case "cpus", "Cpus":
		self.present["cpus"] = false

	case "memoryMb", "MemoryMb":
		self.present["memoryMb"] = false

	case "numPorts", "NumPorts":
		self.present["numPorts"] = false

	}

	return nil
}

func (self *Resources) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type ResourcesList []*Resources

func (self *ResourcesList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*ResourcesList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A ResourcesList cannot copy the values from %#v", other)
}

func (list *ResourcesList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *ResourcesList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *ResourcesList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
