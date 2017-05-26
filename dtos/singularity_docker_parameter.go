package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDockerParameter struct {
	present map[string]bool

	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

func (self *SingularityDockerParameter) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDockerParameter) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerParameter); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerParameter cannot copy the values from %#v", other)
}

func (self *SingularityDockerParameter) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDockerParameter) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDockerParameter) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDockerParameter) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDockerParameter) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerParameter", name)

	case "key", "Key":
		v, ok := value.(string)
		if ok {
			self.Key = v
			self.present["key"] = true
			return nil
		} else {
			return fmt.Errorf("Field key/Key: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "value", "Value":
		v, ok := value.(string)
		if ok {
			self.Value = v
			self.present["value"] = true
			return nil
		} else {
			return fmt.Errorf("Field value/Value: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityDockerParameter) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDockerParameter", name)

	case "key", "Key":
		if self.present != nil {
			if _, ok := self.present["key"]; ok {
				return self.Key, nil
			}
		}
		return nil, fmt.Errorf("Field Key no set on Key %+v", self)

	case "value", "Value":
		if self.present != nil {
			if _, ok := self.present["value"]; ok {
				return self.Value, nil
			}
		}
		return nil, fmt.Errorf("Field Value no set on Value %+v", self)

	}
}

func (self *SingularityDockerParameter) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDockerParameter", name)

	case "key", "Key":
		self.present["key"] = false

	case "value", "Value":
		self.present["value"] = false

	}

	return nil
}

func (self *SingularityDockerParameter) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDockerParameterList []*SingularityDockerParameter

func (self *SingularityDockerParameterList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDockerParameterList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDockerParameterList cannot copy the values from %#v", other)
}

func (list *SingularityDockerParameterList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDockerParameterList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDockerParameterList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
