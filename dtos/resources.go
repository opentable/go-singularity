package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type Resources struct {
	MemoryMb *float64 `json:"memoryMb,omitempty"`
	NumPorts *int32   `json:"numPorts,omitempty"`
	DiskMb   *float64 `json:"diskMb,omitempty"`
	Cpus     *float64 `json:"cpus,omitempty"`
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

func (self *Resources) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *Resources) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *Resources) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on Resources", name)

	case "memoryMb", "MemoryMb":
		v, ok := value.(float64)
		if ok {
			self.MemoryMb = &v
			return nil
		}
		return fmt.Errorf("Field memoryMb/MemoryMb: value %v(%T) couldn't be cast to type float64", value, value)

	case "numPorts", "NumPorts":
		v, ok := value.(int32)
		if ok {
			self.NumPorts = &v
			return nil
		}
		return fmt.Errorf("Field numPorts/NumPorts: value %v(%T) couldn't be cast to type int32", value, value)

	case "diskMb", "DiskMb":
		v, ok := value.(float64)
		if ok {
			self.DiskMb = &v
			return nil
		}
		return fmt.Errorf("Field diskMb/DiskMb: value %v(%T) couldn't be cast to type float64", value, value)

	case "cpus", "Cpus":
		v, ok := value.(float64)
		if ok {
			self.Cpus = &v
			return nil
		}
		return fmt.Errorf("Field cpus/Cpus: value %v(%T) couldn't be cast to type float64", value, value)

	}
}

func (self *Resources) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on Resources", name)

	case "memoryMb", "MemoryMb":
		return *self.MemoryMb, nil
		return nil, fmt.Errorf("Field MemoryMb no set on MemoryMb %+v", self)

	case "numPorts", "NumPorts":
		return *self.NumPorts, nil
		return nil, fmt.Errorf("Field NumPorts no set on NumPorts %+v", self)

	case "diskMb", "DiskMb":
		return *self.DiskMb, nil
		return nil, fmt.Errorf("Field DiskMb no set on DiskMb %+v", self)

	case "cpus", "Cpus":
		return *self.Cpus, nil
		return nil, fmt.Errorf("Field Cpus no set on Cpus %+v", self)

	}
}

func (self *Resources) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on Resources", name)

	case "memoryMb", "MemoryMb":
		self.MemoryMb = nil

	case "numPorts", "NumPorts":
		self.NumPorts = nil

	case "diskMb", "DiskMb":
		self.DiskMb = nil

	case "cpus", "Cpus":
		self.Cpus = nil

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
