package dtos

import (
	"fmt"
	"io"

	"github.com/opentable/swaggering"
)

type SingularityExpiringMachineStateMachineState string

const (
	SingularityExpiringMachineStateMachineStateMISSING_ON_STARTUP    SingularityExpiringMachineStateMachineState = "MISSING_ON_STARTUP"
	SingularityExpiringMachineStateMachineStateACTIVE                SingularityExpiringMachineStateMachineState = "ACTIVE"
	SingularityExpiringMachineStateMachineStateSTARTING_DECOMMISSION SingularityExpiringMachineStateMachineState = "STARTING_DECOMMISSION"
	SingularityExpiringMachineStateMachineStateDECOMMISSIONING       SingularityExpiringMachineStateMachineState = "DECOMMISSIONING"
	SingularityExpiringMachineStateMachineStateDECOMMISSIONED        SingularityExpiringMachineStateMachineState = "DECOMMISSIONED"
	SingularityExpiringMachineStateMachineStateDEAD                  SingularityExpiringMachineStateMachineState = "DEAD"
	SingularityExpiringMachineStateMachineStateFROZEN                SingularityExpiringMachineStateMachineState = "FROZEN"
)

type SingularityExpiringMachineState struct {
	KillTasksOnDecommissionTimeout *bool   `json:"killTasksOnDecommissionTimeout,omitempty"`
	StartMillis                    *int64  `json:"startMillis,omitempty"`
	ActionId                       *string `json:"actionId,omitempty"`
	User                           *string `json:"user,omitempty"`
	// Invalid field: ExpiringAPIRequestObject *notfound.T `json:"expiringAPIRequestObject,omitempty"`
	MachineId     *string                                      `json:"machineId,omitempty"`
	RevertToState *SingularityExpiringMachineStateMachineState `json:"revertToState,omitempty"`
}

func (self *SingularityExpiringMachineState) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, self)
}

func (self *SingularityExpiringMachineState) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringMachineState); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringMachineState cannot copy the values from %#v", other)
}

func (self *SingularityExpiringMachineState) FormatText() string {
	return swaggering.FormatText(self)
}

func (self *SingularityExpiringMachineState) FormatJSON() string {
	return swaggering.FormatJSON(self)
}

func (self *SingularityExpiringMachineState) SetField(name string, value interface{}) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringMachineState", name)

	case "killTasksOnDecommissionTimeout", "KillTasksOnDecommissionTimeout":
		v, ok := value.(bool)
		if ok {
			self.KillTasksOnDecommissionTimeout = &v
			return nil
		}
		return fmt.Errorf("Field killTasksOnDecommissionTimeout/KillTasksOnDecommissionTimeout: value %v(%T) couldn't be cast to type bool", value, value)

	case "startMillis", "StartMillis":
		v, ok := value.(int64)
		if ok {
			self.StartMillis = &v
			return nil
		}
		return fmt.Errorf("Field startMillis/StartMillis: value %v(%T) couldn't be cast to type int64", value, value)

	case "actionId", "ActionId":
		v, ok := value.(string)
		if ok {
			self.ActionId = &v
			return nil
		}
		return fmt.Errorf("Field actionId/ActionId: value %v(%T) couldn't be cast to type string", value, value)

	case "user", "User":
		v, ok := value.(string)
		if ok {
			self.User = &v
			return nil
		}
		return fmt.Errorf("Field user/User: value %v(%T) couldn't be cast to type string", value, value)

	case "machineId", "MachineId":
		v, ok := value.(string)
		if ok {
			self.MachineId = &v
			return nil
		}
		return fmt.Errorf("Field machineId/MachineId: value %v(%T) couldn't be cast to type string", value, value)

	case "revertToState", "RevertToState":
		v, ok := value.(SingularityExpiringMachineStateMachineState)
		if ok {
			self.RevertToState = &v
			return nil
		}
		return fmt.Errorf("Field revertToState/RevertToState: value %v(%T) couldn't be cast to type SingularityExpiringMachineStateMachineState", value, value)

	}
}

func (self *SingularityExpiringMachineState) GetField(name string) (interface{}, error) {
	switch name {
	default:
		return nil, fmt.Errorf("No such field %s on SingularityExpiringMachineState", name)

	case "killTasksOnDecommissionTimeout", "KillTasksOnDecommissionTimeout":
		return *self.KillTasksOnDecommissionTimeout, nil
		return nil, fmt.Errorf("Field KillTasksOnDecommissionTimeout no set on KillTasksOnDecommissionTimeout %+v", self)

	case "startMillis", "StartMillis":
		return *self.StartMillis, nil
		return nil, fmt.Errorf("Field StartMillis no set on StartMillis %+v", self)

	case "actionId", "ActionId":
		return *self.ActionId, nil
		return nil, fmt.Errorf("Field ActionId no set on ActionId %+v", self)

	case "user", "User":
		return *self.User, nil
		return nil, fmt.Errorf("Field User no set on User %+v", self)

	case "machineId", "MachineId":
		return *self.MachineId, nil
		return nil, fmt.Errorf("Field MachineId no set on MachineId %+v", self)

	case "revertToState", "RevertToState":
		return *self.RevertToState, nil
		return nil, fmt.Errorf("Field RevertToState no set on RevertToState %+v", self)

	}
}

func (self *SingularityExpiringMachineState) ClearField(name string) error {
	switch name {
	default:
		return fmt.Errorf("No such field %s on SingularityExpiringMachineState", name)

	case "killTasksOnDecommissionTimeout", "KillTasksOnDecommissionTimeout":
		self.KillTasksOnDecommissionTimeout = nil

	case "startMillis", "StartMillis":
		self.StartMillis = nil

	case "actionId", "ActionId":
		self.ActionId = nil

	case "user", "User":
		self.User = nil

	case "machineId", "MachineId":
		self.MachineId = nil

	case "revertToState", "RevertToState":
		self.RevertToState = nil

	}

	return nil
}

func (self *SingularityExpiringMachineState) LoadMap(from map[string]interface{}) error {
	return swaggering.LoadMapIntoDTO(from, self)
}

type SingularityExpiringMachineStateList []*SingularityExpiringMachineState

func (self *SingularityExpiringMachineStateList) Absorb(other swaggering.DTO) error {
	if like, ok := other.(*SingularityExpiringMachineStateList); ok {
		*self = *like
		return nil
	}
	return fmt.Errorf("A SingularityExpiringMachineStateList cannot copy the values from %#v", other)
}

func (list *SingularityExpiringMachineStateList) Populate(jsonReader io.ReadCloser) (err error) {
	return swaggering.ReadPopulate(jsonReader, list)
}

func (list *SingularityExpiringMachineStateList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *SingularityExpiringMachineStateList) FormatJSON() string {
	return swaggering.FormatJSON(list)
}
