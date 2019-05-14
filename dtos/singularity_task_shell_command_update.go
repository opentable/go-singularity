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
	ShellRequestId *SingularityTaskShellCommandRequestId        `json:"shellRequestId,omitempty"`
	Timestamp      *int64                                       `json:"timestamp,omitempty"`
	Message        *string                                      `json:"message,omitempty"`
	OutputFilename *string                                      `json:"outputFilename,omitempty"`
	UpdateType     *SingularityTaskShellCommandUpdateUpdateType `json:"updateType,omitempty"`
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

func (self *SingularityTaskShellCommandUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityTaskShellCommandUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityTaskShellCommandUpdate) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "shellRequestId", "ShellRequestId":
		v, ok := value.(*SingularityTaskShellCommandRequestId)
		if ok {
			self.ShellRequestId = v
			return nil
		}
		return fmt.Errorf("Field shellRequestId/ShellRequestId: value %v(%T) couldn't be cast to type *SingularityTaskShellCommandRequestId", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "outputFilename", "OutputFilename":
		v, ok := value.(string)
		if ok {
			self.OutputFilename = &v
			return nil
		}
		return fmt.Errorf("Field outputFilename/OutputFilename: value %v(%T) couldn't be cast to type string", value, value)

	case "updateType", "UpdateType":
		v, ok := value.(SingularityTaskShellCommandUpdateUpdateType)
		if ok {
			self.UpdateType = &v
			return nil
		}
		return fmt.Errorf("Field updateType/UpdateType: value %v(%T) couldn't be cast to type SingularityTaskShellCommandUpdateUpdateType", value, value)

	}
}

func (self *SingularityTaskShellCommandUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "shellRequestId", "ShellRequestId":
		return self.ShellRequestId, nil
		return nil, fmt.Errorf("Field ShellRequestId no set on ShellRequestId %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "outputFilename", "OutputFilename":
		return *self.OutputFilename, nil
		return nil, fmt.Errorf("Field OutputFilename no set on OutputFilename %+v", self)

	case "updateType", "UpdateType":
		return *self.UpdateType, nil
		return nil, fmt.Errorf("Field UpdateType no set on UpdateType %+v", self)

	}
}

func (self *SingularityTaskShellCommandUpdate) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityTaskShellCommandUpdate", name)

	case "shellRequestId", "ShellRequestId":
		self.ShellRequestId = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	case "message", "Message":
		self.Message = nil

	case "outputFilename", "OutputFilename":
		self.OutputFilename = nil

	case "updateType", "UpdateType":
		self.UpdateType = nil

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
