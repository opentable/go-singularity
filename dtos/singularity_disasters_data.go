package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityDisastersData struct {
	present map[string]bool

	AutomatedActionsDisabled bool `json:"automatedActionsDisabled"`

	Disasters SingularityDisasterList `json:"disasters"`

	Stats SingularityDisasterDataPointList `json:"stats"`
}

func (self *SingularityDisastersData) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityDisastersData) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisastersData); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisastersData cannot copy the values from %#v", other)
}

func (self *SingularityDisastersData) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityDisastersData) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityDisastersData) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityDisastersData) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityDisastersData) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisastersData", name)

	case "automatedActionsDisabled", "AutomatedActionsDisabled":
		v, ok := value.(bool)
		if ok {
			self.AutomatedActionsDisabled = v
			self.present["automatedActionsDisabled"] = true
			return nil
		} else {
			return fmt.Errorf("Field automatedActionsDisabled/AutomatedActionsDisabled: value %v(%T) couldn't be cast to type bool", value, value)
		}

	case "disasters", "Disasters":
		v, ok := value.(SingularityDisasterList)
		if ok {
			self.Disasters = v
			self.present["disasters"] = true
			return nil
		} else {
			return fmt.Errorf("Field disasters/Disasters: value %v(%T) couldn't be cast to type SingularityDisasterList", value, value)
		}

	case "stats", "Stats":
		v, ok := value.(SingularityDisasterDataPointList)
		if ok {
			self.Stats = v
			self.present["stats"] = true
			return nil
		} else {
			return fmt.Errorf("Field stats/Stats: value %v(%T) couldn't be cast to type SingularityDisasterDataPointList", value, value)
		}

	}
}

func (self *SingularityDisastersData) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityDisastersData", name)

	case "automatedActionsDisabled", "AutomatedActionsDisabled":
		if self.present != nil {
			if _, ok := self.present["automatedActionsDisabled"]; ok {
				return self.AutomatedActionsDisabled, nil
			}
		}
		return nil, fmt.Errorf("Field AutomatedActionsDisabled no set on AutomatedActionsDisabled %+v", self)

	case "disasters", "Disasters":
		if self.present != nil {
			if _, ok := self.present["disasters"]; ok {
				return self.Disasters, nil
			}
		}
		return nil, fmt.Errorf("Field Disasters no set on Disasters %+v", self)

	case "stats", "Stats":
		if self.present != nil {
			if _, ok := self.present["stats"]; ok {
				return self.Stats, nil
			}
		}
		return nil, fmt.Errorf("Field Stats no set on Stats %+v", self)

	}
}

func (self *SingularityDisastersData) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityDisastersData", name)

	case "automatedActionsDisabled", "AutomatedActionsDisabled":
		self.present["automatedActionsDisabled"] = false

	case "disasters", "Disasters":
		self.present["disasters"] = false

	case "stats", "Stats":
		self.present["stats"] = false

	}

	return nil
}

func (self *SingularityDisastersData) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityDisastersDataList []*SingularityDisastersData

func (self *SingularityDisastersDataList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityDisastersDataList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityDisastersDataList cannot copy the values from %#v", other)
}

func (list *SingularityDisastersDataList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityDisastersDataList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityDisastersDataList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
