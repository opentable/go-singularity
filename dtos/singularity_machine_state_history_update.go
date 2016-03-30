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

type SingularityMachineStateHistoryUpdateList []*SingularityMachineStateHistoryUpdate

func (list SingularityMachineStateHistoryUpdateList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list SingularityMachineStateHistoryUpdateList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list SingularityMachineStateHistoryUpdateList) FormatJSON() string {
	return FormatJSON(list)
}
