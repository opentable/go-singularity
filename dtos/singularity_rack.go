package dtos

import (
	"fmt"
	"io"
)

type SingularityRack struct {
	present      map[string]bool
	CurrentState *SingularityMachineStateHistoryUpdate `json:"currentState"`
	FirstSeenAt  int64                                 `json:"firstSeenAt"`
	Id           string                                `json:"id,omitempty"`
}

func (self *SingularityRack) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityRack) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityRack) FormatText() string {
	return FormatText(self)
}

func (self *SingularityRack) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityRack) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityRack) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRack", name)

	case "currentState", "CurrentState":
		v, ok := value.(*SingularityMachineStateHistoryUpdate)
		if ok {
			self.CurrentState = v
			self.present["currentState"] = true
			return nil
		} else {
			return fmt.Errorf("Field currentState/CurrentState: value %v(%T) couldn't be cast to type *SingularityMachineStateHistoryUpdate", value, value)
		}

	case "firstSeenAt", "FirstSeenAt":
		v, ok := value.(int64)
		if ok {
			self.FirstSeenAt = v
			self.present["firstSeenAt"] = true
			return nil
		} else {
			return fmt.Errorf("Field firstSeenAt/FirstSeenAt: value %v(%T) couldn't be cast to type int64", value, value)
		}

	case "id", "Id":
		v, ok := value.(string)
		if ok {
			self.Id = v
			self.present["id"] = true
			return nil
		} else {
			return fmt.Errorf("Field id/Id: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityRack) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityRack", name)

	case "currentState", "CurrentState":
		if self.present != nil {
			if _, ok := self.present["currentState"]; ok {
				return self.CurrentState, nil
			}
		}
		return nil, fmt.Errorf("Field CurrentState no set on CurrentState %+v", self)

	case "firstSeenAt", "FirstSeenAt":
		if self.present != nil {
			if _, ok := self.present["firstSeenAt"]; ok {
				return self.FirstSeenAt, nil
			}
		}
		return nil, fmt.Errorf("Field FirstSeenAt no set on FirstSeenAt %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	}
}

func (self *SingularityRack) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityRack", name)

	case "currentState", "CurrentState":
		self.present["currentState"] = false

	case "firstSeenAt", "FirstSeenAt":
		self.present["firstSeenAt"] = false

	case "id", "Id":
		self.present["id"] = false

	}

	return nil
}

func (self *SingularityRack) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityRackList []*SingularityRack

func (list *SingularityRackList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityRackList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityRackList) FormatJSON() string {
	return FormatJSON(list)
}
