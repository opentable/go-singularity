package dtos

import "io"

type SingularityMachineStateHistoryUpdate struct {
	Message  string
	ObjectId string
	//	State *MachineState
	Timestamp int64
	User      string
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
