package dtos

import (
	"fmt"
	"io"
)

type SingularitySlave struct {
	present      map[string]bool
	Attributes   map[string]string                     `json:"attributes"`
	CurrentState *SingularityMachineStateHistoryUpdate `json:"currentState"`
	FirstSeenAt  int64                                 `json:"firstSeenAt"`
	Host         string                                `json:"host,omitempty"`
	Id           string                                `json:"id,omitempty"`
	RackId       string                                `json:"rackId,omitempty"`
}

func (self *SingularitySlave) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularitySlave) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularitySlave) FormatText() string {
	return FormatText(self)
}

func (self *SingularitySlave) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularitySlave) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularitySlave) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySlave", name)

	case "attributes", "Attributes":
		v, ok := value.(map[string]string)
		if ok {
			self.Attributes = v
			self.present["attributes"] = true
			return nil
		} else {
			return fmt.Errorf("Field attributes/Attributes: value %v(%T) couldn't be cast to type map[string]string", value, value)
		}

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

	case "host", "Host":
		v, ok := value.(string)
		if ok {
			self.Host = v
			self.present["host"] = true
			return nil
		} else {
			return fmt.Errorf("Field host/Host: value %v(%T) couldn't be cast to type string", value, value)
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

	case "rackId", "RackId":
		v, ok := value.(string)
		if ok {
			self.RackId = v
			self.present["rackId"] = true
			return nil
		} else {
			return fmt.Errorf("Field rackId/RackId: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularitySlave) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularitySlave", name)

	case "attributes", "Attributes":
		if self.present != nil {
			if _, ok := self.present["attributes"]; ok {
				return self.Attributes, nil
			}
		}
		return nil, fmt.Errorf("Field Attributes no set on Attributes %+v", self)

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

	case "host", "Host":
		if self.present != nil {
			if _, ok := self.present["host"]; ok {
				return self.Host, nil
			}
		}
		return nil, fmt.Errorf("Field Host no set on Host %+v", self)

	case "id", "Id":
		if self.present != nil {
			if _, ok := self.present["id"]; ok {
				return self.Id, nil
			}
		}
		return nil, fmt.Errorf("Field Id no set on Id %+v", self)

	case "rackId", "RackId":
		if self.present != nil {
			if _, ok := self.present["rackId"]; ok {
				return self.RackId, nil
			}
		}
		return nil, fmt.Errorf("Field RackId no set on RackId %+v", self)

	}
}

func (self *SingularitySlave) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularitySlave", name)

	case "attributes", "Attributes":
		self.present["attributes"] = false

	case "currentState", "CurrentState":
		self.present["currentState"] = false

	case "firstSeenAt", "FirstSeenAt":
		self.present["firstSeenAt"] = false

	case "host", "Host":
		self.present["host"] = false

	case "id", "Id":
		self.present["id"] = false

	case "rackId", "RackId":
		self.present["rackId"] = false

	}

	return nil
}

func (self *SingularitySlave) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularitySlaveList []*SingularitySlave

func (list *SingularitySlaveList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularitySlaveList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularitySlaveList) FormatJSON() string {
	return FormatJSON(list)
}
