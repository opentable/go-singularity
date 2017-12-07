package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityTaskShellCommandUpdateUpdateType string

const (
	SingularityTaskShellCommandUpdateUpdateTypeINVALID  SingularityTaskShellCommandUpdateUpdateType = "INVALID"
	SingularityTaskShellCommandUpdateUpdateTypeACKED    SingularityTaskShellCommandUpdateUpdateType = "ACKED"
	SingularityTaskShellCommandUpdateUpdateTypeSTARTED  SingularityTaskShellCommandUpdateUpdateType = "STARTED"
	SingularityTaskShellCommandUpdateUpdateTypeFINISHED SingularityTaskShellCommandUpdateUpdateType = "FINISHED"
	SingularityTaskShellCommandUpdateUpdateTypeFAILED   SingularityTaskShellCommandUpdateUpdateType = "FAILED"
)

type SingularityTaskShellCommandUpdate struct {
	present map[string]bool

	UpdateType SingularityTaskShellCommandUpdateUpdateType `json:"updateType"`

	UpdateType SingularityTaskShellCommandUpdateUpdateType `json:"updateType"`

	ShellRequestId *SingularityTaskShellCommandRequestId `json:"shellRequestId"`

	Timestamp int64 `json:"timestamp"`
}

func (self *SingularityTaskShellCommandUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityTaskShellCommandUpdate) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandUpdate); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandUpdate cannot copy the values from %#v", other)
}

func (self *SingularityTaskShellCommandUpdate) MarshalJSON() ([]byte, error) {
	return swaggering.MarshalJSON(self)
}

func (self *SingularityTaskShellCommandUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandUpdate) FieldsPresent() []string {
	return swaggering.PresenceFromMap(self.present)
}

func (self *SingularityTaskShellCommandUpdate) SetField(name string, value interface{}) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "updateType", "UpdateType":
		v, ok := value.(SingularityTaskShellCommandUpdateUpdateType)
		if ok {
			self.UpdateType = v
			self.present["updateType"] = true
			return nil
		} else {
			return fmt.Errorf("Field updateType/UpdateType: value %v(%T) couldn't be cast to type SingularityTaskShellCommandUpdateUpdateType", value, value)
		}

	case "updateType", "UpdateType":
		v, ok := value.(SingularityTaskShellCommandUpdateUpdateType)
		if ok {
			self.UpdateType = v
			self.present["updateType"] = true
			return nil
		} else {
			return fmt.Errorf("Field updateType/UpdateType: value %v(%T) couldn't be cast to type SingularityTaskShellCommandUpdateUpdateType", value, value)
		}

	case "shellRequestId", "ShellRequestId":
		v, ok := value.(*SingularityTaskShellCommandRequestId)
		if ok {
			self.ShellRequestId = v
			self.present["shellRequestId"] = true
			return nil
		} else {
			return fmt.Errorf("Field shellRequestId/ShellRequestId: value %v(%T) couldn't be cast to type *SingularityTaskShellCommandRequestId", value, value)
		}

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = v
			self.present["timestamp"] = true
			return nil
		} else {
			return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)
		}

	}
}

func (self *SingularityTaskShellCommandUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "updateType", "UpdateType":
		if self.present != nil {
			if _, ok := self.present["updateType"]; ok {
				return self.UpdateType, nil
			}
		}
		return nil, fmt.Errorf("Field UpdateType no set on UpdateType %+v", self)

	case "updateType", "UpdateType":
		if self.present != nil {
			if _, ok := self.present["updateType"]; ok {
				return self.UpdateType, nil
			}
		}
		return nil, fmt.Errorf("Field UpdateType no set on UpdateType %+v", self)

	case "shellRequestId", "ShellRequestId":
		if self.present != nil {
			if _, ok := self.present["shellRequestId"]; ok {
				return self.ShellRequestId, nil
			}
		}
		return nil, fmt.Errorf("Field ShellRequestId no set on ShellRequestId %+v", self)

	case "timestamp", "Timestamp":
		if self.present != nil {
			if _, ok := self.present["timestamp"]; ok {
				return self.Timestamp, nil
			}
		}
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityTaskShellCommandUpdate) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "updateType", "UpdateType":
		self.present["updateType"] = false

	case "updateType", "UpdateType":
		self.present["updateType"] = false

	case "shellRequestId", "ShellRequestId":
		self.present["shellRequestId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	}

	return nil
}

func (self *SingularityTaskShellCommandUpdate) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityTaskShellCommandUpdateList []*SingularityTaskShellCommandUpdate

func (self *SingularityTaskShellCommandUpdateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityTaskShellCommandUpdateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityTaskShellCommandUpdateList cannot copy the values from %#v", other)
}

func (list *SingularityTaskShellCommandUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityTaskShellCommandUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityTaskShellCommandUpdateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
