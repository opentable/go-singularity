package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityMachineStateHistoryUpdateMachineState string

const (
	SingularityMachineStateHistoryUpdateMachineStateMISSING_ON_STARTUP    SingularityMachineStateHistoryUpdateMachineState = "MISSING_ON_STARTUP"
	SingularityMachineStateHistoryUpdateMachineStateACTIVE                SingularityMachineStateHistoryUpdateMachineState = "ACTIVE"
	SingularityMachineStateHistoryUpdateMachineStateSTARTING_DECOMMISSION SingularityMachineStateHistoryUpdateMachineState = "STARTING_DECOMMISSION"
	SingularityMachineStateHistoryUpdateMachineStateDECOMMISSIONING       SingularityMachineStateHistoryUpdateMachineState = "DECOMMISSIONING"
	SingularityMachineStateHistoryUpdateMachineStateDECOMMISSIONED        SingularityMachineStateHistoryUpdateMachineState = "DECOMMISSIONED"
	SingularityMachineStateHistoryUpdateMachineStateDEAD                  SingularityMachineStateHistoryUpdateMachineState = "DEAD"
	SingularityMachineStateHistoryUpdateMachineStateFROZEN                SingularityMachineStateHistoryUpdateMachineState = "FROZEN"
)

type SingularityMachineStateHistoryUpdate struct {
	ObjectId  *string                                           `json:"objectId,omitempty"`
	State     *SingularityMachineStateHistoryUpdateMachineState `json:"state,omitempty"`
	User      *string                                           `json:"user,omitempty"`
	Message   *string                                           `json:"message,omitempty"`
	Timestamp *int64                                            `json:"timestamp,omitempty"`
}

func (self *SingularityMachineStateHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityMachineStateHistoryUpdate) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMachineStateHistoryUpdate); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMachineStateHistoryUpdate cannot copy the values from %#v", other)
}

func (self *SingularityMachineStateHistoryUpdate) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityMachineStateHistoryUpdate) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityMachineStateHistoryUpdate) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMachineStateHistoryUpdate", name)

	case "objectId", "ObjectId":
		v, ok := value.(string)
		if ok {
			self.ObjectId = &v
			return nil
		}
		return fmt.Errorf("Field objectId/ObjectId: value %v(%T) couldn't be cast to type string", value, value)

	case "state", "State":
		v, ok := value.(SingularityMachineStateHistoryUpdateMachineState)
		if ok {
			self.State = &v
			return nil
		}
		return fmt.Errorf("Field state/State: value %v(%T) couldn't be cast to type SingularityMachineStateHistoryUpdateMachineState", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "message", "Message":
		v, ok := value.(string)
		if ok {
			self.Message = &v
			return nil
		}
		return fmt.Errorf("Field message/Message: value %v(%T) couldn't be cast to type string", value, value)

	case "timestamp", "Timestamp":
		v, ok := value.(int64)
		if ok {
			self.Timestamp = &v
			return nil
		}
		return fmt.Errorf("Field timestamp/Timestamp: value %v(%T) couldn't be cast to type int64", value, value)

	}
}

func (self *SingularityMachineStateHistoryUpdate) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityMachineStateHistoryUpdate", name)

	case "objectId", "ObjectId":
		return *self.ObjectId, nil
		return nil, fmt.Errorf("Field ObjectId no set on ObjectId %+v", self)

	case "state", "State":
		return *self.State, nil
		return nil, fmt.Errorf("Field State no set on State %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "message", "Message":
		return *self.Message, nil
		return nil, fmt.Errorf("Field Message no set on Message %+v", self)

	case "timestamp", "Timestamp":
		return *self.Timestamp, nil
		return nil, fmt.Errorf("Field Timestamp no set on Timestamp %+v", self)

	}
}

func (self *SingularityMachineStateHistoryUpdate) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityMachineStateHistoryUpdate", name)

	case "objectId", "ObjectId":
		self.ObjectId = nil

	case "state", "State":
		self.State = nil

	case "user", "User":
		self.User = nil

	case "message", "Message":
		self.Message = nil

	case "timestamp", "Timestamp":
		self.Timestamp = nil

	}

	return nil
}

func (self *SingularityMachineStateHistoryUpdate) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityMachineStateHistoryUpdateList []*SingularityMachineStateHistoryUpdate

func (self *SingularityMachineStateHistoryUpdateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityMachineStateHistoryUpdateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityMachineStateHistoryUpdateList cannot copy the values from %#v", other)
}

func (list *SingularityMachineStateHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityMachineStateHistoryUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityMachineStateHistoryUpdateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
