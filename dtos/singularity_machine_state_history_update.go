package dtos

import "io"

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
	Message   string                                           `json:"message"`
	ObjectId  string                                           `json:"objectId"`
	State     SingularityMachineStateHistoryUpdateMachineState `json:"state"`
	Timestamp int64                                            `json:"timestamp"`
	User      string                                           `json:"user"`
}

func (self *SingularityMachineStateHistoryUpdate) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *SingularityMachineStateHistoryUpdate) FormatText() string {
	return FormatText(self)
}

func (self *SingularityMachineStateHistoryUpdate) FormatJSON() string {
	return FormatJSON(self)
}

type SingularityMachineStateHistoryUpdateList []*SingularityMachineStateHistoryUpdate

func (list *SingularityMachineStateHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
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
	return FormatJSON(list)
}
