package dtos

import (
	"fmt"
	"io"
)

type SingularityMachineChangeRequest struct {
	present map[string]bool
	Message string `json:"message,omitempty"`
}

func (self *SingularityMachineChangeRequest) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, self)
}

func (self *SingularityMachineChangeRequest) MarshalJSON() ([]byte, error) {
	return MarshalJSON(self)
}

func (self *SingularityMachineChangeRequest) FormatText() string {
	return FormatText(self)
}

func (self *SingularityMachineChangeRequest) FormatJSON() string {
	return FormatJSON(self)
}

func (self *SingularityMachineChangeRequest) FieldsPresent() []string {
	return presenceFromMap(self.present)
}

func (self *SingularityMachineChangeRequest) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMachineChangeRequest", name)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	}
}

func (self *SingularityMachineChangeRequest) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityMachineChangeRequest", name)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	}
}

func (self *SingularityMachineChangeRequest) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMachineChangeRequest", name)

	case "message", "Message":
		self.present["message"] = false

	}

	return nil
}

func (self *SingularityMachineChangeRequest) LoadMap(from map[string]interface{}) error {
	return loadMapIntoDTO(from, self)
}

type SingularityMachineChangeRequestList []*SingularityMachineChangeRequest

func (list *SingularityMachineChangeRequestList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *SingularityMachineChangeRequestList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityMachineChangeRequestList) FormatJSON() string {
	return FormatJSON(list)
}
