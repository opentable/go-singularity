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

	Message string `json:"message,omitempty"`

	OutputFilename string `json:"outputFilename,omitempty"`

	ShellRequestId *SingularityTaskShellCommandRequestId `json:"shellRequestId"`

	Timestamp int64 `json:"timestamp"`

	UpdateType SingularityTaskShellCommandUpdateUpdateType `json:"updateType"`
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

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = v
			self.present["message"] = true
			return nil
		} else {
			return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)
		}

	case "outputFilename", "OutputFilename":
		v, ok := value.(string)
		if ok {
			self.OutputFilename = v
			self.present["outputFilename"] = true
			return nil
		} else {
			return fmt.Errorf("Field outputFilename/OutputFilename: value %v(%T) couldn't be cast to type string", value, value)
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

	case "updateType", "UpdateType":
		v, ok := value.(SingularityTaskShellCommandUpdateUpdateType)
		if ok {
			self.UpdateType = v
			self.present["updateType"] = true
			return nil
		} else {
			return fmt.Errorf("Field updateType/UpdateType: value %v(%T) couldn't be cast to type SingularityTaskShellCommandUpdateUpdateType", value, value)
		}

	}
}

func (self *SingularityTaskShellCommandUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "message", "Message":
		if self.present != nil {
			if _, ok := self.present["message"]; ok {
				return self.Message, nil
			}
		}
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "outputFilename", "OutputFilename":
		if self.present != nil {
			if _, ok := self.present["outputFilename"]; ok {
				return self.OutputFilename, nil
			}
		}
		return nil, fmt.Errorf("Field OutputFilename no set on OutputFilename %+v", self)

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

	case "updateType", "UpdateType":
		if self.present != nil {
			if _, ok := self.present["updateType"]; ok {
				return self.UpdateType, nil
			}
		}
		return nil, fmt.Errorf("Field UpdateType no set on UpdateType %+v", self)

	}
}

func (self *SingularityTaskShellCommandUpdate) ClearField(name string) error {
	if self.present == nil {
		self.present = make(map[string]bool)
	}
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "message", "Message":
		self.present["message"] = false

	case "outputFilename", "OutputFilename":
		self.present["outputFilename"] = false

	case "shellRequestId", "ShellRequestId":
		self.present["shellRequestId"] = false

	case "timestamp", "Timestamp":
		self.present["timestamp"] = false

	case "updateType", "UpdateType":
		self.present["updateType"] = false

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
